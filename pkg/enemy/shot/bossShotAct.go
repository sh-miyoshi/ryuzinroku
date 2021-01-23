package shot

import (
	"math"
	"math/rand"

	"github.com/sh-miyoshi/ryuzinroku/pkg/bullet"
	"github.com/sh-miyoshi/ryuzinroku/pkg/common"
	"github.com/sh-miyoshi/ryuzinroku/pkg/laser"
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

// レーザーサンプル
func bossShotAct3(ex, ey float64, s *Shot) {
	const tm = 420
	t := s.count % tm

	if t == 0 {
		for j := 0; j < 2; j++ {
			num := 4 + s.count/tm
			for i := 0; i < num; i++ {
				fi := float64(i)
				fj := float64(j)
				fn := float64(num)
				angle := math.Pi*2/fn*fi + math.Pi*2/(fn*2)*fj + math.Pi*2/(fn*4)*float64((num+1)%2)
				l := laser.Laser{
					RotOrigin: common.Coordinates{X: ex, Y: ey},
					Angle:     angle,
					Width:     2,
					Length:    240,
					Act: func(l *laser.Laser) bool {
						if l.Count == 80 {
							l.Width = 60
							l.EnableHit = true
						}
						if l.Count >= 260 && l.Count <= 320 {
							if l.Count == 280 {
								l.EnableHit = false
							}
							l.Width = (10 * (60 - float64(l.Count-260)) / 30.0)
						}

						return l.Count >= 320
					},
				}
				if j == 0 {
					l.Color = laser.ColorBlue
					l.SetRotate(math.Pi/fn, 80)
				} else {
					l.Color = laser.ColorPink
					l.SetRotate(-math.Pi/fn, 80)
				}
				laser.Register(l)
			}
		}
		sound.PlaySound(sound.SELaser)
	}
}

// ケロちゃん風雨に負けず
func bossShotAct4(ex, ey float64, s *Shot) {
	const tm = 200
	t := s.count % tm
	if t == 0 {
		s.baseAngle = 190 + common.RandomAngle(30)
	}

	angle := math.Pi*1.5 + math.Pi/6*math.Sin(math.Pi*2/s.baseAngle*float64(s.count))
	if s.count%4 == 0 {
		for i := 0; i < 8; i++ {
			b := s.bulletInfo
			b.CharID = s.charID
			b.ShotID = s.id
			b.Type = 4
			b.Color = 0
			b.Angle = 0
			b.X = ex
			b.Y = ey
			b.VX = math.Cos(angle-math.Pi/8*4+math.Pi/8*float64(i)+math.Pi/16) * 3
			b.VY = math.Sin(angle-math.Pi/8*4+math.Pi/8*float64(i)+math.Pi/16) * 3
			b.Act = func(b *bullet.Bullet) {
				if b.Count < 150 {
					b.VY += 0.03
				}
			}
			bullet.Register(b)
		}
		sound.PlaySound(sound.SEEnemyShot)
	}
	if s.count > 80 {
		num := 1
		if t%2 == 1 {
			num = 2
		}
		for n := 0; n < num; n++ {
			angle = math.Pi*1.5 - math.Pi/2 + math.Pi/12*float64(s.count%13) + common.RandomAngle(math.Pi/15)
			b := s.bulletInfo
			b.CharID = s.charID
			b.ShotID = s.id
			b.Type = 8
			b.Color = 4
			b.Angle = 0
			b.X = ex
			b.Y = ey
			b.VX = math.Cos(angle) * 1.4 * 1.2
			b.VY = math.Sin(angle) * 1.4
			b.Act = func(b *bullet.Bullet) {
				if b.Count < 160 {
					b.VY += 0.03
				}
				b.Angle = math.Atan2(b.VY, b.VX)
			}
			bullet.Register(b)
		}
	}
}

// 反魂蝶～八部咲き～
func bossShotAct5(ex, ey float64, s *Shot) {
	const tm = 420
	t := s.count % tm

	if t == 0 {
		num := 4 + (s.count / tm)
		for j := 0; j < 2; j++ {
			for i := 0; i < num; i++ {
				fi := float64(i)
				fj := float64(j)
				fn := float64(num)
				angle := math.Pi*2/fn*fi + math.Pi*2/(fn*2)*fj + math.Pi*2/(fn*4)*float64((num+1)%2)
				l := laser.Laser{
					RotOrigin: common.Coordinates{X: ex, Y: ey},
					Angle:     angle,
					Width:     2,
					Length:    310,
					Act: func(l *laser.Laser) bool {
						if l.Count == 80 {
							l.Width = 30
							l.EnableHit = true
						}
						if l.Count >= 260 && l.Count <= 320 {
							if l.Count == 280 {
								l.EnableHit = false
							}
							l.Width = (10 * (60 - float64(l.Count-260)) / 30.0)
						}

						return l.Count >= 320
					},
				}
				if j == 0 {
					l.Color = laser.ColorBlue
					l.SetRotate(math.Pi/fn, 80)
				} else {
					l.Color = laser.ColorPink
					l.SetRotate(-math.Pi/fn, 80)
				}
				laser.Register(l)
			}
		}
		sound.PlaySound(sound.SELaser)
	}

	act0 := func(b *bullet.Bullet) {
		if b.Count > 90 && b.Count <= 100 {
			b.Speed -= b.Speed / 220
		}
	}
	act1 := func(b *bullet.Bullet) {
		if b.Count > 50 {
			b.Speed += b.Speed / 45
		}
	}
	act2 := func(b *bullet.Bullet) {
		if b.Count > 65 {
			b.Speed += b.Speed / 90
		}
	}

	if t == 50 {
		angle := common.RandomAngle(math.Pi)
		for l := 0; l < 2; l++ {
			for k := 0; k < 3; k++ {
				for j := 0; j < 3; j++ {
					for i := 0; i < 30; i++ {
						b := s.bulletInfo
						b.CharID = s.charID
						b.ShotID = s.id
						b.Type = 11
						b.Color = l
						b.Angle = angle + math.Pi*2/30*float64(i) + math.Pi*2/60*float64(l)
						b.Speed = 1.8 - 0.2*float64(j) + 0.1*float64(l)
						b.X = ex
						b.Y = ey
						switch k {
						case 0:
							b.Act = act0
						case 1:
							b.Act = act1
						case 2:
							b.Act = act2
						}
						bullet.Register(b)
					}
				}
			}
		}
		sound.PlaySound(sound.SEEnemyShot)
	}

	if t >= 170 && t < 310 && (t-170)%35 == 0 {
		angle := common.RandomAngle(math.Pi)
		for k := 0; k < 2; k++ { // 速度の違う2つの弾がある
			for j := 0; j < 3; j++ { // 1箇所から3つにわかれる
				for i := 0; i < 30; i++ { // 1周30個
					b := s.bulletInfo
					b.CharID = s.charID
					b.ShotID = s.id
					b.Type = 11
					b.Color = 2
					b.Angle = angle + math.Pi*2/30*float64(i)
					b.Speed = 2 - 0.3*float64(k)
					b.X = ex
					b.Y = ey
					b.Act = func(b *bullet.Bullet) {
						switch j {
						case 0:
							act0(b)
						case 1:
							act1(b)
						case 2:
							act2(b)
						}

						if b.Count > 15 && b.Count <= 80 {
							baseAng := math.Pi / 300
							if (t-170)%70 == 0 {
								baseAng *= -1
							}
							b.Angle += baseAng
						}
					}

					bullet.Register(b)
				}
			}
		}
		sound.PlaySound(sound.SEEnemyShot)
	}

	if t == 360 {
		angle := common.RandomAngle(math.Pi)
		for j := 0; j < 3; j++ { // 1箇所から3つにわかれる
			for i := 0; i < 30; i++ { // 1周30個
				b := s.bulletInfo
				b.CharID = s.charID
				b.ShotID = s.id
				b.Type = 0
				b.Color = 1
				b.Angle = angle + math.Pi*2/30*float64(i)
				b.Speed = 1.8
				b.X = ex
				b.Y = ey
				switch j {
				case 0:
					b.Act = act0
				case 1:
					b.Act = act1
				case 2:
					b.Act = act2
				}
				bullet.Register(b)
			}
		}
		sound.PlaySound(sound.SEEnemyShot)
	}
}
