package title

import (
	"fmt"
	"io/ioutil"

	"github.com/sh-miyoshi/dxlib"
	"github.com/sh-miyoshi/ryuzinroku/pkg/common"
	"gopkg.in/yaml.v2"
)

type titleInfo struct {
	Title struct {
		StartTime int    `yaml:"startCount"`
		ImageFile string `yaml:"file"`
	} `yaml:"title"`

	enable     bool
	image      int32
	count      int
	blendParam int32
}

var (
	title titleInfo
)

// StoryInit ...
func StoryInit(storyFile string) error {
	// Load story data
	buf, err := ioutil.ReadFile(storyFile)
	if err != nil {
		return err
	}
	if err := yaml.Unmarshal(buf, &title); err != nil {
		return err
	}

	title.enable = false

	if title.Title.ImageFile == "" {
		// No title data
		return nil
	}

	title.image = dxlib.LoadGraph(title.Title.ImageFile)
	if title.image == -1 {
		return fmt.Errorf("Failed to load title image: %s", title.Title.ImageFile)
	}

	return nil
}

// StoryEnd ...
func StoryEnd() {
	// TODO remove resources
}

// Process ...
func Process() {
	title.count++

	if title.count == title.Title.StartTime {
		title.enable = true
	}

	if !title.enable {
		return
	}

	cnt := title.count - title.Title.StartTime
	if cnt < 128 {
		title.blendParam += 2
	}
	if cnt > 128+128 {
		title.blendParam -= 2
	}
	if cnt > 128+128+128 {
		// 終了
		title.enable = false
	}
}

// Draw ...
func Draw() {
	if title.enable {
		dxlib.SetDrawBlendMode(dxlib.DX_BLENDMODE_ALPHA, title.blendParam)
		dxlib.DrawGraph(120+common.FieldTopX, 10+common.FieldTopY, title.image, dxlib.TRUE)
		dxlib.SetDrawBlendMode(dxlib.DX_BLENDMODE_NOBLEND, 0)
	}
}
