package maps

import (
	"fmt"
)

type MineMap struct {
	Sectors []MineSector
}

func (this *MineMap) GetSector(x, y, floor int) *MineSector {
	for _, sector := range this.Sectors {
		if sector.X == x && sector.Y == y && sector.Floor == floor {
			return &sector
		}
	}
	return emptyMineSector(x, y, floor)
}

type MineSector struct {
	X       int
	Y       int
	Floor   int
	Type    MineSectorType
	Object  MineObjectType
	OwnerId int
	Ore     OreType
	Ammount int
	//
	IsEmptySector bool
}

func emptyMineSector(x, y, floor int) *MineSector {
	return &MineSector{X: x, Y: y, Floor: floor, IsEmptySector: true}
}

func (this *MineSector) SetTypeI(t int) error {
	this.Type = MineSectorType(t)
	return nil
}
func (this *MineSector) SetObjectI(obj int) error {
	this.Object = MineObjectType(obj)
	return nil
}
func (this *MineSector) SetOreI(ore int) error {
	if ore < 0 || ore > 7 {
		return fmt.Errorf("Unknown ore type!")
	}
	this.Ore = OreType(ore)
	return nil
}

type MineSectorType int
type MineObjectType int
type OreType int

const (
	ORETYPE_NONE       OreType = 0
	ORETYPE_COAL       OreType = 1
	ORETYPE_COPPER     OreType = 2
	ORETYPE_TIN        OreType = 3
	ORETYPE_IRON       OreType = 4
	ORETYPE_PLATINUM   OreType = 5
	ORETYPE_MITHRIL    OreType = 6
	ORETYPE_ADAMANTINE OreType = 7
)
