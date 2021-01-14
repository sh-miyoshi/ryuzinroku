package shot

import (
	"math"
	"math/rand"

	"github.com/sh-miyoshi/ryuzinroku/pkg/bullet"
	"github.com/sh-miyoshi/ryuzinroku/pkg/common"
	"github.com/sh-miyoshi/ryuzinroku/pkg/mover"
	"github.com/sh-miyoshi/ryuzinroku/pkg/player"
	"github.com/sh-miyoshi/ryuzinroku/pkg/sound"
)

type box struct {
	x1, y1 float64 // the position of top left corner
	x2, y2 float64 // the position of bottom right corner
}

func moveRandom(charID string, ex, ey float64, dist float64, targetCnt int, moveRange box) {
	// 1000回トライしてだめならあきらめる
	for i := 0; i < 1000; i++ {
		x := ex
		y := ey
		angle := common.RandomAngle(math.Pi)
		x += math.Cos(angle) * dist
		y += math.Sin(angle) * dist
		// その点が移動可能範囲なら
		if moveRange.x1 <= x && x <= moveRange.x2 && moveRange.y1 <= y && y <= moveRange.y2 {
			mover.MoveTo(charID, x, y, targetCnt)
			return
		}
	}
}

// 円形照射
func bossShotAct0(ex, ey float64, s *Shot) {
	const tm = 120
	t := s.count % tm

	if t < 60 && t%10 == 0 {
		px, py := player.GetPlayerPos()
		angle := math.Atan2(py-ey, px-ex)
		for i := 0; i < 30; i++ {
			b := s.bulletInfo
			b.CharID = s.charID
			b.ShotID = s.id
			b.X = ex
			b.Y = ey
			b.Angle = angle + math.Pi*2/30*float64(i)
			b.Speed = 3
			bullet.Register(b)
			sound.PlaySound(sound.SEEnemyShot)
		}
	}
}

// サイレントセレナ
func bossShotAct1(ex, ey float64, s *Shot) {
	const turnCnt = 60
	const loopCnt = turnCnt * 4

	t := s.count % turnCnt
	if t == 0 { // 1ターンの最初の初期化
		px, py := player.GetPlayerPos()
		s.baseAngle = math.Atan2(py-ey, px-ex)
		if s.count%loopCnt == (turnCnt * 3) { // 4ターンに１回移動
			moveRandom(s.charID, ex, ey, 60, 60, box{x1: 40, y1: 30, x2: float64(common.FiledSizeX) - 40, y2: 120})
		}
	}

	// 1ターンの最初は自機狙い、半分からは自機狙いからずらす
	if t == turnCnt/2-1 {
		s.baseAngle += math.Pi / 20 / 2
	}

	// 1ターンに10回円形発射の弾をうつ
	if t%(turnCnt/10) == 0 {
		for i := 0; i < 20; i++ {
			b := s.bulletInfo
			b.CharID = s.charID
			b.ShotID = s.id
			b.X = ex
			b.Y = ey
			b.Angle = s.baseAngle + math.Pi*2/20*float64(i) // ベース角度から20個回転して発射
			b.Speed = 2.7
			b.Color = 4
			b.Type = 8
			bullet.Register(b)
			sound.PlaySound(sound.SEEnemyShot)
		}
	}

	// 4カウントに1回下に落ちる弾を登録
	if t%4 == 0 {
		b := s.bulletInfo
		b.CharID = s.charID
		b.ShotID = s.id
		b.X = float64(rand.Intn(common.FiledSizeX))
		b.Y = float64(rand.Intn(200))
		b.Angle = math.Pi / 2 // 真下
		b.Speed = 1 + common.RandomAngle(0.5)
		b.Color = 0
		b.Type = 8
		bullet.Register(b)
		sound.PlaySound(sound.SEEnemyShot)
	}
}
