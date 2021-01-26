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
	opt   option
}

// New ...
func New(optImg int32, initPower int) *Shot {
	return &Shot{
		Power: initPower,
		opt: option{
			img: optImg,
		},
	}
}

// Process ...
func (s *Shot) Process(px, py, ex, ey float64, slow bool) {
	s.opt.process(px, py, slow)

	if inputs.CheckKey(dxlib.KEY_INPUT_Z) > 0 {
		s.count++
		if s.count%3 == 0 {
			num := 2
			if s.Power > 200 {
				num = 4
			}

			ofsx := []float64{-10, 10, -30, 30}
			ofsy := []float64{-30, -30, -10, -10}
			for i := 0; i < num; i++ {
				if slow {
					registerBullet(px+ofsx[i]/3, py+ofsy[i]/2, s.Power)
				} else {
					registerBullet(px+ofsx[i], py+ofsy[i], s.Power)
				}
			}

			s.opt.registerOptShot(ex, ey, s.Power)
		}
	} else {
		s.count = 0
	}
}

// Draw ...
func (s *Shot) Draw() {
	if s.Power >= 100 {
		s.opt.draw()
	}
}

func registerBullet(x, y float64, power int) {
	bpw := 12 + power/100
	if power < 200 {
		bpw += 8
	}

	b := bullet.Bullet{
		Type:     15,
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
