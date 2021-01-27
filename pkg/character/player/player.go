package player

import (
	"fmt"
	"math"

	"github.com/sh-miyoshi/dxlib"
	"github.com/sh-miyoshi/ryuzinroku/pkg/bullet"
	"github.com/sh-miyoshi/ryuzinroku/pkg/character/player/shot"
	"github.com/sh-miyoshi/ryuzinroku/pkg/common"
	"github.com/sh-miyoshi/ryuzinroku/pkg/effect"
	"github.com/sh-miyoshi/ryuzinroku/pkg/inputs"
	"github.com/sh-miyoshi/ryuzinroku/pkg/item"
	"github.com/sh-miyoshi/ryuzinroku/pkg/laser"
	"github.com/sh-miyoshi/ryuzinroku/pkg/score"
	"github.com/sh-miyoshi/ryuzinroku/pkg/sound"
)

const (
	// InitRemainNum ...
	InitRemainNum = 5

	initShotPower   = 300
	hitRange        = 2.0
	itemGetBorder   = 100.0
	itemAbsorbRange = 70.0
	itemHitRange    = 20.0
	maxShotPower    = 500
)

const (
	stateNormal int = iota
	stateDead
)

type player struct {
	x, y            float64
	count           int
	imgNo           int
	images          []int32
	plyrShot        *shot.Shot
	invincibleCount int
	state           int
	slow            bool
	hitImg          int32
	gaveOver        bool
}

func create(img common.ImageInfo, hitImg, optImg int32) (*player, error) {
	if img.AllNum <= 0 {
		return nil, fmt.Errorf("image num must be positive integer, but got %d", img.AllNum)
	}

	res := player{
		x:        common.FiledSizeX / 2,
		y:        common.FiledSizeY * 3 / 4,
		plyrShot: shot.New(optImg, initShotPower),
		state:    stateNormal,
		hitImg:   hitImg,
		gaveOver: false,
	}
	res.images = make([]int32, img.AllNum)
	r := dxlib.LoadDivGraph(img.FileName, img.AllNum, img.XNum, img.YNum, img.XSize, img.YSize, res.images, dxlib.FALSE)
	if r != 0 {
		return nil, fmt.Errorf("Failed to load player image")
	}

	score.Set(score.TypeRemainNum, InitRemainNum)
	score.Set(score.TypePlayerPower, initShotPower)

	return &res, nil
}

func (p *player) draw() {
	if p.invincibleCount%2 == 0 {
		common.CharDraw(p.x, p.y, p.images[p.imgNo], dxlib.TRUE)
	}

	if p.slow {
		dxlib.DrawRotaGraphFast(int32(p.x)+common.FieldTopX, int32(p.y)+common.FieldTopY, 1, math.Pi*2*float32(p.count%120)/120, p.hitImg, dxlib.TRUE, dxlib.FALSE, dxlib.FALSE)
	}

	p.plyrShot.Draw()
}

func (p *player) process(ex, ey float64) {
	p.count++
	p.imgNo = (p.count / 6) % 4

	p.slow = inputs.CheckKey(dxlib.KEY_INPUT_LSHIFT) > 0

	switch p.state {
	case stateNormal:
		p.move()
		p.plyrShot.Process(p.x, p.y, ex, ey, p.slow)
		if inputs.CheckKey(dxlib.KEY_INPUT_X) == 1 {
			if err := effect.Register(effect.Controller{
				Type: effect.ControllerTypeBomb,
				X:    p.x,
				Y:    p.y,
			}); err == nil {
				// 無敵状態に
				p.invincibleCount = 1
				// ダメージの登録
				bullet.PushBomb(p.plyrShot.Power / 20)
			}
		}
	case stateDead:
		p.y -= 1.5

		input := inputs.CheckKey(dxlib.KEY_INPUT_LEFT) + inputs.CheckKey(dxlib.KEY_INPUT_RIGHT) + inputs.CheckKey(dxlib.KEY_INPUT_UP) + inputs.CheckKey(dxlib.KEY_INPUT_DOWN)
		//１秒以上か、キャラがある程度上にいて、何かおされたら
		if p.count > 60 || (p.y < float64(common.FiledSizeY)-20 && input != 0) {
			p.count = 0
			p.state = stateNormal
		}
	}

	if p.invincibleCount > 0 {
		p.invincibleCount++
		if p.invincibleCount > 120 {
			p.invincibleCount = 0 // 無敵状態終了
		}
	}
}

