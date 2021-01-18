package laser

import (
	"fmt"
	"math"

	"github.com/sh-miyoshi/dxlib"
	"github.com/sh-miyoshi/ryuzinroku/pkg/common"
)

const (
	// ColorBlue ...
	ColorBlue int = iota
	// ColorPink ...
	ColorPink
)

// Laser ...
type Laser struct {
	RotOrigin common.Coordinates
	Width     float64
	Length    float64
	Angle     float64
	Color     int
	Act       func(l *Laser) bool
	Count     int

	isRotate    bool
	baseAngle   float64
	targetAngle float64
	targetCount int
	viewOrigin  common.Coordinates
	viewRect    [4]common.Coordinates
}

const (
	viewDist = 60
)

var (
	imgMain   []int32
	imgOrigin []int32
	lasers    []*Laser
)

// Init ...
func Init() error {
	imgMain = make([]int32, 2)
	fname := "data/image/bullet/laser_main.png"
	if res := dxlib.LoadDivGraph(fname, 2, 2, 1, 30, 460, imgMain, dxlib.FALSE); res == -1 {
		return fmt.Errorf("Failed to load image: %s", fname)
	}
	imgOrigin = make([]int32, 2)
	fname = "data/image/bullet/laser_origin.png"
	if res := dxlib.LoadDivGraph(fname, 2, 2, 1, 70, 70, imgOrigin, dxlib.FALSE); res == -1 {
		return fmt.Errorf("Failed to load image: %s", fname)
	}

	return nil
}

// Register ...
func Register(l Laser) {
	l.viewOrigin.X = l.RotOrigin.X + math.Cos(l.Angle)*viewDist
	l.viewOrigin.Y = l.RotOrigin.Y + math.Sin(l.Angle)*viewDist

	if l.Act == nil {
		panic("laser act must be set")
	}

	lasers = append(lasers, &l)
}

// MgrProcess ...
func MgrProcess() {
	newLasers := []*Laser{}
	for _, l := range lasers {
		o := l.RotOrigin

		if l.isRotate {
			if l.Count >= l.targetCount {
				l.isRotate = false
			} else {
				t := float64(l.targetCount)
				c := float64(l.Count)
				delta := 2*l.targetAngle*c/t - l.targetAngle*c*c/(t*t)

				l.Angle = l.baseAngle + delta
				l.viewOrigin.X = o.X + math.Cos(l.Angle)*viewDist
				l.viewOrigin.Y = o.Y + math.Sin(l.Angle)*viewDist
			}
		}

		// 座標変換
		v := common.Coordinates{X: o.X + viewDist, Y: o.Y}
		l.viewRect[0].X, l.viewRect[0].Y = common.Rotate(o.X, o.Y, v.X, v.Y+l.Width/2, l.Angle)
		l.viewRect[1].X, l.viewRect[1].Y = common.Rotate(o.X, o.Y, v.X, v.Y-l.Width/2, l.Angle)
		l.viewRect[2].X, l.viewRect[2].Y = common.Rotate(o.X, o.Y, v.X+l.Length, v.Y-l.Width/2, l.Angle)
		l.viewRect[3].X, l.viewRect[3].Y = common.Rotate(o.X, o.Y, v.X+l.Length, v.Y+l.Width/2, l.Angle)

		if !l.Act(l) {
			// まだ終了していないならnewLasersに追加
			newLasers = append(newLasers, l)
		}

		l.Count++
	}
	lasers = newLasers
}

// MgrDraw ...
func MgrDraw() {
	if len(lasers) == 0 {
		return
	}

	dxlib.SetDrawMode(dxlib.DX_DRAWMODE_BILINEAR)
	dxlib.SetDrawBlendMode(dxlib.DX_BLENDMODE_ADD, 255)
	for _, l := range lasers {
		fx := int32(common.FieldTopX)
		fy := int32(common.FieldTopY)
		dxlib.DrawRotaGraphFast(int32(l.viewOrigin.X)+fx, int32(l.viewOrigin.Y)+fy, 1, 0, imgOrigin[l.Color], dxlib.TRUE, dxlib.FALSE, dxlib.FALSE)
		dxlib.DrawModiGraph(
			int32(l.viewRect[0].X)+fx, int32(l.viewRect[0].Y)+fy,
			int32(l.viewRect[1].X)+fx, int32(l.viewRect[1].Y)+fy,
			int32(l.viewRect[2].X)+fx, int32(l.viewRect[2].Y)+fy,
			int32(l.viewRect[3].X)+fx, int32(l.viewRect[3].Y)+fy,
			imgMain[l.Color], dxlib.TRUE,
		)
	}
	dxlib.SetDrawBlendMode(dxlib.DX_BLENDMODE_NOBLEND, 0)
	dxlib.SetDrawMode(dxlib.DX_DRAWMODE_NEAREST)
}

// IsHit ...
func IsHit() bool {
	// TODO
	return false
}

// debug
func drawSqure(p [4]common.Coordinates) {
	dxlib.DrawTriangle(int32(p[0].X), int32(p[0].Y), int32(p[1].X), int32(p[1].Y),
		int32(p[3].X), int32(p[3].Y), 0xff0000, dxlib.TRUE)
	dxlib.DrawTriangle(int32(p[1].X), int32(p[1].Y), int32(p[3].X), int32(p[3].Y),
		int32(p[2].X), int32(p[2].Y), 0xff0000, dxlib.TRUE)
}

// SetRotate ...
func (l *Laser) SetRotate(angle float64, tm int) {
	l.baseAngle = l.Angle
	l.targetAngle = angle
	l.targetCount = tm
	l.isRotate = true
}
