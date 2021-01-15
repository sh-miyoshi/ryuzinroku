package player

import (
	"fmt"

	"github.com/sh-miyoshi/dxlib"
	"github.com/sh-miyoshi/ryuzinroku/pkg/bullet"
	"github.com/sh-miyoshi/ryuzinroku/pkg/common"
)

var (
	plyr   *player
	hitImg int32
)

// Init ...
func Init() error {
	hitImg := dxlib.LoadGraph("data/image/effect/player_hit.png", dxlib.FALSE)
	if hitImg == -1 {
		return fmt.Errorf("Failed to load hit image: data/image/effect/player_hit.png")
	}

	var err error
	plyr, err = create(
		common.ImageInfo{FileName: "data/image/char/player.png", AllNum: 12, XNum: 4, YNum: 3, XSize: 73, YSize: 73},
		hitImg,
	)
	return err
}

// MgrProcess ...
func MgrProcess() {
	plyr.process()
	bullets := bullet.GetBullets()
	hits := plyr.hitProc(bullets)
	if len(hits) > 0 {
		bullet.RemoveHitBullets(hits)
	}
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
