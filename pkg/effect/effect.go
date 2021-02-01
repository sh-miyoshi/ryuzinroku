package effect

import (
	"fmt"
	"math"

	"github.com/sh-miyoshi/dxlib"
	"github.com/sh-miyoshi/ryuzinroku/pkg/common"
)

const (
	// ControllerTypeDead ...
	ControllerTypeDead int = iota
	// ControllerTypeBomb ...
	ControllerTypeBomb
)

const (
	effectDead int = iota
	effectBombChar
	effectBombTitle
	effectBombMain

	effectMax
)

type effect struct {
	Type       int
	Color      int
	BlendType  int32
	BlendParam float64
	X, Y       float64
	Angle      float64
	ExtRate    float64
	MoveAngle  float64
	Speed      float64

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

	controllerActs = []func(*Controller) bool{ctrlAct0, ctrlAct1}
	effectActs     = []func(*effect) bool{
		// 順番はeffect typeと合わせる
		effectActDead,
		effectActBombChar,
		effectActBombTitle,
		effectActBombMain,
	}
)

// Init ...
func Init() error {
	effectImgs = make([][]int32, effectMax)
	effectImgs[effectDead] = make([]int32, 5)
	if res := dxlib.LoadDivGraph("data/image/effect/death.png", 5, 5, 1, 140, 140, effectImgs[effectDead]); res == -1 {
		return fmt.Errorf("Failed to load image: data/image/effect/death.png")
	}

	effectImgs[effectBombChar] = make([]int32, 1)
	effectImgs[effectBombChar][0] = dxlib.LoadGraph("data/image/effect/bom_char.png")
	if effectImgs[effectBombChar][0] == -1 {
		return fmt.Errorf("Failed to load image: data/image/effect/bom_char.png")
	}
	effectImgs[effectBombTitle] = make([]int32, 1)
	effectImgs[effectBombTitle][0] = dxlib.LoadGraph("data/image/effect/bom_title.png")
	if effectImgs[effectBombTitle][0] == -1 {
		return fmt.Errorf("Failed to load image: data/image/effect/bom_title.png")
	}
	effectImgs[effectBombMain] = make([]int32, 1)
	effectImgs[effectBombMain][0] = dxlib.LoadGraph("data/image/effect/bom_main.png")
	if effectImgs[effectBombMain][0] == -1 {
		return fmt.Errorf("Failed to load image: data/image/effect/bom_main.png")
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
		e.X += math.Cos(e.MoveAngle) * e.Speed
		e.Y += math.Sin(e.MoveAngle) * e.Speed
		e.count++
	}
	effects = newEffects
}

// MgrDraw ...
func MgrDraw() {
	for _, e := range effects {
		if e.BlendType != dxlib.DX_BLENDMODE_NOBLEND {
			dxlib.SetDrawBlendMode(e.BlendType, int32(e.BlendParam))
		}

		dxlib.DrawRotaGraphFast(int32(e.X)+common.FieldTopX, int32(e.Y)+common.FieldTopY, float32(e.ExtRate), float32(e.Angle), effectImgs[e.Type][e.Color], dxlib.TRUE)

		if e.BlendType != dxlib.DX_BLENDMODE_NOBLEND {
			dxlib.SetDrawBlendMode(dxlib.DX_BLENDMODE_NOBLEND, 0)
		}
	}
}

// Register ...
func Register(e Controller) error {
	// ボムは二重登録禁止
	if e.Type == ControllerTypeBomb {
		for _, c := range effectControllers {
			if c.Type == ControllerTypeBomb {
				return fmt.Errorf("the effect was already exists")
			}
		}
	}

	e.count = 0
	effectControllers = append(effectControllers, &e)
	return nil
}
