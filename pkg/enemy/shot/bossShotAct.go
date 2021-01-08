package shot

import (
	"math"

	"github.com/sh-miyoshi/ryuzinroku/pkg/bullet"
	"github.com/sh-miyoshi/ryuzinroku/pkg/player"
	"github.com/sh-miyoshi/ryuzinroku/pkg/sound"
)

// 円形照射
func bossShotAct0(ex, ey float64, s *Shot) {
	const tm = 120
	t := s.count % tm

	if t < 60 && t%10 == 0 {
		px, py := player.GetPlayerPos()
		angle := math.Atan2(py-ey, px-ex)
		for i := 0; i < 30; i++ {
			b := s.bulletInfo
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
