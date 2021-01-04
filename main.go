package main

import (
	"fmt"
	"os"

	"github.com/sh-miyoshi/dxlib"
	"github.com/sh-miyoshi/ryuzinroku/pkg/background"
	"github.com/sh-miyoshi/ryuzinroku/pkg/bullet"
	"github.com/sh-miyoshi/ryuzinroku/pkg/common"
	"github.com/sh-miyoshi/ryuzinroku/pkg/effect"
	"github.com/sh-miyoshi/ryuzinroku/pkg/enemy"
	"github.com/sh-miyoshi/ryuzinroku/pkg/inputs"
	"github.com/sh-miyoshi/ryuzinroku/pkg/player"
	"github.com/sh-miyoshi/ryuzinroku/pkg/sound"
)

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

	board, err := background.NewBoard(
		"data/image/background/board_top.png",
		"data/image/background/board_bottom.png",
		"data/image/background/board_left.png",
		"data/image/background/board_right.png",
	)
	if err != nil {
		fmt.Printf("Failed to init board: %v\n", err)
		os.Exit(1)
	}
	back, err := background.NewBack("data/image/background/back0.png")
	if err != nil {
		fmt.Printf("Failed to init back: %v\n", err)
		os.Exit(1)
	}

	if err := enemy.StoryInit("data/story/story.yaml"); err != nil {
		fmt.Printf("Failed to init enemy: %v\n", err)
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

	count := 0
	for dxlib.ScreenFlip() == 0 && dxlib.ProcessMessage() == 0 && dxlib.ClearDrawScreen() == 0 {
		// 処理関係
		inputs.KeyStateUpdate()
		player.MgrProcess()
		enemy.MgrProcess()
		bullet.MgrProcess()
		effect.MgrProcess()

		// 描画関係
		back.Draw(count)
		player.MgrDraw()
		bullet.MgrDraw()
		enemy.MgrDraw()
		board.Draw()
		effect.MgrDraw()

		if dxlib.CheckHitKey(dxlib.KEY_INPUT_ESCAPE) == 1 {
			break
		}
		count++
	}

	dxlib.DxLib_End()
}
