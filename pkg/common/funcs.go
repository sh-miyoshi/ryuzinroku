package common

import (
	"math"
	"math/rand"

	"github.com/sh-miyoshi/dxlib"
)

// MoveOK ...
func MoveOK(x, y int) bool {
	if x >= 10 && x <= FiledSizeX-10 && y >= 5 && y <= FiledSizeY-5 {
		return true
	}
	return false
}

// CharDraw ...
func CharDraw(x float64, y float64, grHandle int32, transFlag int32) {
	dxlib.DrawRotaGraphFast(int32(x)+FieldTopX, int32(y)+FieldTopY, 1, 0, grHandle, transFlag)
}

// RandomAngle method return random value in -angle to angle
func RandomAngle(angle float64) float64 {
	return -angle + angle*2*rand.Float64()
}

// Rotate method rotates (x, y) by an angle around (bx, by).
func Rotate(bx, by, x, y, angle float64) (float64, float64) {
	x -= bx
	y -= by
	rx := math.Cos(angle)*x - math.Sin(angle)*y
	ry := math.Sin(angle)*x + math.Cos(angle)*y
	return rx + bx, ry + by
}

// UpMax ...
func UpMax(current, upVal, max int) int {
	current += upVal
	if current > max {
		return max
	}
	return current
}
