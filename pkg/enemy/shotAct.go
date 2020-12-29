package enemy

import (
	"math"

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
		b.x = float64(e.X)
		b.y = float64(e.Y)
		px, py := player.GetPlayerPos()
		b.angle = math.Atan2(float64(py)-b.y, float64(px)-b.x)
		b.speed = 3
		s.bullets = append(s.bullets, &b)
	}

	if !s.finished && len(s.bullets) == 0 {
		s.finished = true
	}
}
