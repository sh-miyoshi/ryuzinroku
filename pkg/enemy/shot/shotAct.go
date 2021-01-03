package shot

import (
	"math"

	"github.com/sh-miyoshi/ryuzinroku/pkg/bullet"
	"github.com/sh-miyoshi/ryuzinroku/pkg/common"
	"github.com/sh-miyoshi/ryuzinroku/pkg/player"
	"github.com/sh-miyoshi/ryuzinroku/pkg/sound"
)

// 1発だけ、自機に向かって直線移動
func shotAct0(ex, ey float64, s *Shot) {
	if s.count == 0 {
		// register bullet
		b := s.bulletInfo
		b.ShotID = s.id
		b.X = ex
		b.Y = ey
		px, py := player.GetPlayerPos()
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
func shotAct1(ex, ey float64, s *Shot) {
	if s.count <= 100 && s.count%10 == 0 { //100カウント中10カウントに1回
		b := s.bulletInfo
		b.ShotID = s.id
		b.X = ex
		b.Y = ey
		px, py := player.GetPlayerPos()
		b.Angle = math.Atan2(py-b.Y, px-b.X)
		b.Speed = 3
		bullet.Register(b)
	}
}

// 100カウント中に10発、自機に向かって直線発射(角度記憶)
func shotAct2(ex, ey float64, s *Shot) {
	if s.count <= 100 && s.count%10 == 0 {
		if s.count == 0 {
			px, py := player.GetPlayerPos()
			s.baseAngle = math.Atan2(py-ey, px-ex)
		}
		b := s.bulletInfo
		b.ShotID = s.id
		b.X = ex
		b.Y = ey
		b.Angle = s.baseAngle
		b.Speed = 3
		bullet.Register(b)
	}
}

// 100カウント中に10発、自機に向かってスピード変化直線発射
func shotAct3(ex, ey float64, s *Shot) {
	if s.count <= 100 && s.count%10 == 0 {
		b := s.bulletInfo
		b.ShotID = s.id
		b.X = ex
		b.Y = ey
		px, py := player.GetPlayerPos()
		b.Angle = math.Atan2(py-b.Y, px-b.X)
		b.Speed = 1 + 5.0/100*float64(s.count)
		bullet.Register(b)
	}
}

// 0.5秒に1回ずつ円形発射
func shotAct4(ex, ey float64, s *Shot) {
	if s.count < 120 && s.count%20 == 0 {
		px, py := player.GetPlayerPos()
		angle := math.Atan2(py-ey, px-ex)
		for i := 0; i < 20; i++ {
			b := s.bulletInfo
			b.ShotID = s.id
			b.X = ex
			b.Y = ey
			b.Angle = angle + math.Pi*2/20*float64(i)
			b.Speed = 4
			bullet.Register(b)
		}
	}
}

// ばらまきショット
func shotAct5(ex, ey float64, s *Shot) {
	if s.count < 120 && s.count%2 == 0 {
		b := s.bulletInfo
		b.ShotID = s.id
		b.X = ex
		b.Y = ey
		px, py := player.GetPlayerPos()
		b.Angle = math.Atan2(py-b.Y, px-b.X) + common.RandomAngle(math.Pi/4)
		b.Speed = float64(3) + common.RandomAngle(1.5)
		bullet.Register(b)
	}
}

// ばらまきショット(減速)
func shotAct6(ex, ey float64, s *Shot) {
	if s.count < 120 && s.count%2 == 0 {
		b := s.bulletInfo
		b.ShotID = s.id
		b.X = ex
		b.Y = ey
		px, py := player.GetPlayerPos()
		b.Angle = math.Atan2(py-b.Y, px-b.X) + common.RandomAngle(math.Pi/4)
		b.Speed = float64(4) + common.RandomAngle(2)
		b.ActType = 1
		bullet.Register(b)
	}
}
