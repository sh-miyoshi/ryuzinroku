package background

import "fmt"

const (
	// BackNormal ...
	BackNormal int = iota
	// BackSpellCard ...
	BackSpellCard
)

var (
	back  *Back
	board *Board
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
	back, err = newBack("data/image/background/back0.png")
	if err != nil {
		return fmt.Errorf("Failed to init back: %v", err)
	}

	return nil
}

// DrawBack ...
func DrawBack(count int) {
	if back != nil {
		back.Draw(count)
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

}
