package background

import (
	"fmt"

	"github.com/sh-miyoshi/dxlib"
	"github.com/sh-miyoshi/ryuzinroku/pkg/common"
)

// Back ...
type Back struct {
	image      int32
	imageSizeY int32
}

// NewBack ...
func NewBack(filePath string) (*Back, error) {
	res := &Back{}
	res.image = dxlib.LoadGraph(filePath, dxlib.FALSE)
	if res.image == -1 {
		return nil, fmt.Errorf("Failed to load image %s", filePath)
	}
	dxlib.GetGraphSize(res.image, nil, &res.imageSizeY)
	if res.imageSizeY <= 0 {
		return nil, fmt.Errorf("Failed to get graph size")
	}

	return res, nil
}

// Draw ...
func (b *Back) Draw(count int) {
	dxlib.SetDrawArea(common.FieldTopX, common.FieldTopY, common.FieldTopX+common.FiledSizeX, common.FieldTopY+common.FiledSizeY)
	y := common.FieldTopY + count%int(b.imageSizeY)
	dxlib.DrawGraph(common.FieldTopX, int32(y)-b.imageSizeY, b.image, dxlib.FALSE)
	dxlib.DrawGraph(common.FieldTopX, int32(y), b.image, dxlib.FALSE)
	dxlib.SetDrawArea(0, 0, common.ScreenX, common.ScreenY)
}