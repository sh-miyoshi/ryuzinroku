package shot

import (
	"math"

	"github.com/sh-miyoshi/ryuzinroku/pkg/bullet"
	"github.com/sh-miyoshi/ryuzinroku/pkg/player"
)

// //1発だけ、自機に向かって直線移動
func shotAct0(ex, ey float64, s *Shot) {
	if s.count == 0 {
		// register bullet
		b := s.bulletInfo
		b.ShotID = s.id
		b.X = float64(ex)
		b.Y = float64(ey)
		px, py := player.GetPlayerPos()
		b.Angle = math.Atan2(py-b.Y, px-b.X)
		b.Speed = 3
		bullet.Register(b)
	}

	if !s.finished && !bullet.Exists(s.id) {
		s.finished = true
	}
}
