package common

import "github.com/sh-miyoshi/dxlib"

// MoveOK ...
func MoveOK(x, y int) bool {
	if x >= 10 && x <= FiledSizeX-10 && y >= 5 && y <= FiledSizeY-5 {
		return true
	}
	return false
}

// CharDraw ...
func CharDraw(x float64, y float64, grHandle int32, transFlag int32) {
	dxlib.DrawRotaGraph(int32(x)+FieldTopX, int32(y)+FieldTopY, 1, 0, grHandle, transFlag, dxlib.FALSE, dxlib.FALSE)
}
