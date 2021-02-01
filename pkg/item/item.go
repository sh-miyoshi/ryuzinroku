package item

import (
	"fmt"
	"math"

	"github.com/sh-miyoshi/dxlib"
	"github.com/sh-miyoshi/ryuzinroku/pkg/common"
)

// Item ...
type Item struct {
	Type uint `yaml:"type"`

	X, Y   float64
	VX, VY float64
	State  int

	count   int
	extRate float64
}

const (
	// StateNormal ...
	StateNormal int = iota
	// StateAbsorb ...
	StateAbsorb
	// StateGot ...
	StateGot
)

const (
	// TypePowerS is small power
	TypePowerS uint = iota
	// TypePointS is small point
	TypePointS
	// TypeMoneyS is small money
	TypeMoneyS
	// TypePowerL is large power
	TypePowerL
	// TypeMoneyL is large money
	TypeMoneyL

	typeMax
)

type itemDefine struct {
	info    common.ImageInfo
	images  []int32
	extRate float64
}

var (
	itemDefs = [typeMax]itemDefine{
		{ // TypePowerS
			info:    common.ImageInfo{XSize: 35, YSize: 35},
			extRate: 0.6,
		},
		{ // TypePointS
			info:    common.ImageInfo{XSize: 35, YSize: 35},
			extRate: 0.6,
		},
		{ // TypeMoneyS
			info:    common.ImageInfo{XSize: 35, YSize: 35},
			extRate: 0.6,
		},
		{ // TypePowerL
			info:    common.ImageInfo{XSize: 35, YSize: 35},
			extRate: 1,
		},
		{ // TypeMoneyL
			info:    common.ImageInfo{XSize: 35, YSize: 35},
			extRate: 1,
		},
	}

	items []*Item
)

// Init ...
func Init() error {
	for i, def := range itemDefs {
		itemDefs[i].images = make([]int32, 2)
		fname := fmt.Sprintf("data/image/item/p%d.png", i)
		if res := dxlib.LoadDivGraph(fname, 2, 2, 1, def.info.XSize, def.info.YSize, itemDefs[i].images); res == -1 {
			return fmt.Errorf("Failed to load item image: %s", fname)
		}
	}

	return nil
}

// Register ...
func Register(i Item) {
	if i.Type >= typeMax {
		panic(fmt.Sprintf("Item type overflow. No :%d, Max: %d", i.Type, typeMax))
	}

	i.extRate = itemDefs[int(i.Type)].extRate

	items = append(items, &i)
}

// MgrProcess ...
func MgrProcess() {
	newList := []*Item{}
	for _, i := range items {
		if i.State == StateGot {
			continue
		}

		i.count++

		i.X += i.VX
		i.Y += i.VY

		if i.VY < 2.5 {
			i.VY += 0.06
		}

		if i.Y <= float64(common.FiledSizeY)+50 {
			newList = append(newList, i)
		}
	}
	items = newList
}

// MgrDraw ...
func MgrDraw() {
	for _, i := range items {
		angle := math.Pi * 2 * float32(i.count%120) / 120
		dxlib.DrawRotaGraphFast(int32(i.X)+common.FieldTopX, int32(i.Y)+common.FieldTopY, float32(i.extRate), angle, itemDefs[int(i.Type)].images[1], dxlib.TRUE)
		dxlib.DrawRotaGraphFast(int32(i.X)+common.FieldTopX, int32(i.Y)+common.FieldTopY, float32(i.extRate)*0.8, -angle, itemDefs[int(i.Type)].images[1], dxlib.TRUE)
		dxlib.DrawRotaGraphFast(int32(i.X)+common.FieldTopX, int32(i.Y)+common.FieldTopY, float32(i.extRate), 0, itemDefs[int(i.Type)].images[0], dxlib.TRUE)
	}
}

// GetItems ...
func GetItems() []*Item {
	return items
}
