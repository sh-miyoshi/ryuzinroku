package main

import (
	"fmt"
	"os"

	"github.com/sh-miyoshi/dxlib"
	"github.com/sh-miyoshi/ryuzinroku/pkg/background"
	"github.com/sh-miyoshi/ryuzinroku/pkg/common"
	"github.com/sh-miyoshi/ryuzinroku/pkg/enemy"
	"github.com/sh-miyoshi/ryuzinroku/pkg/inputs"
	"github.com/sh-miyoshi/ryuzinroku/pkg/player"
)

func main() {
	dxlib.Init("DxLib.dll")

	dxlib.ChangeWindowMode(dxlib.TRUE)
	dxlib.SetGraphMode(common.ScreenX, common.ScreenY, 16)
	dxlib.SetOutApplicationLogValidFlag(dxlib.TRUE)

	dxlib.DxLib_Init()
	dxlib.SetDrawScreen(dxlib.DX_SCREEN_BACK)

	board, err := background.NewBoard(
		"data/image/background/board_top.png",
		"data/image/background/board_bottom.png",
		"data/image/background/board_left.png",
		"data/image/background/board_right.png",
	)
	if err != nil {
		fmt.Printf("Failed to init back ground: %v\n", err)
		os.Exit(1)
	}
	if err := enemy.StoryInit("data/story/story.yaml"); err != nil {
		fmt.Printf("Failed to init enemy: %v\n", err)
		os.Exit(1)
	}
	if err := enemy.ShotInit(); err != nil {
		fmt.Printf("Failed to init enemy shot: %v\n", err)
		os.Exit(1)
	}
	if err := player.Init(); err != nil {
		fmt.Printf("Failed to init player: %v\n", err)
		os.Exit(1)
	}

	for dxlib.ScreenFlip() == 0 && dxlib.ProcessMessage() == 0 && dxlib.ClearDrawScreen() == 0 {
		// 処理関係
		inputs.KeyStateUpdate()
		player.MgrProcess()
		enemy.MgrProcess()

		// 描画関係
		player.MgrDraw()
		enemy.MgrDraw()
		board.Draw()

		if dxlib.CheckHitKey(dxlib.KEY_INPUT_ESCAPE) == 1 {
			break
		}
	}

	dxlib.DxLib_End()
}
