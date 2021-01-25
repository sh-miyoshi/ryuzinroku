package end

import (
	"fmt"
	"math"

	"github.com/sh-miyoshi/dxlib"
	"github.com/sh-miyoshi/ryuzinroku/pkg/character/player"
	"github.com/sh-miyoshi/ryuzinroku/pkg/inputs"
	"github.com/sh-miyoshi/ryuzinroku/pkg/score"
)

var (
	isWin    bool
	boardImg int32
	clearImg int32
	count    int
	font     int32
)

// Init ...
func Init(win bool) error {
	isWin = win

	fname := "data/image/background/result_board.png"
	boardImg = dxlib.LoadGraph(fname, dxlib.FALSE)
	if boardImg == -1 {
		return fmt.Errorf("Failed to load image: %s", fname)
	}

	fname = "data/image/etc/clear.png"
	clearImg = dxlib.LoadGraph(fname, dxlib.FALSE)
	if boardImg == -1 {
		return fmt.Errorf("Failed to load image: %s", fname)
	}

	font = dxlib.CreateFontToHandle("", 20, -1, -1, -1, -1, dxlib.FALSE, -1)

	// background.SetBright(180)
	count = 0
	return nil
}

// Draw ...
func Draw() {
	dxlib.SetDrawBlendMode(dxlib.DX_BLENDMODE_ALPHA, 140)
	dxlib.DrawGraph(50, 50, boardImg, dxlib.TRUE)
	dxlib.SetDrawBlendMode(dxlib.DX_BLENDMODE_NOBLEND, 0)

	if count > 140 {
		dxlib.DrawFormatStringToHandle(160, 80, 0xf0f0f0, font, "残機数   %d", score.Get(score.TypeRemainNum))
	}
	if count > 170 {
		dxlib.DrawFormatStringToHandle(160, 110, 0xf0f0f0, font, "被弾数   %d", player.InitRemainNum-score.Get(score.TypeRemainNum))
	}
	if count > 200 {
		dxlib.DrawFormatStringToHandle(160, 140, 0xf0f0f0, font, "スコア   %d", score.Get(score.TypeScore))
	}
	if count > 230 {
		dxlib.DrawFormatStringToHandle(160, 170, 0xf0f0f0, font, "お金     %d", score.Get(score.TypeMoney))
	}

	if count > 300 {
		if isWin {
			dxlib.SetDrawMode(dxlib.DX_DRAWMODE_BILINEAR)
			dxlib.DrawRotaGraph(300, 260, 1, -math.Pi/20, clearImg, dxlib.TRUE, dxlib.FALSE, dxlib.FALSE)
			dxlib.SetDrawMode(dxlib.DX_DRAWMODE_NEAREST)
		} else {
			dxlib.DrawBox(140, 240, 320, 280, 0x2f2f2f, dxlib.TRUE)
			dxlib.DrawStringToHandle(160, 250, "Game Over ...", 0xff0000, font, 0, dxlib.FALSE)
		}
	}
}

// Process ...
func Process() bool {
	count++

	if count > 320 {
		if inputs.CheckKey(dxlib.KEY_INPUT_Z) >= 1 {
			return true
		}
	}

	return false
}
