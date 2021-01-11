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
	hpImg       int32
	currentBarr int
	mode        int
	move        mover
	shotProc    *shot.Shot
	currentHP   int
}

// Init ...
func (b *Boss) Init(imgs []int32, hpImg int32) {
	b.count = 0
	b.currentBarr = -1
	b.x = float64(common.FiledSizeX) / 2
	b.y = -30
	b.images = imgs
	b.mode = modeWait
	b.hpImg = hpImg
	b.shotProc = nil
	b.currentHP = 0
	b.move.moveTo(b.x, b.y, stdPosX, stdPosY, 60)
}

// Process ...
func (b *Boss) Process() bool {
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
			b.currentHP = barr.HP
			return false
		}
	case modeBarr:
		b.shotProc.Process(b.x, b.y)

		// HPが0以下になるかendTimeになれば待機モードに

		// TODO Check bullet hit and dead

		if b.count >= endTime {
			// TODO Stop Shot
			if b.currentBarr == len(b.Barrages)-1 {
				return true // finish
			}
			b.mode = modeWait
			b.count = 0
			b.move.moveTo(b.x, b.y, stdPosX, stdPosY, 60)
		}
	}

	b.count++
	return false
}

// Draw ...
func (b *Boss) Draw() {
	common.CharDraw(b.x, b.y, b.images[0], dxlib.TRUE)

	// HP描画
	// TODO hpの色を背景色に合わせて変える
	if b.currentHP > 0 && b.currentBarr < len(b.Barrages) {
		hpSize := common.FiledSizeX * 0.98 * float64(b.currentHP) / float64(b.Barrages[b.currentBarr].HP)
		for i := 0; i < int(hpSize); i++ {
			dxlib.DrawGraph(3+int32(i)+common.FieldTopX, 2+common.FieldTopY, b.hpImg, dxlib.FALSE)
		}
	}
}
