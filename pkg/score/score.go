package score

import (
	"fmt"

	"github.com/sh-miyoshi/dxlib"
)

// Type ...
type Type int

const (
	// TypeHighScore ...
	TypeHighScore Type = iota
	// TypeScore ...
	TypeScore
	// TypePlayerPower ...
	TypePlayerPower
	// TypeMoney ...
	TypeMoney
	// TypeRemainNum ...
	TypeRemainNum

	typeMax
)

var (
	numImgs   []int32
	starImg   int32
	scoreData [typeMax]int
)

// Init ...
func Init() error {
	// load num imgs
	numImgs = make([]int32, 10)
	fname := "data/image/etc/num0.png"
	if res := dxlib.LoadDivGraph(fname, 10, 10, 1, 16, 18, numImgs, dxlib.FALSE); res == -1 {
		return fmt.Errorf("Failed to load number image: %s", fname)
	}

	fname = "data/image/etc/star.png"
	starImg = dxlib.LoadGraph(fname, dxlib.FALSE)
	if starImg == -1 {
		return fmt.Errorf("Failed to load image: %s", fname)
	}

	return nil
}

// Set ...
func Set(typ Type, value int) {
	scoreData[typ] = value
}

// Get ...
func Get(typ Type) int {
	return scoreData[typ]
}

// Draw ...
func Draw() {
	// スコア・ハイスコア表示
	high := scoreData[TypeHighScore]
	score := scoreData[TypeScore]
	for i := int32(0); i < 9; i++ {
		dxlib.DrawRotaGraphFast(625-15*i, 30, 1, 0, numImgs[high%10], dxlib.TRUE, dxlib.FALSE, dxlib.FALSE)
		dxlib.DrawRotaGraphFast(625-15*i, 50, 1, 0, numImgs[score%10], dxlib.TRUE, dxlib.FALSE, dxlib.FALSE)
		high /= 10
		score /= 10
	}

	// 残機数表示
	for i := 0; i < scoreData[TypeRemainNum]; i++ {
		dxlib.DrawGraph(499+12*int32(i), 63, starImg, dxlib.TRUE)
	}

	// パワー表示
	power := scoreData[TypePlayerPower]
	dxlib.DrawRotaGraph(547, 91, 0.9, 0, numImgs[power%10], dxlib.TRUE, dxlib.FALSE, dxlib.FALSE)
	power /= 10
	dxlib.DrawRotaGraph(536, 91, 0.9, 0, numImgs[power%10], dxlib.TRUE, dxlib.FALSE, dxlib.FALSE)
	power /= 10
	dxlib.DrawRotaGraph(513, 91, 1.0, 0, numImgs[power%10], dxlib.TRUE, dxlib.FALSE, dxlib.FALSE)
	dxlib.DrawString(522, 82, ".", 0xffffff, 0)

	// お金表示
	money := scoreData[TypeMoney]
	for i := int32(0); i < 6; i++ {
		dxlib.DrawRotaGraph(578-14*i, 154, 1, 0, numImgs[money%10], dxlib.TRUE, dxlib.FALSE, dxlib.FALSE)
		money /= 10
	}
}
