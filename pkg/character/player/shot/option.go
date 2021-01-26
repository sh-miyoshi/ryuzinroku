package shot

import (
	"math"

	"github.com/sh-miyoshi/dxlib"
	"github.com/sh-miyoshi/ryuzinroku/pkg/bullet"
	"github.com/sh-miyoshi/ryuzinroku/pkg/common"
)

type option struct {
	img    int32
	count  int
	viewX  [2]float64
	viewY  [2]float64
	slow   bool
	px, py float64
}

func (o *option) draw() {
	for i := 0; i < 2; i++ {
		common.CharDraw(o.viewX[i], o.viewY[i], o.img, dxlib.TRUE)
	}
}

func (o *option) process(px, py float64, slow bool) {
	o.count++
	o.slow = slow
	o.px = px
	o.py = py

	const baseX = 25.0
	const baseY = 25.0
	if slow {
		for i := 0; i < 2; i++ {
			o.viewX[i] = px
			o.viewY[i] = py + baseY/2
		}
		o.viewX[0] += -baseX / 2
		o.viewX[1] += baseX / 2
	} else {
		for i := 0; i < 2; i++ {
			o.viewX[i] = px
			o.viewY[i] = py + baseY + math.Sin(math.Pi*2/150*float64(o.count))*10
		}
		o.viewX[0] += -baseX
		o.viewX[1] += baseX
	}
}

func (o *option) registerOptShot(ex, ey float64, power int) {
	num := 2
	if power >= 300 {
		num = 4
	}

	baseAng := [4]float64{-math.Pi / 2, -math.Pi / 2, -math.Pi/2 - math.Pi/4, -math.Pi/2 + math.Pi/4}
	for i := 0; i < num; i++ {
		ang := baseAng[i]
		if ex >= 0 && ey >= 0 {
			// TODO
		}
		b := bullet.Bullet{
			Type:     16,
			Angle:    ang,
			Speed:    20,
			IsPlayer: true,
			X:        o.viewX[i%2],
			Y:        o.viewY[i%2],
			Power:    10 - 7*(i/2),
		}
		bullet.Register(b)
	}
}
