package bullet

import (
	"fmt"
	"math"
	"sort"

	"github.com/sh-miyoshi/dxlib"
	"github.com/sh-miyoshi/ryuzinroku/pkg/common"
)

const (
	maxImgSize = 120
)

// Bullet ...
type Bullet struct {
	Color int `yaml:"color"`
	Type  int `yaml:"type"`

	ShotID   string
	X, Y     float64
	Speed    float64
	Angle    float64
	ActType  int
	IsPlayer bool
	Power    int
	HitRange float64
}

var (
	bullets    []*Bullet
	bulletImgs [][]int32
	bulletActs = []func(*Bullet){bulletAct0, bulletAct1}
	hitRanges  = []float64{17, 4, 2.5, 2, 2, 3.5, 2, 2.5, 1.5, 2, 6}
)

// Init ...
func Init() error {
	bulletImgs = make([][]int32, 11)
	bulletImgs[0] = make([]int32, 5)
	if res := dxlib.LoadDivGraph("data/image/bullet/b0.png", 5, 5, 1, 76, 76, bulletImgs[0], dxlib.FALSE); res == -1 {
		return fmt.Errorf("Failed to load image: data/image/bullet/b0.png")
	}
	bulletImgs[1] = make([]int32, 6)
	if res := dxlib.LoadDivGraph("data/image/bullet/b1.png", 6, 6, 1, 22, 22, bulletImgs[1], dxlib.FALSE); res == -1 {
		return fmt.Errorf("Failed to load image: data/image/bullet/b1.png")
	}
	bulletImgs[2] = make([]int32, 10)
	if res := dxlib.LoadDivGraph("data/image/bullet/b2.png", 10, 10, 1, 5, 120, bulletImgs[2], dxlib.FALSE); res == -1 {
		return fmt.Errorf("Failed to load image: data/image/bullet/b2.png")
	}
	bulletImgs[3] = make([]int32, 5)
	if res := dxlib.LoadDivGraph("data/image/bullet/b3.png", 5, 5, 1, 19, 34, bulletImgs[3], dxlib.FALSE); res == -1 {
		return fmt.Errorf("Failed to load image: data/image/bullet/b3.png")
	}
	bulletImgs[4] = make([]int32, 10)
	if res := dxlib.LoadDivGraph("data/image/bullet/b4.png", 10, 10, 1, 38, 38, bulletImgs[4], dxlib.FALSE); res == -1 {
		return fmt.Errorf("Failed to load image: data/image/bullet/b4.png")
	}
	bulletImgs[5] = make([]int32, 3)
	if res := dxlib.LoadDivGraph("data/image/bullet/b5.png", 3, 3, 1, 14, 16, bulletImgs[5], dxlib.FALSE); res == -1 {
		return fmt.Errorf("Failed to load image: data/image/bullet/b5.png")
	}
	bulletImgs[6] = make([]int32, 3)
	if res := dxlib.LoadDivGraph("data/image/bullet/b6.png", 3, 3, 1, 14, 18, bulletImgs[6], dxlib.FALSE); res == -1 {
		return fmt.Errorf("Failed to load image: data/image/bullet/b6.png")
	}
	bulletImgs[7] = make([]int32, 9)
	if res := dxlib.LoadDivGraph("data/image/bullet/b7.png", 9, 9, 1, 16, 16, bulletImgs[7], dxlib.FALSE); res == -1 {
		return fmt.Errorf("Failed to load image: data/image/bullet/b7.png")
	}
	bulletImgs[8] = make([]int32, 10)
	if res := dxlib.LoadDivGraph("data/image/bullet/b8.png", 10, 10, 1, 12, 18, bulletImgs[8], dxlib.FALSE); res == -1 {
		return fmt.Errorf("Failed to load image: data/image/bullet/b8.png")
	}
	bulletImgs[9] = make([]int32, 3)
	if res := dxlib.LoadDivGraph("data/image/bullet/b9.png", 3, 3, 1, 13, 19, bulletImgs[9], dxlib.FALSE); res == -1 {
		return fmt.Errorf("Failed to load image: data/image/bullet/b9.png")
	}
	bulletImgs[10] = make([]int32, 1)
	bulletImgs[10][0] = dxlib.LoadGraph("data/image/bullet/player_b0.png", dxlib.FALSE)
	if bulletImgs[10][0] == -1 {
		return fmt.Errorf("Failed to load image: data/image/bullet/player_b0.png")
	}

	return nil
}

// Register ...
func Register(b Bullet) {
	b.HitRange = hitRanges[b.Type]
	bullets = append(bullets, &b)
}

// MgrProcess ...
func MgrProcess() {
	newBullets := []*Bullet{}
	for _, b := range bullets {
		b.X += math.Cos(b.Angle) * b.Speed
		b.Y += math.Sin(b.Angle) * b.Speed

		bulletActs[b.ActType](b)

		out := b.Speed + float64(maxImgSize)/2
		if b.X < -out || b.X > common.FiledSizeX+out || b.Y < -out || b.Y > common.FiledSizeY+out {
			continue
		}
		newBullets = append(newBullets, b)
	}
	bullets = newBullets
}

// MgrDraw ...
func MgrDraw() {
	// Show bullets
	// TODO SetDrawMode(DX_DRAWMODE_BILINEAR)
	for _, b := range bullets {
		dxlib.DrawRotaGraphFast(int32(b.X)+common.FieldTopX, int32(b.Y)+common.FieldTopY, 1, float32(b.Angle+math.Pi/2), bulletImgs[b.Type][b.Color], dxlib.TRUE, dxlib.FALSE, dxlib.FALSE)
	}
	// SetDrawMode(DX_DRAWMODE_NEAREST)
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

// GetBullets ...
func GetBullets() []*Bullet {
	return bullets
}

// RemoveHitBullets ...
func RemoveHitBullets(hits []int) {
	sort.Sort(sort.Reverse(sort.IntSlice(hits)))
	for _, hit := range hits {
		bullets[hit] = bullets[len(bullets)-1]
		bullets = bullets[:len(bullets)-1]
	}
}
