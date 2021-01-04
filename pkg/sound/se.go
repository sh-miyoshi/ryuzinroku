package sound

import (
	"fmt"

	"github.com/sh-miyoshi/dxlib"
)

// SEType ...
type SEType int32

const (
	// SEEnemyShot ...
	SEEnemyShot SEType = iota
	// SEPlayerShot ...
	SEPlayerShot
	// SEEnemyDead ...
	SEEnemyDead

	// SEMax ...
	SEMax
)

var (
	soundEffects []int32
)

// Init ...
func Init() error {
	soundEffects = make([]int32, int32(SEMax))
	soundEffects[int(SEEnemyShot)] = dxlib.LoadSoundMem("data/se/enemy_shot.wav", 3, -1)
	soundEffects[int(SEPlayerShot)] = dxlib.LoadSoundMem("data/se/cshot.wav", 3, -1)
	soundEffects[int(SEEnemyDead)] = dxlib.LoadSoundMem("data/se/enemy_death.wav", 3, -1)
	for i, s := range soundEffects {
		if s == -1 {
			return fmt.Errorf("Failed to load %d sound", i)
		}
	}
	return nil
}

// PlaySound ...
func PlaySound(typ SEType) {
	dxlib.PlaySoundMem(soundEffects[typ], dxlib.DX_PLAYTYPE_BACK, dxlib.TRUE)
}
