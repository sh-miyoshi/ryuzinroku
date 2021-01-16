package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/sh-miyoshi/dxlib"
	"github.com/sh-miyoshi/ryuzinroku/pkg/background"
	"github.com/sh-miyoshi/ryuzinroku/pkg/bullet"
	"github.com/sh-miyoshi/ryuzinroku/pkg/common"
	"github.com/sh-miyoshi/ryuzinroku/pkg/effect"
	"github.com/sh-miyoshi/ryuzinroku/pkg/enemy"
	"github.com/sh-miyoshi/ryuzinroku/pkg/inputs"
	"github.com/sh-miyoshi/ryuzinroku/pkg/mover"
	"github.com/sh-miyoshi/ryuzinroku/pkg/player"
	"github.com/sh-miyoshi/ryuzinroku/pkg/sound"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	dxlib.Init("DxLib.dll")

	dxlib.ChangeWindowMode(dxlib.TRUE)
	dxlib.SetGraphMode(common.ScreenX, common.ScreenY, 16, 60)
	dxlib.SetOutApplicationLogValidFlag(dxlib.TRUE)

	dxlib.DxLib_Init()
	dxlib.SetDrawScreen(dxlib.DX_SCREEN_BACK)

	if err := sound.Init(); err != nil {
		fmt.Printf("Failed to init sound: %v\n", err)
		os.Exit(1)
	}
	if err := background.Init(); err != nil {
		fmt.Printf("Failed to init background: %v\n", err)
		os.Exit(1)
	}
	if err := bullet.Init(); err != nil {
		fmt.Printf("Failed to init bullet: %v\n", err)
		os.Exit(1)
	}
	if err := player.Init(); err != nil {
		fmt.Printf("Failed to init player: %v\n", err)
		os.Exit(1)
	}
	if err := effect.Init(); err != nil {
		fmt.Printf("Failed to init effect: %v\n", err)
		os.Exit(1)
	}

	// TODO set per story
	if err := enemy.StoryInit("data/story/story.yaml"); err != nil {
		fmt.Printf("Failed to init enemy: %v\n", err)
		os.Exit(1)
	}

	count := 0
	for dxlib.ScreenFlip() == 0 && dxlib.ProcessMessage() == 0 && dxlib.ClearDrawScreen() == 0 {
		// 処理関係
		inputs.KeyStateUpdate()
		player.MgrProcess()
		enemy.MgrProcess()
		bullet.MgrProcess()
		effect.MgrProcess()
		mover.Process()

		// 描画関係
		draw(count)

		if dxlib.CheckHitKey(dxlib.KEY_INPUT_ESCAPE) == 1 {
			break
		}
		count++
		dxlib.DrawFormatString(0, 0, dxlib.GetColor(255, 255, 255), "%d", count)
	}

	dxlib.DxLib_End()
}

func draw(count int) {
	bright := background.GetBright()
	if bright != 255 {
		dxlib.SetDrawBright(bright, bright, bright)
	}
	background.DrawBack(count)
	enemy.MgrDraw()
	if bright != 255 {
		dxlib.SetDrawBright(255, 255, 255)
	}

	player.MgrDraw()

	if bright != 255 {
		dxlib.SetDrawBright(bright, bright, bright)
	}

	bullet.MgrDraw()
	background.DrawBoard()

	if bright != 255 {
		dxlib.SetDrawBright(255, 255, 255)
	}

	effect.MgrDraw()
}
