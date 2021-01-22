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

	Count    int
	CharID   string
	ShotID   string
	X, Y     float64
	VX, VY   float64
	Speed    float64
	Angle    float64
	IsPlayer bool
	Power    int
	HitRange float64
	Rotate   bool // 表示される弾を回転させるか
	Act      func(b *Bullet)
}

var (
	bullets    []*Bullet
	bulletImgs [][]int32
	hitRanges  = [16]float64{17, 4, 2.5, 2, 2, 3.5, 2, 2.5, 1.5, 2, 1, 2, 1.5, 4, 0.5, 6}
)

// Init ...
func Init() error {
	bulletImgs = make([][]int32, 16)
	imgs := []common.ImageInfo{
		{AllNum: 5, XNum: 5, YNum: 1, XSize: 76, YSize: 76},
		{AllNum: 6, XNum: 6, YNum: 1, XSize: 22, YSize: 22},
		{AllNum: 10, XNum: 10, YNum: 1, XSize: 5, YSize: 120},
		{AllNum: 5, XNum: 5, YNum: 1, XSize: 19, YSize: 34},
		{AllNum: 10, XNum: 10, YNum: 1, XSize: 38, YSize: 38},
		{AllNum: 3, XNum: 3, YNum: 1, XSize: 14, YSize: 16},
		{AllNum: 3, XNum: 3, YNum: 1, XSize: 14, YSize: 18},
		{AllNum: 10, XNum: 10, YNum: 1, XSize: 16, YSize: 16},
		{AllNum: 10, XNum: 10, YNum: 1, XSize: 12, YSize: 18},
		{AllNum: 3, XNum: 3, YNum: 1, XSize: 13, YSize: 19},
		{AllNum: 8, XNum: 8, YNum: 1, XSize: 8, YSize: 8},
		{AllNum: 8, XNum: 8, YNum: 1, XSize: 35, YSize: 32},
		{AllNum: 10, XNum: 10, YNum: 1, XSize: 12, YSize: 12},
		{AllNum: 10, XNum: 10, YNum: 1, XSize: 22, YSize: 22},
		{AllNum: 4, XNum: 4, YNum: 1, XSize: 6, YSize: 6},
	}

	for i, img := range imgs {
		fname := fmt.Sprintf("data/image/bullet/b%d.png", i)
		bulletImgs[i] = make([]int32, int(img.AllNum))
		if res := dxlib.LoadDivGraph(fname, img.AllNum, img.XNum, img.YNum, img.XSize, img.YSize, bulletImgs[i], dxlib.FALSE); res == -1 {
			return fmt.Errorf("Failed to load image: %s", fname)
		}
	}

	// Load player bullet
	bulletImgs[15] = make([]int32, 1)
	bulletImgs[15][0] = dxlib.LoadGraph("data/image/bullet/player_b0.png", dxlib.FALSE)
	if bulletImgs[15][0] == -1 {
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
		b.Count++
		b.X += math.Cos(b.Angle) * b.Speed
		b.Y += math.Sin(b.Angle) * b.Speed
		b.X += b.VX
		b.Y += b.VY

		if b.Act != nil {
			b.Act(b)
		}

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
	if len(bullets) == 0 {
		return
	}

	// Show bullets
	dxlib.SetDrawMode(dxlib.DX_DRAWMODE_BILINEAR)
	for _, b := range bullets {
		ang := b.Angle + math.Pi/2
		if b.Rotate {
			ang = math.Pi * 2 * float64(b.Count%120) / 120
		}

		dxlib.DrawRotaGraphFast(int32(b.X)+common.FieldTopX, int32(b.Y)+common.FieldTopY, 1, float32(ang), bulletImgs[b.Type][b.Color], dxlib.TRUE, dxlib.FALSE, dxlib.FALSE)
	}
	dxlib.SetDrawMode(dxlib.DX_DRAWMODE_NEAREST)
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

// RemoveCharBullets ...
func RemoveCharBullets(charID string) {
	newBullets := []*Bullet{}
	for _, b := range bullets {
		if b.CharID != charID {
			newBullets = append(newBullets, b)
		}
	}

	bullets = newBullets
}
