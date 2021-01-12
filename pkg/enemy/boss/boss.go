package boss

import (
	"math"

	"github.com/google/uuid"
	"github.com/sh-miyoshi/dxlib"
	"github.com/sh-miyoshi/ryuzinroku/pkg/background"
	"github.com/sh-miyoshi/ryuzinroku/pkg/bullet"
	"github.com/sh-miyoshi/ryuzinroku/pkg/common"
	"github.com/sh-miyoshi/ryuzinroku/pkg/enemy/shot"
	"github.com/sh-miyoshi/ryuzinroku/pkg/sound"
)

const (
	stdPosX  = common.FiledSizeX / 2
	stdPosY  = 100.0
	waitTime = 140
	endTime  = 99 * 60
	hitRange = 40.0
)

const (
	modeWait int = iota
	modeBarr
)

type barrage struct {
	Type      int           `yaml:"type"`
	HP        int           `yaml:"hp"`
	Bullet    bullet.Bullet `yaml:"bullet"`
	SpellCard bool          `yaml:"spellcard"`
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
	charID      string
}

// Init ...
func (b *Boss) Init(imgs []int32, hpImg int32) {
	b.count = 0
	b.currentBarr = 0
	b.x = float64(common.FiledSizeX) / 2
	b.y = -30
	b.images = imgs
	b.mode = modeWait
	b.hpImg = hpImg
	b.charID = uuid.New().String()
	b.setBarr()
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
			return false
		}
	case modeBarr:
		b.shotProc.Process(b.x, b.y)

		// Check bullet hit and dead
		bullets := bullet.GetBullets()
		hits := b.hitProc(bullets)
		if len(hits) > 0 {
			bullet.RemoveHitBullets(hits)
		}

		// HPが0以下になるかendTimeになれば待機モードに
		if b.currentHP <= 0 || b.count >= endTime {
			sound.PlaySound(sound.SEEnemyDead)
			bullet.RemoveCharBullets(b.charID)
			if b.currentBarr == len(b.Barrages)-1 {
				background.SetBack(background.BackNormal)
				return true // finish
			}
			b.currentBarr++
			b.mode = modeWait
			b.count = 0
			b.setBarr()
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

func (b *Boss) hitProc(bullets []*bullet.Bullet) []int {
	res := []int{}
	for i, bl := range bullets {
		if !bl.IsPlayer {
			continue
		}

		x := bl.X - b.x
		y := bl.Y - b.y
		r := bl.HitRange + hitRange

		if x*x+y*y < r*r { // 当たり判定内なら
			b.currentHP -= bl.Power
			sound.PlaySound(sound.SEEnemyHit)
			res = append(res, i)
			continue
		}

		// 中間を計算する必要があれば
		if bl.Speed > r {
			// 1フレーム前にいた位置
			preX := bl.X + math.Cos(bl.Angle+math.Pi)*bl.Speed
			preY := bl.Y + math.Sin(bl.Angle+math.Pi)*bl.Speed
			for j := 0; j < int(bl.Speed/r); j++ { // 進んだ分÷当たり判定分ループ
				px := preX - b.x
				py := preY - b.y
				if px*px+py*py < r*r {
					b.currentHP -= bl.Power
					sound.PlaySound(sound.SEEnemyHit)
					res = append(res, i)
					break
				}
				preX += math.Cos(bl.Angle) * bl.Speed
				preY += math.Sin(bl.Angle) * bl.Speed
			}
		}
	}
	return res
}

func (b *Boss) setBarr() {
	barr := b.Barrages[b.currentBarr]
	b.shotProc = shot.New(barr.Type, b.charID, barr.Bullet)
	b.currentHP = barr.HP
	if barr.SpellCard {
		background.SetBack(background.BackSpellCard)
	} else {
		background.SetBack(background.BackNormal)
	}
}
