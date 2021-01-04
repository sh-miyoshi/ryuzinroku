package effect

import (
	"fmt"

	"github.com/sh-miyoshi/dxlib"
	"github.com/sh-miyoshi/ryuzinroku/pkg/common"
)

const (
	// ControllerTypeDead ...
	ControllerTypeDead int = iota
)

const (
	effectDead int = iota
	effectMax
)

type effect struct {
	Type       int
	Color      int
	BlendType  int32
	BlendParam int32
	X, Y       float64
	Angle      float64
	ExtRate    float64

	count int
}

// Controller ...
type Controller struct {
	Type  int
	Color int
	X, Y  float64

	count int
}

var (
	effects           []*effect
	effectControllers []*Controller
	effectImgs        [][]int32

	controllerActs = []func(*Controller) bool{ctrlAct0}
	effectActs     = []func(*effect) bool{effectAct0}
)

// Init ...
func Init() error {
	effectImgs = make([][]int32, effectMax)
	effectImgs[0] = make([]int32, 5)
	if res := dxlib.LoadDivGraph("data/image/effect/death.png", 5, 5, 1, 140, 140, effectImgs[0], dxlib.FALSE); res == -1 {
		return fmt.Errorf("Failed to load image: data/image/effect/death.png")
	}

	return nil
}

// MgrProcess ...
func MgrProcess() {
	// Controller処理
	newCtrls := []*Controller{}
	for _, c := range effectControllers {
		if !controllerActs[c.Type](c) {
			newCtrls = append(newCtrls, c)
		}
		c.count++
	}
	effectControllers = newCtrls

	// Effect処理
	newEffects := []*effect{}
	for _, e := range effects {
		if !effectActs[e.Type](e) {
			newEffects = append(newEffects, e)
		}
		e.count++
	}
	effects = newEffects
}

// MgrDraw ...
func MgrDraw() {
	for _, e := range effects {
		if e.BlendType != dxlib.DX_BLENDMODE_NOBLEND {
			dxlib.SetDrawBlendMode(e.BlendType, e.BlendParam)
		}

		dxlib.DrawRotaGraph(int32(e.X)+common.FieldTopX, int32(e.Y)+common.FieldTopY, e.ExtRate, e.Angle, effectImgs[e.Type][e.Color], dxlib.TRUE, dxlib.FALSE, dxlib.FALSE)

		if e.BlendType != dxlib.DX_BLENDMODE_NOBLEND {
			dxlib.SetDrawBlendMode(dxlib.DX_BLENDMODE_NOBLEND, 0)
		}
	}
}

// Register ...
func Register(e Controller) {
	effectControllers = append(effectControllers, &e)
}
