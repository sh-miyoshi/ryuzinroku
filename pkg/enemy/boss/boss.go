package boss

import (
	"github.com/sh-miyoshi/dxlib"
	"github.com/sh-miyoshi/ryuzinroku/pkg/bullet"
	"github.com/sh-miyoshi/ryuzinroku/pkg/common"
	"github.com/sh-miyoshi/ryuzinroku/pkg/enemy/shot"
)

const (
	stdPosX  = common.FiledSizeX / 2
	stdPosY  = 100.0
	waitTime = 140
	endTime  = 99 * 60
)

const (
	modeWait int = iota
	modeBarr
)

type barrage struct {
	Type   int           `yaml:"type"`
	HP     int           `yaml:"hp"`
	Bullet bullet.Bullet `yaml:"bullet"`
}

// Boss ...
type Boss struct {
	AppearCount int       `yaml:"appearCount"`
	Final       bool      `yaml:"final"`
	Barrages    []barrage `yaml:"barrages"`

	x, y        float64
	count       int
	images      []int32
	currentBarr int
	mode        int
	move        mover
	shotProc    *shot.Shot
}

// Init ...
func (b *Boss) Init(imgs []int32) {
	b.count = 0
	b.currentBarr = -1
	b.x = float64(common.FiledSizeX) / 2
	b.y = -30
	b.images = imgs
	b.mode = modeWait
	b.shotProc = nil
	b.move.moveTo(b.x, b.y, stdPosX, stdPosY, 60)
}

// Process ...
func (b *Boss) Process() {
	// Move
	b.move.process()
	b.x, b.y = b.move.currentPos()

	// 初期状態は待機モード
	// 今が待機モードならwaitTime分待機する
	// 待機が終了したら弾幕を登録し、弾幕モードにする
	// HPが0になるか、endTime時間がたつとその弾幕を中止し、待機モードへ

	switch b.mode {
	case modeWait:
		if b.count == waitTime {
			b.count = 0
			b.mode = modeBarr
			b.currentBarr++
			barr := b.Barrages[b.currentBarr]
			b.shotProc = shot.New(barr.Type, barr.Bullet)
			return
		}
	case modeBarr:
		if b.shotProc != nil {
			b.shotProc.Process(b.x, b.y)
		}
	}

	// Check bullet hit and dead
	// 待機モードの場合はHPを減らさない
	// 弾幕モードの時、HPが0以下になるかendTimeになれば待機モードに
	// TODO
	b.count++
}

// Draw ...
func (b *Boss) Draw() {
	common.CharDraw(b.x, b.y, b.images[0], dxlib.TRUE)
}
