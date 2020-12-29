package enemy

import (
	"fmt"
	"math"

	"github.com/sh-miyoshi/dxlib"
	"github.com/sh-miyoshi/ryuzinroku/pkg/common"
)

type bullet struct {
	Color int `yaml:"color"`
	Type  int `yaml:"type"`

	x, y  float64
	speed float64
	angle float64
}

type shot struct {
	Type       int    `yaml:"type"`
	StartCount int    `yaml:"startCount"`
	BulletInfo bullet `yaml:"bullet"`

	owner    string
	finished bool
	count    int
	bullets  []*bullet
}

var (
	shotActs   = []func(*shot){shotAct0}
	bulletImgs [][]int32
	shots      []*shot
)

// ShotInit ...
func ShotInit() error {
	bulletImgs = make([][]int32, 10)
	bulletImgs[0] = make([]int32, 5)
	if res := dxlib.LoadDivGraph("data/image/bullet/b0.png", 5, 5, 1, 76, 76, bulletImgs[0]); res == -1 {
		return fmt.Errorf("Failed to load image: data/image/bullet/b0.png")
	}

	return nil
}

func shotMgrProcess() {
	newShots := []*shot{}
	for _, s := range shots {
		shotActs[s.Type](s)
		s.count++
		if !s.finished {
			newShots = append(newShots, s)
		}

		newBullets := []*bullet{}
		for _, b := range s.bullets {
			b.x += math.Cos(b.angle) * b.speed
			b.y += math.Sin(b.angle) * b.speed

			if b.x < -50 || b.x > common.FiledSizeX+50 || b.y < -50 || b.y > common.FiledSizeY+50 {
				continue
			}
			newBullets = append(newBullets, b)
		}
		s.bullets = newBullets
	}
	shots = newShots
}

func shotMgrDraw() {
	// Show bullets
	for _, s := range shots {
		for _, b := range s.bullets {
			dxlib.DrawRotaGraph(int32(b.x)+common.FieldTopX, int32(b.y)+common.FieldTopX, 1, b.angle+math.Pi, bulletImgs[b.Type][b.Color], dxlib.TRUE, dxlib.FALSE, dxlib.FALSE)
		}
	}
}

func shotRegister(enemyID string, s shot) {
	s.owner = enemyID
	s.finished = false
	s.count = 0
	shots = append(shots, &s)
}
