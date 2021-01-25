package shot

import (
	"math"

	"github.com/sh-miyoshi/ryuzinroku/pkg/bullet"
	"github.com/sh-miyoshi/ryuzinroku/pkg/common"
	"github.com/sh-miyoshi/ryuzinroku/pkg/sound"
)

func bulletActSlow(b *bullet.Bullet) {
	if b.Speed > 1.5 {
		b.Speed -= 0.04
	}
}

// 1発だけ、自機に向かって直線移動
func shotAct0(ex, ey, px, py float64, s *Shot) {
	if s.count == 0 {
		// register bullet
		b := s.bulletInfo
		b.CharID = s.charID
		b.ShotID = s.id
		b.X = ex
		b.Y = ey
		b.Angle = math.Atan2(py-b.Y, px-b.X)
		b.Speed = 3
		bullet.Register(b)
		sound.PlaySound(sound.SEEnemyShot)
	}

	if !s.finished && !bullet.Exists(s.id) {
		s.finished = true
	}
}

// 100カウント中に10発、自機に向かって直線発射(常に自機狙い)
func shotAct1(ex, ey, px, py float64, s *Shot) {
	if s.count <= 100 && s.count%10 == 0 { //100カウント中10カウントに1回
		b := s.bulletInfo
		b.CharID = s.charID
		b.ShotID = s.id
		b.X = ex
		b.Y = ey
		b.Angle = math.Atan2(py-b.Y, px-b.X)
		b.Speed = 3
		bullet.Register(b)
		sound.PlaySound(sound.SEEnemyShot)
	}
}

// 100カウント中に10発、自機に向かって直線発射(角度記憶)
func shotAct2(ex, ey, px, py float64, s *Shot) {
	if s.count <= 100 && s.count%10 == 0 {
		if s.count == 0 {
			s.baseAngle = math.Atan2(py-ey, px-ex)
		}
		b := s.bulletInfo
		b.CharID = s.charID
		b.ShotID = s.id
		b.X = ex
		b.Y = ey
		b.Angle = s.baseAngle
		b.Speed = 3
		bullet.Register(b)
		sound.PlaySound(sound.SEEnemyShot)
	}
}

// 100カウント中に10発、自機に向かってスピード変化直線発射
func shotAct3(ex, ey, px, py float64, s *Shot) {
	if s.count <= 100 && s.count%10 == 0 {
		b := s.bulletInfo
		b.CharID = s.charID
		b.ShotID = s.id
		b.X = ex
		b.Y = ey
		b.Angle = math.Atan2(py-b.Y, px-b.X)
		b.Speed = 1 + 5.0/100*float64(s.count)
		bullet.Register(b)
		sound.PlaySound(sound.SEEnemyShot)
	}
}

// 0.5秒に1回ずつ円形発射
func shotAct4(ex, ey, px, py float64, s *Shot) {
	if s.count < 120 && s.count%20 == 0 {
		angle := math.Atan2(py-ey, px-ex)
		for i := 0; i < 20; i++ {
			b := s.bulletInfo
			b.CharID = s.charID
			b.ShotID = s.id
			b.X = ex
			b.Y = ey
			b.Angle = angle + math.Pi*2/20*float64(i)
			b.Speed = 4
			bullet.Register(b)
			sound.PlaySound(sound.SEEnemyShot)
		}
	}
}

// ばらまきショット
func shotAct5(ex, ey, px, py float64, s *Shot) {
	if s.count < 120 && s.count%2 == 0 {
		b := s.bulletInfo
		b.CharID = s.charID
		b.ShotID = s.id
		b.X = ex
		b.Y = ey
		b.Angle = math.Atan2(py-b.Y, px-b.X) + common.RandomAngle(math.Pi/4)
		b.Speed = float64(3) + common.RandomAngle(1.5)
		bullet.Register(b)
		sound.PlaySound(sound.SEEnemyShot)
	}
}

// ばらまきショット(減速)
func shotAct6(ex, ey, px, py float64, s *Shot) {
	if s.count < 120 && s.count%2 == 0 {
		b := s.bulletInfo
		b.CharID = s.charID
		b.ShotID = s.id
		b.X = ex
		b.Y = ey
		b.Angle = math.Atan2(py-b.Y, px-b.X) + common.RandomAngle(math.Pi/4)
		b.Speed = float64(4) + common.RandomAngle(2)
		b.Act = bulletActSlow
		bullet.Register(b)
		sound.PlaySound(sound.SEEnemyShot)
	}
}

// みょん弾幕
func shotAct7(ex, ey, px, py float64, s *Shot) {
	if s.count <= 150 && s.count%10 == 0 {
		for i := 0; i < 20; i++ {
			b := s.bulletInfo
			b.CharID = s.charID
			b.Type = 7
			b.ShotID = s.id
			b.X = ex + math.Cos(math.Pi/2+math.Pi/150*float64(s.count))*100
			b.Y = ey + math.Sin(math.Pi/2+math.Pi/150*float64(s.count))*100
			b.Angle = math.Pi * 2 / 20 * float64(i)
			b.Speed = 1.2
			b.Color = 2
			b.Act = bulletActSlow
			bullet.Register(b)
		}
		for i := 0; i < 20; i++ {
			b := s.bulletInfo
			b.CharID = s.charID
			b.Type = 7
			b.ShotID = s.id
			b.X = ex + math.Cos(math.Pi/2-math.Pi/150*float64(s.count))*100
			b.Y = ey + math.Sin(math.Pi/2-math.Pi/150*float64(s.count))*100
			b.Angle = math.Pi * 2 / 20 * float64(i)
			b.Speed = 1.2
			b.Color = 4
			b.Act = bulletActSlow
			bullet.Register(b)
			sound.PlaySound(sound.SEEnemyShot)
		}
	}
}

// ミシャグジさま弾幕
func shotAct8(ex, ey, px, py float64, s *Shot) {
	if s.count < 1200 && s.count%90 == 0 {
		angle := common.RandomAngle(math.Pi)
		for j := 0; j < 2; j++ { // 途中から２分裂する分
			for i := 0; i < 60; i++ { // 一度に60個
				b := s.bulletInfo
				b.CharID = s.charID
				b.Type = 8
				b.ShotID = s.id
				b.X = ex
				b.Y = ey
				b.Angle = angle + math.Pi*2/60*float64(i)
				b.Speed = 2
				b.Color = 4
				if j == 0 {
					b.Act = func(b *bullet.Bullet) {
						if b.Count > 30 && b.Count < 120 {
							b.Speed -= 1.2 / 90
							// 90カウントかけて90°傾ける
							b.Angle += (math.Pi / 2) / 90 * (1)
						}
					}
				} else {
					b.Act = func(b *bullet.Bullet) {
						if b.Count > 30 && b.Count < 120 {
							b.Speed -= 1.2 / 90
							// 90カウントかけて90°傾ける
							b.Angle += (math.Pi / 2) / 90 * (-1)
						}
					}
				}
				bullet.Register(b)
			}
		}
		sound.PlaySound(sound.SEEnemyShot)
	}
}

// ダミー(なにもしない)
func shotActNon(ex, ey, px, py float64, s *Shot) {
	// nothing to do
}
