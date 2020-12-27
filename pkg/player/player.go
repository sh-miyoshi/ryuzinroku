package player

import (
	"fmt"

	"github.com/sh-miyoshi/dxlib"
	"github.com/sh-miyoshi/ryuzinroku/pkg/common"
)

// Player ...
type Player struct {
	x, y     int32
	count    int
	imgCount int
	images   []int32
}

// New ...
func New(img common.ImageInfo) (*Player, error) {
	if img.AllNum <= 0 {
		return nil, fmt.Errorf("image num must be positive integer, but got %d", img.AllNum)
	}

	res := Player{
		x: common.ScreenX / 2,
		y: common.ScreenY * 3 / 4,
	}
	res.images = make([]int32, img.AllNum)
	r := dxlib.LoadDivGraph(img.FileName, img.AllNum, img.XNum, img.YNum, img.XSize, img.YSize, res.images)
	if r != 0 {
		return nil, fmt.Errorf("Failed to load player image")
	}

	return &res, nil
}

// Draw ...
func (p *Player) Draw() {
	dxlib.DrawGraph(p.x, p.y, p.images[p.imgCount], dxlib.TRUE)
}

// Process ...
func (p *Player) Process() {
	p.count++
	p.imgCount = (p.count / 6) % 4
}
