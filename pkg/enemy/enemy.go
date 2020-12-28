package enemy

import (
	"github.com/sh-miyoshi/dxlib"
	"github.com/sh-miyoshi/ryuzinroku/pkg/common"
)

type enemy struct {
	ApperCount  int   `yaml:"appearCount"`
	MovePattern int   `yaml:"movePattern"`
	Type        int   `yaml:"type"`
	X           int32 `yaml:"x"`
	Y           int32 `yaml:"y"`
	HP          int   `yaml:"hp"`

	images   []int32
	imgSizeX int32
	imgSizeY int32
	imgCount int
	count    int
	dead     bool
}

// Process ...
func (e *enemy) Process() {
	e.count++
	switch e.MovePattern {
	case 0:
		act0(e)
		break
	default:
		panic("Invalid move pattern")
	}

	if e.HP <= 0 {
		e.dead = true
		return
	}
	if e.X < -50 || e.X > common.ScreenX+50 || e.Y < -50 || e.Y > common.ScreenY+50 {
		e.dead = true
		return
	}
}

// Draw ...
func (e *enemy) Draw() {
	centerX := e.X - e.imgSizeX/2
	centerY := e.Y - e.imgSizeY/2
	dxlib.DrawGraph(centerX, centerY, e.images[e.imgCount], dxlib.TRUE)
}
