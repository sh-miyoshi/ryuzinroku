package player

import (
	"fmt"
	"math"

	"github.com/sh-miyoshi/dxlib"
	"github.com/sh-miyoshi/ryuzinroku/pkg/bullet"
	"github.com/sh-miyoshi/ryuzinroku/pkg/common"
	"github.com/sh-miyoshi/ryuzinroku/pkg/inputs"
	"github.com/sh-miyoshi/ryuzinroku/pkg/player/shot"
)

const (
	initShotPower = 500
	hitRange      = 2.0
)

type player struct {
	x, y            float64
	count           int
	imgCount        int
	images          []int32
	plyrShot        *shot.Shot
	invincibleCount int
}

func create(img common.ImageInfo) (*player, error) {
	if img.AllNum <= 0 {
		return nil, fmt.Errorf("image num must be positive integer, but got %d", img.AllNum)
	}

	res := player{
		x:        common.FiledSizeX / 2,
		y:        common.FiledSizeY * 3 / 4,
		plyrShot: &shot.Shot{Power: initShotPower},
	}
	res.images = make([]int32, img.AllNum)
	r := dxlib.LoadDivGraph(img.FileName, img.AllNum, img.XNum, img.YNum, img.XSize, img.YSize, res.images, dxlib.FALSE)
	if r != 0 {
		return nil, fmt.Errorf("Failed to load player image")
	}

	return &res, nil
}

func (p *player) draw() {
	if p.invincibleCount%2 == 0 {
		common.CharDraw(p.x, p.y, p.images[p.imgCount], dxlib.TRUE)
	}
}

func (p *player) process() {
	p.count++
	p.imgCount = (p.count / 6) % 4

	p.move()
	p.plyrShot.Process(p.x, p.y)
}

func (p *player) move() {
	// Check left and right moves
	moveX := 0
	if inputs.CheckKey(dxlib.KEY_INPUT_LEFT) > 0 {
		p.imgCount += 4 * 2
		moveX = -4
	} else if inputs.CheckKey(dxlib.KEY_INPUT_RIGHT) > 0 {
		p.imgCount += 4 * 1
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

	if len(hits) > 0 {
		// TODO player death
		fmt.Println("hits")
	}

	return hits
}
