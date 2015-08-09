package server

import (
	"encoding/json"
	"fmt"
	"github.com/Archmage-/er-maps/core"
	"github.com/Archmage-/er-maps/core/util"
	"github.com/Archmage-/er-maps/modules/minemap"
	"net/http"
)

type mapGetSector struct {
	provider minemapprovider.IProvider
}

func NewMapGetSector(provider minemapprovider.IProvider) *mapGetSector {
	return &mapGetSector{provider: provider}
}

func (this *mapGetSector) Handle(r *http.Request, rc *core.RequestContext) (data []byte, e error) {
	//check input params
	fmt.Println("GetSector")
	params := r.URL.Query()
	ints, err := util.ConvertToInts([]string{params.Get("x"), params.Get("y"), params.Get("f")})
	if err != nil {
		return nil, err
	}
	//return sector json
	sector, err := this.provider.GetSector(ints[0], ints[1], ints[2])
	if err != nil {
		return nil, err
	}
	return json.Marshal(sector)
}
