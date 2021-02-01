package background

import (
	"fmt"

	"github.com/sh-miyoshi/dxlib"
	"github.com/sh-miyoshi/ryuzinroku/pkg/common"
)

type addImg struct {
	x, y     int32
	filePath string
	image    int32
}

// Back ...
type Back struct {
	image      int32
	imageSizeY int32
	addImgs    []addImg
}

func newBack(filePath string, addImgs ...addImg) (*Back, error) {
	res := &Back{}
	res.image = dxlib.LoadGraph(filePath)
	if res.image == -1 {
		return nil, fmt.Errorf("Failed to load image %s", filePath)
	}
	dxlib.GetGraphSize(res.image, nil, &res.imageSizeY)
	if res.imageSizeY <= 0 {
		return nil, fmt.Errorf("Failed to get graph size")
	}

	for _, i := range addImgs {
		i.image = dxlib.LoadGraph(i.filePath)
		if i.image == -1 {
			return nil, fmt.Errorf("Failed to load image %s", i.filePath)
		}
		res.addImgs = append(res.addImgs, i)
	}

	return res, nil
}

// Draw ...
func (b *Back) Draw(count int) {
	dxlib.SetDrawArea(common.FieldTopX, common.FieldTopY, common.FieldTopX+common.FiledSizeX, common.FieldTopY+common.FiledSizeY)
	y := common.FieldTopY + count%int(b.imageSizeY)
	dxlib.DrawGraph(common.FieldTopX, int32(y)-b.imageSizeY, b.image, dxlib.FALSE)
	dxlib.DrawGraph(common.FieldTopX, int32(y), b.image, dxlib.FALSE)
	for _, i := range b.addImgs {
		dxlib.DrawGraph(common.FieldTopX+i.x, common.FieldTopY+i.y, i.image, dxlib.TRUE)
	}
	dxlib.SetDrawArea(0, 0, common.ScreenX, common.ScreenY)
}