func (p *player) move() {
	// Check left and right moves
	moveX := 0
	if inputs.CheckKey(dxlib.KEY_INPUT_LEFT) > 0 {
		p.imgNo += 4 * 2
		moveX = -4
	} else if inputs.CheckKey(dxlib.KEY_INPUT_RIGHT) > 0 {
		p.imgNo += 4 * 1
		moveX = 4
	}

	// Check up and down moves
	moveY := 0
	if inputs.CheckKey(dxlib.KEY_INPUT_UP) > 0 {
		moveY = -4
	} else if inputs.CheckKey(dxlib.KEY_INPUT_DOWN) > 0 {
		moveY = 4
	}

	if moveX != 0 || moveY != 0 {
		if moveX != 0 && moveY != 0 {
			// 斜め移動
			moveX = int(float64(moveX) / math.Sqrt(2))
			moveY = int(float64(moveY) / math.Sqrt(2))
		}

		// 低速移動
		if p.slow {
			moveX /= 3
			moveY /= 3
		}

		mx := int(p.x) + moveX
		my := int(p.y) + moveY
		if common.MoveOK(mx, my) {
			p.x = float64(mx)
			p.y = float64(my)
		}
	}
}

func (p *player) hitProc(bullets []*bullet.Bullet) []int {
	hits := []int{}
	for i, b := range bullets {
		if b.IsPlayer {
			continue
		}

		x := b.X - p.x
		y := b.Y - p.y
		r := b.HitRange + hitRange

		if x*x+y*y < r*r { // 当たり判定内なら
			hits = append(hits, i)
			continue
		}

		// 中間を計算する必要があれば
		if b.Speed > r {
			// 1フレーム前にいた位置
			preX := b.X + math.Cos(b.Angle+math.Pi)*b.Speed
			preY := b.Y + math.Sin(b.Angle+math.Pi)*b.Speed
			for j := 0; j < int(b.Speed/r); j++ { // 進んだ分÷当たり判定分ループ
				px := preX - p.x
				py := preY - p.y
				if px*px+py*py < r*r {
					hits = append(hits, i)
					break
				}
				preX += math.Cos(b.Angle) * b.Speed
				preY += math.Sin(b.Angle) * b.Speed
			}
		}
	}

	// ヒットした弾が存在し、無敵状態でないなら
	if len(hits) > 0 && p.invincibleCount == 0 {
		p.death()
	}

	return hits
}

func (p *player) laserHitProc() {
	if laser.IsHit(p.x, p.y, hitRange) && p.invincibleCount == 0 {
		p.death()
	}
}

func (p *player) absorbItem(itm *item.Item) {
	v := 3.0
	if itm.State == item.StateAbsorb {
		v = 8.0
	}
	angle := math.Atan2(p.y-itm.Y, p.x-itm.X)

	itm.X += math.Cos(angle) * v
	itm.Y += math.Sin(angle) * v
}

func (p *player) itemProc(items []*item.Item) {
	for i := 0; i < len(items); i++ {
		x := p.x - items[i].X
		y := p.y - items[i].Y

		// ボーダーラインより上にいればアイテムを引き寄せる
		if p.y < itemGetBorder {
			items[i].State = item.StateAbsorb
		}
		if items[i].State == item.StateAbsorb {
			p.absorbItem(items[i])
		} else {
			// slow modeならアイテムを引き寄せる
			if p.slow && (x*x+y*y) < (itemAbsorbRange*itemAbsorbRange) {
				p.absorbItem(items[i])
			}
		}
		// 一定より近くにあればアイテムを取得する
		if (x*x + y*y) < (itemHitRange * itemHitRange) {
			switch items[i].Type {
			case item.TypePowerS:
				p.plyrShot.Power = common.UpMax(p.plyrShot.Power, 3, maxShotPower)
				score.Set(score.TypePlayerPower, p.plyrShot.Power)
			case item.TypePowerL:
				p.plyrShot.Power = common.UpMax(p.plyrShot.Power, 50, maxShotPower)
				score.Set(score.TypePlayerPower, p.plyrShot.Power)
			case item.TypePointS:
				score.Set(score.TypeScore, common.UpMax(score.Get(score.TypeScore), 100, 999999999))
			case item.TypeMoneyS:
				score.Set(score.TypeMoney, common.UpMax(score.Get(score.TypeMoney), 10, 999999))
			case item.TypeMoneyL:
				score.Set(score.TypeMoney, common.UpMax(score.Get(score.TypeMoney), 100, 999999))
			}
			sound.PlaySound(sound.SEItemGet)
			items[i].State = item.StateGot
		}
	}

}

func (p *player) death() {
	sound.PlaySound(sound.SEPlayerDead)

	remain := score.Get(score.TypeRemainNum)
	remain--
	score.Set(score.TypeRemainNum, remain)
	if remain == 0 {
		// game over
		p.gaveOver = true
		return
	}

	for i := 0; i < 4; i++ {
		item.Register(item.Item{
			Type: item.TypePowerL,
			X:    p.x + common.RandomAngle(40),
			Y:    p.y + common.RandomAngle(40),
			VY:   -3.5,
		})
	}
	p.state = stateDead
	p.invincibleCount++
	p.count = 0
	p.x = float64(common.FiledSizeX) / 2
	p.y = float64(common.FiledSizeY) + 30
}
