package enemy

import (
	"github.com/sh-miyoshi/dxlib"
	"github.com/sh-miyoshi/ryuzinroku/pkg/common"
)

var (
	acts = []func(*enemy){act0, act1}
)

type enemy struct {
	ApperCount  int   `yaml:"appearCount"`
	MovePattern int   `yaml:"movePattern"`
	Type        int   `yaml:"type"`
	X           int32 `yaml:"x"`
	Y           int32 `yaml:"y"`
	HP          int   `yaml:"hp"`
	Wait        int   `yaml:"wait"`

	images   []int32
	imgSizeX int32
	imgSizeY int32
	imgCount int
	count    int
	dead     bool
	direct   common.Direct
	vx, vy   float64
}

// Process ...
func (e *enemy) Process() {
	acts[e.MovePattern](e)

	e.count++
	e.imgCount = (e.count / 6) % 3
	switch e.direct {
	case common.DirectFront:
		e.imgCount += 3
	case common.DirectRight:
		e.imgCount += 6
	}

	e.X = int32(float64(e.X) + e.vx)
	e.Y = int32(float64(e.Y) + e.vy)

	if e.HP <= 0 {
		e.dead = true
		return
	}
	if e.X < -50 || e.X > common.FiledSizeX+50 || e.Y < -50 || e.Y > common.FiledSizeY+50 {
		e.dead = true
		return
	}
}

// Draw ...
func (e *enemy) Draw() {
	common.CharDraw(e.X, e.Y, e.imgSizeX, e.imgSizeY, e.images[e.imgCount], dxlib.TRUE)
}
