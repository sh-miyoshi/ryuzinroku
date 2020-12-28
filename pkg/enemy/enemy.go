package enemy

import "github.com/sh-miyoshi/dxlib"

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
}

// Process ...
func (e *enemy) Process() {
}

// Draw ...
func (e *enemy) Draw() {
	centerX := e.X - e.imgSizeX/2
	centerY := e.Y - e.imgSizeY/2
	dxlib.DrawGraph(centerX, centerY, e.images[e.imgCount], dxlib.TRUE)
}
