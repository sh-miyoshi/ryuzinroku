package background

import (
	"fmt"

	"github.com/sh-miyoshi/dxlib"
)

// Background ...
type Background struct {
	imgTop    int32
	imgBottom int32
	imgLeft   int32
	imgRight  int32
}

// New ...
func New(imgTopFile, imtBottomFile, imgLeftFile, imgRightFile string) (*Background, error) {
	res := Background{
		imgTop:    dxlib.LoadGraph(imgTopFile),
		imgBottom: dxlib.LoadGraph(imtBottomFile),
		imgLeft:   dxlib.LoadGraph(imgLeftFile),
		imgRight:  dxlib.LoadGraph(imgRightFile),
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
func (b *Background) Draw() {
	dxlib.DrawGraph(0, 0, b.imgTop, dxlib.FALSE)
	dxlib.DrawGraph(0, 16, b.imgLeft, dxlib.FALSE)
	dxlib.DrawGraph(0, 464, b.imgBottom, dxlib.FALSE)
	dxlib.DrawGraph(416, 0, b.imgRight, dxlib.FALSE)
}
