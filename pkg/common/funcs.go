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
func CharDraw(x int32, y int32, imgSizeX int32, imgSizeY int32, grHandle int32, transFlag int32) {
	centerX := x - imgSizeX/2 + FieldTopX
	centerY := y - imgSizeY/2 + FieldTopY
	dxlib.DrawGraph(centerX, centerY, grHandle, transFlag)
}
