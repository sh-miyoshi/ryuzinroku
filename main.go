package main

import (
	"fmt"
	"os"

	"github.com/sh-miyoshi/dxlib"
	"github.com/sh-miyoshi/ryuzinroku/pkg/common"
	"github.com/sh-miyoshi/ryuzinroku/pkg/fps"
	"github.com/sh-miyoshi/ryuzinroku/pkg/player"
)

func main() {
	dxlib.Init("DxLib.dll")

	dxlib.ChangeWindowMode(dxlib.TRUE)
	dxlib.SetGraphMode(common.ScreenX, common.ScreenY, 16)
	dxlib.SetOutApplicationLogValidFlag(dxlib.TRUE)

	dxlib.DxLib_Init()
	dxlib.SetDrawScreen(dxlib.DX_SCREEN_BACK)

	plyr, err := player.New(common.ImageInfo{FileName: "data/image/char/player.png", AllNum: 12, XNum: 4, YNum: 3, XSize: 73, YSize: 73})
	if err != nil {
		fmt.Printf("Failed to init player: %v\n", err)
		os.Exit(1)
	}

	for dxlib.ScreenFlip() == 0 && dxlib.ProcessMessage() == 0 && dxlib.ClearDrawScreen() == 0 {
		// 処理関係
		plyr.Process()
		fps.Wait()

		// 描画関係
		plyr.Draw()
		fps.Draw(common.ScreenX-60, 10)

		if dxlib.CheckHitKey(dxlib.KEY_INPUT_ESCAPE) == 1 {
			break
		}
	}

	dxlib.DxLib_End()
}
