package effect

import (
	"math"
	"math/rand"

	"github.com/sh-miyoshi/dxlib"
	"github.com/sh-miyoshi/ryuzinroku/pkg/background"
	"github.com/sh-miyoshi/ryuzinroku/pkg/common"
	"github.com/sh-miyoshi/ryuzinroku/pkg/sound"
)

// 消滅時エフェクト
func ctrlAct0(c *Controller) bool {
	if c.count%2 == 0 {
		// エフェクトの登録
		col := c.Color
		if col < 0 {
			col = rand.Intn(5)
		}

		e := effect{
			Type:       effectDead,
			BlendType:  dxlib.DX_BLENDMODE_ADD,
			BlendParam: 255,
			Color:      col,
			X:          c.X,
			Y:          c.Y,
			ExtRate:    0.5,
			Angle:      common.RandomAngle(math.Pi),
		}
		effects = append(effects, &e)
	}

	return c.count > 8
}

// ボムエフェクト
func ctrlAct1(c *Controller) bool {
	if c.count == 0 {
		sound.PlaySound(sound.SEBombStart)

		// 横線
		e1 := effect{
			Type:       effectBombTitle,
			BlendType:  dxlib.DX_BLENDMODE_ALPHA,
			BlendParam: 0,
			Color:      0,
			X:          100,
			Y:          350,
			ExtRate:    1,
			Angle:      0,
			Speed:      1,
		}
		effects = append(effects, &e1)

		// 縦線
		e2 := effect{
			Type:       effectBombTitle,
			BlendType:  dxlib.DX_BLENDMODE_ALPHA,
			BlendParam: 0,
			Color:      0,
			X:          70,
			Y:          300,
			ExtRate:    1,
			Angle:      math.Pi / 2,
			MoveAngle:  -math.Pi / 2,
			Speed:      1,
		}
		effects = append(effects, &e2)

		// キャラ
		e3 := effect{
			Type:       effectBombChar,
			BlendType:  dxlib.DX_BLENDMODE_ALPHA,
			BlendParam: 0,
			Color:      0,
			X:          260,
			Y:          300,
			ExtRate:    1,
			Angle:      0,
			MoveAngle:  -math.Pi / 2,
			Speed:      0.7,
		}
		effects = append(effects, &e3)
	}

	if c.count%10 == 0 { // 1/6秒に１回
		shotAng := [4]float64{0, math.Pi, math.Pi / 2, math.Pi * 1.5}
		// ボムエフェクトの登録
		n := c.count / 10
		if n < 4 {
			e := effect{
				Type:       effectBombMain,
				BlendType:  dxlib.DX_BLENDMODE_ALPHA,
				BlendParam: 0,
				Color:      0,
				X:          c.X,
				Y:          c.Y,
				ExtRate:    0.5,
				Angle:      common.RandomAngle(math.Pi),
				MoveAngle:  shotAng[n] - math.Pi/4,
				Speed:      13 + common.RandomAngle(2),
			}
			effects = append(effects, &e)
		}
	}

	if c.count < 40 {
		background.SetBright(int32(255 - c.count*5)) // 画面の明るさ設定(暗く)
	}
	if c.count > 90 {
		background.SetBright(int32(255 - 40*5 + (c.count-90)*5)) // 画面の明るさ設定(明るく)
	}

	if c.count > 130 {
		background.SetBright(255)
		return true
	}
	return false
}
