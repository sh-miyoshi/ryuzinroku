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
	// SEEnemyHit ...
	SEEnemyHit
	// SEPlayerDead ...
	SEPlayerDead
	// SEBombStart ...
	SEBombStart
	// SEBomb ...
	SEBomb
	// SELaser ...
	SELaser
	// SEBossDead ...
	SEBossDead
	// SEItemGet ...
	SEItemGet
	// SEBossChange ...
	SEBossChange

	// SEMax ...
	SEMax
)

var (
	soundEffects []int32
)

// Init ...
func Init() error {
	soundEffects = make([]int32, int32(SEMax))
	soundEffects[SEEnemyShot] = dxlib.LoadSoundMem("data/se/enemy_shot.wav")
	soundEffects[SEPlayerShot] = dxlib.LoadSoundMem("data/se/cshot.wav")
	soundEffects[SEEnemyDead] = dxlib.LoadSoundMem("data/se/enemy_death.wav")
	soundEffects[SEEnemyHit] = dxlib.LoadSoundMem("data/se/hit.wav")
	soundEffects[SEPlayerDead] = dxlib.LoadSoundMem("data/se/char_death.wav")
	soundEffects[SEBombStart] = dxlib.LoadSoundMem("data/se/bom0.wav")
	soundEffects[SEBomb] = dxlib.LoadSoundMem("data/se/bom1.wav")
	soundEffects[SELaser] = dxlib.LoadSoundMem("data/se/lazer.wav")
	soundEffects[SEBossDead] = dxlib.LoadSoundMem("data/se/boss_death.wav")
	soundEffects[SEItemGet] = dxlib.LoadSoundMem("data/se/item_get.wav")
	soundEffects[SEBossChange] = dxlib.LoadSoundMem("data/se/boss_change.wav")

	for i, s := range soundEffects {
		if s == -1 {
			return fmt.Errorf("Failed to load %d sound", i)
		}
	}

	// 各素材の再生ボリュームを設定
	dxlib.ChangeVolumeSoundMem(128, soundEffects[SEEnemyShot])
	dxlib.ChangeVolumeSoundMem(96, soundEffects[SEEnemyDead])
	dxlib.ChangeVolumeSoundMem(128, soundEffects[SEPlayerShot])
	dxlib.ChangeVolumeSoundMem(64, soundEffects[SEEnemyHit])
	dxlib.ChangeVolumeSoundMem(64, soundEffects[SEPlayerDead])
	dxlib.ChangeVolumeSoundMem(64, soundEffects[SEBombStart])
	dxlib.ChangeVolumeSoundMem(64, soundEffects[SEBomb])
	dxlib.ChangeVolumeSoundMem(64, soundEffects[SELaser])
	dxlib.ChangeVolumeSoundMem(128, soundEffects[SEBossDead])
	dxlib.ChangeVolumeSoundMem(32, soundEffects[SEItemGet])
	dxlib.ChangeVolumeSoundMem(96, soundEffects[SEBossChange])

	return nil
}

// PlaySound ...
func PlaySound(typ SEType) {
	if dxlib.CheckSoundMem(soundEffects[typ]) == 1 {
		if typ == SEEnemyHit {
			return
		}
		dxlib.StopSoundMem(soundEffects[typ])
	}
	dxlib.PlaySoundMem(soundEffects[typ], dxlib.DX_PLAYTYPE_BACK, dxlib.TRUE)
}
