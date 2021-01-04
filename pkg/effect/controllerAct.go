package effect

import (
	"math"

	"github.com/sh-miyoshi/dxlib"
	"github.com/sh-miyoshi/ryuzinroku/pkg/common"
)

// 消滅時エフェクト
func ctrlAct0(c *Controller) bool {
	if c.count%2 == 0 {
		// エフェクトの登録
		e := effect{
			Type:       effectDead,
			BlendType:  dxlib.DX_BLENDMODE_ADD,
			BlendParam: 255,
			Color:      c.Color,
			X:          c.X,
			Y:          c.Y,
			ExtRate:    0.5,
			Angle:      common.RandomAngle(math.Pi),
		}
		effects = append(effects, &e)
	}

	return c.count > 8
}
