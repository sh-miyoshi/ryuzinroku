package enemy

import (
	"math"

	"github.com/sh-miyoshi/ryuzinroku/pkg/bullet"
	"github.com/sh-miyoshi/ryuzinroku/pkg/player"
)

// //1発だけ、自機に向かって直線移動
func shotAct0(s *shot) {
	if s.count == 0 {
		e, err := getEnemy(s.owner)
		if err != nil || e.dead {
			return
		}
		// register bullet
		b := s.BulletInfo
		b.ShotID = s.id
		b.X = float64(e.X)
		b.Y = float64(e.Y)
		px, py := player.GetPlayerPos()
		b.Angle = math.Atan2(float64(py)-b.Y, float64(px)-b.X)
		b.Speed = 3
		bullet.Register(b)
	}

	if !s.finished && !bullet.Exists(s.id) {
		s.finished = true
	}
}
