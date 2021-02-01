package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/sh-miyoshi/dxlib"
	"github.com/sh-miyoshi/ryuzinroku/pkg/background"
	"github.com/sh-miyoshi/ryuzinroku/pkg/bullet"
	"github.com/sh-miyoshi/ryuzinroku/pkg/character"
	"github.com/sh-miyoshi/ryuzinroku/pkg/common"
	"github.com/sh-miyoshi/ryuzinroku/pkg/effect"
	"github.com/sh-miyoshi/ryuzinroku/pkg/end"
	"github.com/sh-miyoshi/ryuzinroku/pkg/inputs"
	"github.com/sh-miyoshi/ryuzinroku/pkg/item"
	"github.com/sh-miyoshi/ryuzinroku/pkg/laser"
	"github.com/sh-miyoshi/ryuzinroku/pkg/mover"
	"github.com/sh-miyoshi/ryuzinroku/pkg/score"
	"github.com/sh-miyoshi/ryuzinroku/pkg/sound"
	"github.com/sh-miyoshi/ryuzinroku/pkg/title"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	dxlib.Init("DxLib.dll")

	dxlib.ChangeWindowMode(dxlib.TRUE)
	dxlib.SetGraphMode(common.ScreenX, common.ScreenY)
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
	if err := character.Init(); err != nil {
		fmt.Printf("Failed to init character: %v\n", err)
		os.Exit(1)
	}
	if err := effect.Init(); err != nil {
		fmt.Printf("Failed to init effect: %v\n", err)
		os.Exit(1)
	}
	if err := laser.Init(); err != nil {
		fmt.Printf("Failed to init laser: %v\n", err)
		os.Exit(1)
	}
	if err := item.Init(); err != nil {
		fmt.Printf("Failed to init item: %v\n", err)
		os.Exit(1)
	}
	if err := score.Init(); err != nil {
		fmt.Printf("Failed to init score: %v\n", err)
		os.Exit(1)
	}

	// TODO set per story
	if err := character.StoryInit("data/story/story.yaml"); err != nil {
		fmt.Printf("Failed to story init character: %v\n", err)
		os.Exit(1)
	}
	if err := title.StoryInit("data/story/story.yaml"); err != nil {
		fmt.Printf("Failed to init title: %v\n", err)
		os.Exit(1)
	}
	if err := sound.BGMStoryInit("data/story/story.yaml"); err != nil {
		fmt.Printf("Failed to init bgm: %v\n", err)
		os.Exit(1)
	}

	state := 0
	count := 0

MAIN:
	for dxlib.ScreenFlip() == 0 && dxlib.ProcessMessage() == 0 && dxlib.ClearDrawScreen() == 0 {
		// 処理関係
		inputs.KeyStateUpdate()

		switch state {
		case 0:
			switch character.MgrProcess() {
			case character.ResClear:
				end.Init(true)
				state = 1
				continue
			case character.ResGameOver:
				end.Init(false)
				state = 1
				continue
			}

			bullet.MgrProcess()
			effect.MgrProcess()
			mover.Process()
			laser.MgrProcess()
			item.MgrProcess()
			title.Process()
		case 1:
			if end.Process() {
				break MAIN
			}
		}

		// 描画関係
		draw(count)

		if dxlib.CheckHitKey(dxlib.KEY_INPUT_ESCAPE) == 1 {
			break MAIN
		}
		count++

		if count == 60 {
			sound.BGMPlay(sound.TypeNormal)
		}

		// for debug
		// dxlib.DrawFormatString(0, 0, dxlib.GetColor(255, 255, 255), "%d", count)
	}

	dxlib.DxLib_End()
}

func draw(count int) {
	bright := background.GetBright()
	if bright != 255 {
		dxlib.SetDrawBright(bright, bright, bright)
	}

	background.DrawBack(count)
	item.MgrDraw()
	character.MgrDraw()
	bullet.MgrDraw()
	laser.MgrDraw()
	title.Draw()
	background.DrawBoard()
	end.Draw()

	if bright != 255 {
		dxlib.SetDrawBright(255, 255, 255)
	}

	effect.MgrDraw()
}
