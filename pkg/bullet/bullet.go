package bullet

import (
	"fmt"
	"math"

	"github.com/sh-miyoshi/dxlib"
	"github.com/sh-miyoshi/ryuzinroku/pkg/common"
)

// Bullet ...
type Bullet struct {
	Color int `yaml:"color"`
	Type  int `yaml:"type"`

	ShotID string
	X, Y   float64
	Speed  float64
	Angle  float64
}

var (
	bullets    []*Bullet
	bulletImgs [][]int32
)

// Init ...
func Init() error {
	bulletImgs = make([][]int32, 10)
	bulletImgs[0] = make([]int32, 5)
	if res := dxlib.LoadDivGraph("data/image/bullet/b0.png", 5, 5, 1, 76, 76, bulletImgs[0]); res == -1 {
		return fmt.Errorf("Failed to load image: data/image/bullet/b0.png")
	}
	bulletImgs[1] = make([]int32, 6)
	if res := dxlib.LoadDivGraph("data/image/bullet/b1.png", 6, 6, 1, 22, 22, bulletImgs[1]); res == -1 {
		return fmt.Errorf("Failed to load image: data/image/bullet/b1.png")
	}
	bulletImgs[2] = make([]int32, 10)
	if res := dxlib.LoadDivGraph("data/image/bullet/b2.png", 10, 10, 1, 5, 120, bulletImgs[2]); res == -1 {
		return fmt.Errorf("Failed to load image: data/image/bullet/b2.png")
	}
	bulletImgs[3] = make([]int32, 5)
	if res := dxlib.LoadDivGraph("data/image/bullet/b3.png", 5, 5, 1, 19, 34, bulletImgs[3]); res == -1 {
		return fmt.Errorf("Failed to load image: data/image/bullet/b3.png")
	}
	bulletImgs[4] = make([]int32, 10)
	if res := dxlib.LoadDivGraph("data/image/bullet/b4.png", 10, 10, 1, 38, 38, bulletImgs[4]); res == -1 {
		return fmt.Errorf("Failed to load image: data/image/bullet/b4.png")
	}
	bulletImgs[5] = make([]int32, 3)
	if res := dxlib.LoadDivGraph("data/image/bullet/b5.png", 3, 3, 1, 14, 16, bulletImgs[5]); res == -1 {
		return fmt.Errorf("Failed to load image: data/image/bullet/b5.png")
	}
	bulletImgs[6] = make([]int32, 3)
	if res := dxlib.LoadDivGraph("data/image/bullet/b6.png", 3, 3, 1, 14, 18, bulletImgs[6]); res == -1 {
		return fmt.Errorf("Failed to load image: data/image/bullet/b6.png")
	}
	bulletImgs[7] = make([]int32, 9)
	if res := dxlib.LoadDivGraph("data/image/bullet/b7.png", 9, 9, 1, 16, 16, bulletImgs[7]); res == -1 {
		return fmt.Errorf("Failed to load image: data/image/bullet/b7.png")
	}
	bulletImgs[8] = make([]int32, 10)
	if res := dxlib.LoadDivGraph("data/image/bullet/b8.png", 10, 10, 1, 12, 18, bulletImgs[8]); res == -1 {
		return fmt.Errorf("Failed to load image: data/image/bullet/b8.png")
	}
	bulletImgs[9] = make([]int32, 3)
	if res := dxlib.LoadDivGraph("data/image/bullet/b9.png", 3, 3, 1, 13, 19, bulletImgs[9]); res == -1 {
		return fmt.Errorf("Failed to load image: data/image/bullet/b9.png")
	}

	return nil
}

// Register ...
func Register(b Bullet) {
	bullets = append(bullets, &b)
}

// MgrProcess ...
func MgrProcess() {
	newBullets := []*Bullet{}
	for _, b := range bullets {
		b.X += math.Cos(b.Angle) * b.Speed
		b.Y += math.Sin(b.Angle) * b.Speed

		if b.X < -50 || b.X > common.FiledSizeX+50 || b.Y < -50 || b.Y > common.FiledSizeY+50 {
			continue
		}
		newBullets = append(newBullets, b)
	}
	bullets = newBullets
}

// MgrDraw ...
func MgrDraw() {
	// Show bullets
	for _, b := range bullets {
		dxlib.DrawRotaGraph(int32(b.X)+common.FieldTopX, int32(b.Y)+common.FieldTopY, 1, b.Angle+math.Pi, bulletImgs[b.Type][b.Color], dxlib.TRUE, dxlib.FALSE, dxlib.FALSE)
	}
}

// Exists ...
func Exists(shotID string) bool {
	for _, b := range bullets {
		if b.ShotID == shotID {
			return true
		}
	}
	return false
}
