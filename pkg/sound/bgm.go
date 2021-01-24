package sound

import (
	"fmt"
	"io/ioutil"

	"github.com/sh-miyoshi/dxlib"
	"gopkg.in/yaml.v2"
)

// Type ...
type Type int

const (
	// TypeNormal ...
	TypeNormal Type = iota
	// TypeBoss ...
	TypeBoss

	typeMax
)

type musicInfo struct {
	File      string `yaml:"file"`
	LoopCount int32  `yaml:"loopCnt"`
}

type bgmInfo struct {
	Define struct {
		Normal musicInfo `yaml:"normal"`
		Boss   musicInfo `yaml:"boss"`
	} `yaml:"bgm"`

	handles [typeMax]int32
}

var (
	bgm bgmInfo
)

// BGMStoryInit ...
func BGMStoryInit(storyFile string) error {
	// Load story data
	buf, err := ioutil.ReadFile(storyFile)
	if err != nil {
		return err
	}
	if err := yaml.Unmarshal(buf, &bgm); err != nil {
		return err
	}

	if bgm.Define.Normal.File == "" || bgm.Define.Boss.File == "" {
		// No BGM Data
		return nil
	}

	bgm.handles[TypeNormal] = dxlib.LoadSoundMem(bgm.Define.Normal.File, 3, -1)
	if bgm.handles[TypeNormal] == -1 {
		return fmt.Errorf("Failed to load BGM: %s", bgm.Define.Normal.File)
	}
	dxlib.SetLoopPosSoundMem(bgm.Define.Normal.LoopCount, bgm.handles[TypeNormal])
	dxlib.ChangeVolumeSoundMem(128, bgm.handles[TypeNormal])

	bgm.handles[TypeBoss] = dxlib.LoadSoundMem(bgm.Define.Boss.File, 3, -1)
	if bgm.handles[TypeBoss] == -1 {
		return fmt.Errorf("Failed to load BGM: %s", bgm.Define.Boss.File)
	}
	dxlib.SetLoopPosSoundMem(bgm.Define.Boss.LoopCount, bgm.handles[TypeBoss])
	dxlib.ChangeVolumeSoundMem(128, bgm.handles[TypeBoss])

	return nil
}

// BGMStoryEnd ...
func BGMStoryEnd() {
	// TODO remove resources
}

// BGMPlay ...
func BGMPlay(bgmType Type) {
	BGMStopAll()

	dxlib.PlaySoundMem(bgm.handles[bgmType], dxlib.DX_PLAYTYPE_LOOP, dxlib.TRUE)
}

// BGMStopAll ...
func BGMStopAll() {
	for _, h := range bgm.handles {
		dxlib.StopSoundMem(h)
	}
}
