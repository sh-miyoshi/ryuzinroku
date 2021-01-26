package player

import (
	"fmt"

	"github.com/sh-miyoshi/dxlib"
	"github.com/sh-miyoshi/ryuzinroku/pkg/bullet"
	"github.com/sh-miyoshi/ryuzinroku/pkg/common"
	"github.com/sh-miyoshi/ryuzinroku/pkg/item"
)

var (
	plyr *player
)

// Init ...
func Init() error {
	fname := "data/image/effect/player_hit.png"
	hitImg := dxlib.LoadGraph(fname, dxlib.FALSE)
	if hitImg == -1 {
		return fmt.Errorf("Failed to load hit image: %s", fname)
	}

	fname = "data/image/etc/player_option.png"
	optImg := dxlib.LoadGraph(fname, dxlib.FALSE)
	if optImg == -1 {
		return fmt.Errorf("Failed to load option image: %s", fname)
	}

	var err error
	plyr, err = create(
		common.ImageInfo{FileName: "data/image/char/player.png", AllNum: 12, XNum: 4, YNum: 3, XSize: 73, YSize: 73},
		hitImg,
		optImg,
	)
	return err
}

// MgrProcess ...
func MgrProcess(ex, ey float64) bool {
	plyr.process(ex, ey)
	bullets := bullet.GetBullets()
	hits := plyr.hitProc(bullets)
	if len(hits) > 0 {
		bullet.RemoveHitBullets(hits)
	}
	plyr.laserHitProc()
	items := item.GetItems()
	plyr.itemProc(items)

	return plyr.gaveOver
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
