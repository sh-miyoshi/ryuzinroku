package shot

import (
	"math"

	"github.com/sh-miyoshi/dxlib"
	"github.com/sh-miyoshi/ryuzinroku/pkg/bullet"
	"github.com/sh-miyoshi/ryuzinroku/pkg/inputs"
	"github.com/sh-miyoshi/ryuzinroku/pkg/sound"
)

// Shot ...
type Shot struct {
	Power int

	count int
}

// Process ...
func (s *Shot) Process(px, py float64) {
	if inputs.CheckKey(dxlib.KEY_INPUT_X) > 0 {
		s.count++
		if s.count%3 == 0 {
			num := 2
			if s.Power > 200 {
				num = 4
			}

			ofsx := []float64{-10, 10, -30, 30}
			ofsy := []float64{-30, -30, -10, -10}
			for i := 0; i < num; i++ {
				registerBullet(px+ofsx[i], py+ofsy[i], s.Power)
			}
		}
	} else {
		s.count = 0
	}
}

func registerBullet(x, y float64, power int) {
	bpw := 12 + power/100
	if power < 200 {
		bpw += 8
	}

	b := bullet.Bullet{
		Type:     10,
		Angle:    -math.Pi / 2,
		Speed:    20,
		IsPlayer: true,
		X:        x,
		Y:        y,
		Power:    bpw,
	}
	bullet.Register(b)
	sound.PlaySound(sound.SEPlayerShot)
}
