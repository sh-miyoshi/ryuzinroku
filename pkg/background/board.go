package background

import (
	"fmt"

	"github.com/sh-miyoshi/dxlib"
)

// Board ...
type Board struct {
	imgTop    int32
	imgBottom int32
	imgLeft   int32
	imgRight  int32
}

// NewBoard ...
func NewBoard(imgTopFile, imtBottomFile, imgLeftFile, imgRightFile string) (*Board, error) {
	res := Board{
		imgTop:    dxlib.LoadGraph(imgTopFile, dxlib.FALSE),
		imgBottom: dxlib.LoadGraph(imtBottomFile, dxlib.FALSE),
		imgLeft:   dxlib.LoadGraph(imgLeftFile, dxlib.FALSE),
		imgRight:  dxlib.LoadGraph(imgRightFile, dxlib.FALSE),
	}

	if res.imgTop == -1 {
		return nil, fmt.Errorf("Failed to load top image")
	}
	if res.imgBottom == -1 {
		return nil, fmt.Errorf("Failed to load bottom image")
	}
	if res.imgLeft == -1 {
		return nil, fmt.Errorf("Failed to load left image")
	}
	if res.imgRight == -1 {
		return nil, fmt.Errorf("Failed to load right image")
	}
	return &res, nil
}

// Draw ...
func (b *Board) Draw() {
	dxlib.DrawGraph(0, 0, b.imgTop, dxlib.FALSE)
	dxlib.DrawGraph(0, 16, b.imgLeft, dxlib.FALSE)
	dxlib.DrawGraph(0, 464, b.imgBottom, dxlib.FALSE)
	dxlib.DrawGraph(416, 0, b.imgRight, dxlib.FALSE)
}
