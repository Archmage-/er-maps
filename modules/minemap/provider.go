package minemapprovider

import (
	"fmt"
	"github.com/Archmage-/er-maps/core"
	"github.com/Archmage-/er-maps/core/util"
	"github.com/Archmage-/er-maps/data/maps"
	"io/ioutil"
	"strings"
	"time"
)

type provider struct {
	url       string
	refresh   int
	ticker    *time.Ticker
	latestMap *maps.MineMap
}

type IProvider interface {
	GetSector(x, y, floor int) (*maps.MineSector, error)
}

func NewProvider(url string, refresh int) *provider {
	prov := provider{url: url, refresh: refresh}
	prov.ticker = time.NewTicker(time.Duration(refresh) * time.Second)
	prov.tickerFunc()
	return &prov
}

func (this *provider) GetSector(x, y, floor int) (*maps.MineSector, error) {
	if this.latestMap == nil {
		return nil, fmt.Errorf("Map not initialzied!")
	}
	return this.latestMap.GetSector(x, y, floor), nil
}

func (this *provider) tickerFunc() {
	go func() {
		this.loadMap(this.url)
		for _ = range this.ticker.C {
			this.loadMap(this.url)
		}
	}()
}

func (this *provider) loadMap(url string) error {
	resp, err := core.SimpleClient.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	minemap, err := this.parseBody(string(body))
	if err != nil {
		return err
	}
	this.latestMap = minemap
	return nil
}

func (this *provider) parseBody(body string) (*maps.MineMap, error) {
	lines := strings.Split(body, "\n")
	m := maps.MineMap{}
	for _, line := range lines {
		items := strings.Split(line, "|")
		sector, err := this.parseSector(items)
		if err != nil {
			//log
			continue
		}
		m.Sectors = append(m.Sectors, *sector)
	}
	return &m, nil
}

func (this *provider) parseSector(strItems []string) (*maps.MineSector, error) {
	if len(strItems) < 8 {
		return nil, fmt.Errorf("Wrong number of items (%d), %#v", len(strItems), strItems)
	}
	intItems, err := util.ConvertToInts(strItems)
	if err != nil {
		return nil, err
	}
	sector := maps.MineSector{}
	sector.X = intItems[0]
	sector.Y = intItems[1]
	sector.Floor = intItems[2]
	err = sector.SetTypeI(intItems[3]) //MineSectorType
	if err != nil {
		return nil, err
	}
	err = sector.SetObjectI(intItems[4]) //MineObjectType
	if err != nil {
		return nil, err
	}
	sector.OwnerId = intItems[5]
	err = sector.SetOreI(intItems[6]) //OreType
	if err != nil {
		return nil, err
	}
	sector.Ammount = intItems[7]
	return &sector, nil
}
