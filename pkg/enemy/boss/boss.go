package boss

import (
	"github.com/sh-miyoshi/dxlib"
	"github.com/sh-miyoshi/ryuzinroku/pkg/common"
)

const (
	stdPosX = common.FiledSizeX / 2
	stdPosY = 100.0
)

// Boss ...
type Boss struct {
	AppearCount    int  `yaml:"appearCount"`
	Final          bool `yaml:"final"`
	ShotStartCount int  `yaml:"shotStartCount"`
	HP             int  `yaml:"hp"`

	x, y    float64
	endTime int
	count   int
	images  []int32
	move    mover
}

// Init ...
func (b *Boss) Init(imgs []int32) {
	b.count = 0
	b.endTime = 99 * 60
	b.x = float64(common.FiledSizeX) / 2
	b.y = -30
	b.images = imgs
	b.move.moveTo(b.x, b.y, stdPosX, stdPosY, 60)
}

// Process ...
func (b *Boss) Process() {
	b.move.process()
	b.x, b.y = b.move.currentPos()
}

// Draw ...
func (b *Boss) Draw() {
	common.CharDraw(b.x, b.y, b.images[0], dxlib.TRUE)
}
