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

// パーフェクトフリーズ
func bossShotAct2(ex, ey float64, s *Shot) {
	const tm = 650

	t := s.count % tm
	if t == 0 || t == 210 {
		moveRandom(s.charID, ex, ey, 100, 80, box{x1: 40, y1: 50, x2: float64(common.FiledSizeX) - 40, y2: 150})
	}

	// 最初のランダム発射
	if t < 180 {
		for i := 0; i < 2; i++ {
			b := s.bulletInfo
			b.CharID = s.charID
			b.ShotID = s.id
			b.X = ex
			b.Y = ey
			b.Color = rand.Intn(6)
			b.Type = 7
			b.Angle = common.RandomAngle(math.Pi*2/20) + math.Pi*2/10*float64(t)
			b.Speed = 3.2 + common.RandomAngle(2.1)
			b.Count = t // 同時に止めるためあえて0からスタートさせない
			b.Rotate = true
			b.Act = func(b *bullet.Bullet) {
				if b.Count == 190 {
					// 停止
					b.Speed = 0
					b.Color = 9
					b.Rotate = false
				}
				if b.Count == 390 {
					b.Angle = common.RandomAngle(math.Pi)
					b.Rotate = true
				}
				if b.Count > 390 {
					b.Speed += 0.01
				}
			}
			bullet.Register(b)
		}
		if t%10 == 0 {
			sound.PlaySound(sound.SEEnemyShot)
		}
	}

	// 自機依存による8方向発射
	if t > 210 && t < 270 && t%3 == 0 {
		px, py := player.GetPlayerPos()
		angle := math.Atan2(py-ey, px-ex)
		for i := 0; i < 8; i++ {
			b := s.bulletInfo
			b.CharID = s.charID
			b.ShotID = s.id
			b.X = ex
			b.Y = ey
			b.Color = 0
			b.Type = 7
			b.Angle = angle - math.Pi/2*0.8 + math.Pi*0.8/7*float64(i) + common.RandomAngle(math.Pi/180)
			b.Speed = 3 + common.RandomAngle(0.3)
			b.Rotate = true
			bullet.Register(b)
		}
		if t%10 == 0 {
			sound.PlaySound(sound.SEEnemyShot)
		}
	}
}
