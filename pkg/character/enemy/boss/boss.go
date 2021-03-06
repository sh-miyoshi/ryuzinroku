package boss

import (
	"fmt"
	"math"

	"github.com/google/uuid"
	"github.com/sh-miyoshi/dxlib"
	"github.com/sh-miyoshi/ryuzinroku/pkg/background"
	"github.com/sh-miyoshi/ryuzinroku/pkg/bullet"
	"github.com/sh-miyoshi/ryuzinroku/pkg/character/enemy/shot"
	"github.com/sh-miyoshi/ryuzinroku/pkg/common"
	"github.com/sh-miyoshi/ryuzinroku/pkg/laser"
	"github.com/sh-miyoshi/ryuzinroku/pkg/mover"
	"github.com/sh-miyoshi/ryuzinroku/pkg/score"
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

const (
	// HPColNormal ...
	HPColNormal int = iota
	// HPColBright ...
	HPColBright

	// HPColMax ...
	HPColMax
)

const (
	// TypeRiria ...
	TypeRiria int = iota

	// TypeMax ...
	TypeMax
)

// Define ...
type Define struct {
	Type        int  `yaml:"type"`
	AppearCount int  `yaml:"appearCount"`
	Final       bool `yaml:"final"`
	Barrages    []struct {
		Type      int           `yaml:"type"`
		HP        int           `yaml:"hp"`
		Bullet    bullet.Bullet `yaml:"bullet"`
		SpellCard bool          `yaml:"spellcard"`
	} `yaml:"barrages"`
}

// Boss ...
type Boss interface {
	Process(px, py float64) bool
	Draw()
	Clear()
	GetPos() (float64, float64)
}

// Riria ...
type Riria struct {
	Define

	x, y        float64
	count       int
	images      []int32
	hpImg       [HPColMax]int32
	backImgs    []int32
	currentBarr int
	mode        int
	shotProc    *shot.Shot
	currentHP   int
	charID      string
	imgCount    int
}

// NewRiria ...
func NewRiria(def Define, charImg []int32, hpImg [HPColMax]int32, backImgs []int32) (*Riria, error) {
	inst := Riria{
		Define:      def,
		x:           float64(common.FiledSizeX) / 2,
		y:           -30,
		count:       0,
		currentBarr: 0,
		mode:        modeWait,
		charID:      uuid.New().String(),
		images:      charImg,
		hpImg:       hpImg,
		backImgs:    backImgs,
		imgCount:    0,
	}

	if len(backImgs) != 3 {
		return nil, fmt.Errorf("required 3 back images, but got %d", len(backImgs))
	}

	inst.setBarr()
	mover.CharRegist(inst.charID, &inst.x, &inst.y)
	mover.MoveTo(inst.charID, stdPosX, stdPosY, 60)

	return &inst, nil
}

// Process ...
func (r *Riria) Process(px, py float64) bool {
	// 初期状態は待機モード
	// 今が待機モードならwaitTime分待機する
	// 待機が終了したら弾幕を登録し、弾幕モードにする
	// HPが0になるか、endTime時間がたつとその弾幕を中止し、待機モードへ

	switch r.mode {
	case modeWait:
		if r.count == waitTime {
			r.count = 0
			r.mode = modeBarr
			return false
		}
	case modeBarr:
		r.shotProc.Process(r.x, r.y, px, py)

		// Check bullet hit and dead
		bullets := bullet.GetBullets()
		hits := r.hitProc(bullets)
		if len(hits) > 0 {
			bullet.RemoveHitBullets(hits)
		}

		bombDm := bullet.PopBomb()
		if bombDm > 0 {
			r.currentHP -= bombDm
		}

		// HPが0以下になるかendTimeになれば待機モードに
		if r.currentHP <= 0 || r.count >= endTime {
			sound.PlaySound(sound.SEEnemyDead)
			bullet.RemoveCharBullets(r.charID)
			laser.RemoveCharLaser(r.charID)
			if r.currentBarr == len(r.Barrages)-1 {
				background.SetBack(background.BackNormal)
				sound.PlaySound(sound.SEBossDead)
				return true // finish
			}
			sound.PlaySound(sound.SEBossChange)
			r.currentBarr++
			r.mode = modeWait
			r.count = 0
			r.setBarr()
			mover.MoveTo(r.charID, stdPosX, stdPosY, 60)
		}
	}

	r.count++
	r.imgCount++
	return false
}

// Draw ...
func (r *Riria) Draw() {
	cnt := r.imgCount
	y := r.y + math.Sin(math.Pi*2/130*float64(cnt%130))*10

	dxlib.SetDrawBlendMode(dxlib.DX_BLENDMODE_ALPHA, 150)
	dxlib.DrawRotaGraphFast(int32(r.x)+common.FieldTopX, int32(y)+common.FieldTopY, float32(0.4+0.05*math.Sin(math.Pi*2/360*float64(cnt%360)))*3, float32(math.Pi*2*float64(cnt%580)/580), r.backImgs[1], dxlib.TRUE)
	dxlib.DrawRotaGraphFast(int32(r.x)+common.FieldTopX, int32(y)+common.FieldTopY, float32(0.5+0.1*math.Sin(math.Pi*2/360*float64(cnt%360)))*2, 2*math.Pi*float32(cnt%340)/340, r.backImgs[0], dxlib.TRUE)
	dxlib.DrawRotaGraphFast(int32(r.x+60*math.Sin(math.Pi*2/153*float64(cnt%153))+common.FieldTopX), int32(y+80*math.Sin(math.Pi*2/120*float64(cnt%120))+common.FieldTopY), float32(0.4+0.05*math.Sin(math.Pi*2/120*float64(cnt%120))), 2*math.Pi*float32(cnt%30)/30, r.backImgs[2], dxlib.TRUE)
	dxlib.DrawRotaGraphFast(int32(r.x+60*math.Sin(math.Pi*2/200*float64((cnt+20)%200))+common.FieldTopX), int32(y+80*math.Sin(math.Pi*2/177*float64((cnt+20)%177))+common.FieldTopY), float32(0.3+0.05*math.Sin(math.Pi*2/120*float64(cnt%120))), 2*math.Pi*float32(cnt%35)/35, r.backImgs[2], dxlib.TRUE)
	dxlib.SetDrawBlendMode(dxlib.DX_BLENDMODE_NOBLEND, 0)
	dxlib.DrawRotaGraphFast(int32(r.x+60*math.Sin(math.Pi*2/230*float64((cnt+40)%230))+common.FieldTopX), int32(y+80*math.Sin(math.Pi*2/189*float64((cnt+40)%189))+common.FieldTopY), float32(0.6+0.05*math.Sin(math.Pi*2/120*float64(cnt%120))), 2*math.Pi*float32(cnt%40)/40, r.backImgs[2], dxlib.TRUE)

	common.CharDraw(r.x, y, r.images[0], dxlib.TRUE)

	// HP描画
	if r.currentHP > 0 && r.currentBarr < len(r.Barrages) {
		col := HPColNormal
		if r.Barrages[r.currentBarr].SpellCard {
			col = HPColBright
		}

		hpSize := common.FiledSizeX * 0.98 * float64(r.currentHP) / float64(r.Barrages[r.currentBarr].HP)
		for i := 0; i < int(hpSize); i++ {
			dxlib.DrawGraph(3+int32(i)+common.FieldTopX, 2+common.FieldTopY, r.hpImg[col], dxlib.FALSE)
		}
	}
}

// Clear ...
func (r *Riria) Clear() {
	mover.CharRemove(r.charID)
}

// GetPos ...
func (r *Riria) GetPos() (float64, float64) {
	if r.mode == modeBarr {
		return r.x, r.y
	}
	return -1, -1
}

func (r *Riria) hitProc(bullets []*bullet.Bullet) []int {
	res := []int{}
	for i, bl := range bullets {
		if !bl.IsPlayer {
			continue
		}

		x := bl.X - r.x
		y := bl.Y - r.y
		hr := bl.HitRange + hitRange

		if x*x+y*y < hr*hr { // 当たり判定内なら
			r.currentHP -= bl.Power
			sound.PlaySound(sound.SEEnemyHit)
			res = append(res, i)
			score.Set(score.TypeScore, common.UpMax(score.Get(score.TypeScore), 1, 999999999))
			continue
		}

		// 中間を計算する必要があれば
		if bl.Speed > hr {
			// 1フレーム前にいた位置
			preX := bl.X + math.Cos(bl.Angle+math.Pi)*bl.Speed
			preY := bl.Y + math.Sin(bl.Angle+math.Pi)*bl.Speed
			for j := 0; j < int(bl.Speed/hr); j++ { // 進んだ分÷当たり判定分ループ
				px := preX - r.x
				py := preY - r.y
				if px*px+py*py < hr*hr {
					r.currentHP -= bl.Power
					sound.PlaySound(sound.SEEnemyHit)
					res = append(res, i)
					score.Set(score.TypeScore, common.UpMax(score.Get(score.TypeScore), 1, 999999999))
					break
				}
				preX += math.Cos(bl.Angle) * bl.Speed
				preY += math.Sin(bl.Angle) * bl.Speed
			}
		}
	}
	return res
}

func (r *Riria) setBarr() {
	barr := r.Barrages[r.currentBarr]
	r.shotProc = shot.New(barr.Type, r.charID, barr.Bullet)
	r.currentHP = barr.HP
	if barr.SpellCard {
		background.SetBack(background.BackSpellCard)
	} else {
		background.SetBack(background.BackNormal)
	}
}
