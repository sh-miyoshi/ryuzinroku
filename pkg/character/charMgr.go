package character

import (
	"github.com/sh-miyoshi/ryuzinroku/pkg/character/enemy"
	"github.com/sh-miyoshi/ryuzinroku/pkg/character/player"
)

const (
	// ResContinue ...
	ResContinue int = iota
	// ResClear ...
	ResClear
	// ResGameOver ...
	ResGameOver
)

// Init ...
func Init() error {
	return player.Init()
}

// StoryInit ...
func StoryInit(storyFile string) error {
	return enemy.StoryInit(storyFile)
}

// StoryEnd ...
func StoryEnd() {
	enemy.StoryEnd()
}

// MgrProcess ...
func MgrProcess() int {
	px, py := player.GetPlayerPos()
	ex, ey := enemy.GetClosestEnemy(px, py)

	if player.MgrProcess(ex, ey) {
		return ResGameOver
	}

	if enemy.MgrProcess(px, py) {
		return ResClear
	}

	return ResContinue
}

// MgrDraw ...
func MgrDraw() {
	enemy.MgrDraw()
	player.MgrDraw()
}
