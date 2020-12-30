package player

import (
	"github.com/sh-miyoshi/ryuzinroku/pkg/common"
)

var (
	plyr *player
)

// Init ...
func Init() error {
	var err error
	plyr, err = create(common.ImageInfo{FileName: "data/image/char/player.png", AllNum: 12, XNum: 4, YNum: 3, XSize: 73, YSize: 73})
	return err
}

// MgrProcess ...
func MgrProcess() {
	plyr.process()
}

// MgrDraw ...
func MgrDraw() {
	plyr.draw()
}

// GetPlayerPos ...
func GetPlayerPos() (x, y float64) {
	if plyr != nil {
		return plyr.x, plyr.y
	}
	return 0, 0
}
