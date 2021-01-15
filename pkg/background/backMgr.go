package background

import "fmt"

const (
	// BackNormal ...
	BackNormal int = iota
	// BackSpellCard ...
	BackSpellCard

	backMax
)

var (
	bright      int32 = 255
	currentBack       = BackNormal
	backs       [backMax]*Back
	board       *Board
)

// Init ...
func Init() error {
	var err error
	board, err = newBoard(
		"data/image/background/board_top.png",
		"data/image/background/board_bottom.png",
		"data/image/background/board_left.png",
		"data/image/background/board_right.png",
	)
	if err != nil {
		return fmt.Errorf("Failed to init board: %v", err)
	}
	backs[BackNormal], err = newBack("data/image/background/back0.png")
	if err != nil {
		return fmt.Errorf("Failed to init back normal: %v", err)
	}

	backs[BackSpellCard], err = newBack(
		"data/image/background/back1.png",
		addImg{x: 0, y: 0, filePath: "data/image/background/back1_fixed0.png"},
	)
	if err != nil {
		return fmt.Errorf("Failed to init back spellcard: %v", err)
	}

	return nil
}

// DrawBack ...
func DrawBack(count int) {
	if backs[currentBack] != nil {
		backs[currentBack].Draw(count)
	}
}

// DrawBoard ...
func DrawBoard() {
	if board != nil {
		board.Draw()
	}
}

// SetBack ...
func SetBack(backType int) {
	currentBack = backType
}

// GetBright ...
func GetBright() int32 {
	return bright
}

// SetBright ...
func SetBright(param int32) {
	bright = param
}
