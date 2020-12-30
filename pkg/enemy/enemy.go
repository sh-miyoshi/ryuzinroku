package enemy

import (
	"github.com/sh-miyoshi/dxlib"
	"github.com/sh-miyoshi/ryuzinroku/pkg/bullet"
	"github.com/sh-miyoshi/ryuzinroku/pkg/common"
	"github.com/sh-miyoshi/ryuzinroku/pkg/enemy/shot"
)

var (
	acts = []func(*enemy){act0, act1, act2, act3, act4, act5, act6, act7, act8, act9}
)

type enemyShot struct {
	Type       int           `yaml:"type"`
	StartCount int           `yaml:"startCount"`
	BulletInfo bullet.Bullet `yaml:"bullet"`
}

type enemy struct {
	ApperCount  int       `yaml:"appearCount"`
	MovePattern int       `yaml:"movePattern"`
	Type        int       `yaml:"type"`
	X           float64   `yaml:"x"`
	Y           float64   `yaml:"y"`
	HP          int       `yaml:"hp"`
	Wait        int       `yaml:"wait"`
	Shot        enemyShot `yaml:"shot"`

	images   []int32
	imgCount int
	count    int
	dead     bool
	direct   common.Direct
	vx, vy   float64
	shotProc *shot.Shot
}

// Process ...
func (e *enemy) Process() {
	acts[e.MovePattern](e)

	e.count++
	e.imgCount = (e.count / 6) % 3
	switch e.direct {
	case common.DirectFront:
		e.imgCount += 3
	case common.DirectRight:
		e.imgCount += 6
	}

	e.X += e.vx
	e.Y += e.vy

	if e.HP <= 0 {
		e.dead = true
		return
	}
	if e.X < -50 || e.X > common.FiledSizeX+50 || e.Y < -50 || e.Y > common.FiledSizeY+50 {
		e.dead = true
		return
	}

	// Shot Process
	if e.count == e.Shot.StartCount {
		e.shotProc = shot.New(e.Shot.Type, e.Shot.BulletInfo)
	}

	if e.shotProc != nil {
		if e.shotProc.Process(e.X, e.Y) {
			e.shotProc = nil
		}
	}
}

// Draw ...
func (e *enemy) Draw() {
	common.CharDraw(e.X, e.Y, e.images[e.imgCount], dxlib.TRUE)
}
