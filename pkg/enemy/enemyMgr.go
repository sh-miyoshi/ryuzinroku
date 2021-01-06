package enemy

import (
	"fmt"
	"io/ioutil"

	"github.com/sh-miyoshi/dxlib"
	"github.com/sh-miyoshi/ryuzinroku/pkg/bullet"
	"github.com/sh-miyoshi/ryuzinroku/pkg/common"
	"github.com/sh-miyoshi/ryuzinroku/pkg/enemy/minion"
	yaml "gopkg.in/yaml.v2"
)

type story struct {
	Minions []minion.Minion `yaml:"enemies"`
}

type imageInfo struct {
	info   common.ImageInfo
	loaded bool
	images []int32
}

var (
	minionImgInfo = []*imageInfo{
		{
			info:   common.ImageInfo{AllNum: 9, XNum: 3, YNum: 3, XSize: 32, YSize: 32},
			loaded: false,
		},
	}

	storyInfo story
	minions   []*minion.Minion
	count     int
)

// StoryInit ...
func StoryInit(storyFile string) error {
	// Load story data
	buf, err := ioutil.ReadFile(storyFile)
	if err != nil {
		return err
	}
	if err := yaml.Unmarshal(buf, &storyInfo); err != nil {
		return err
	}

	// Load enemy images
	for _, e := range storyInfo.Minions {
		if e.Type >= len(minionImgInfo) {
			return fmt.Errorf("Invalid story file: enemy type %d is not defined", e.Type)
		}
		if err := load(e.Type); err != nil {
			return err
		}
	}

	count = 0
	return nil
}

// StoryEnd ...
func StoryEnd() {
	// Delete images
	for _, imgInfos := range minionImgInfo {
		if imgInfos.loaded {
			for i := 0; i < int(imgInfos.info.AllNum); i++ {
				dxlib.DeleteGraph(imgInfos.images[i])
			}
			imgInfos.images = nil
			imgInfos.loaded = false
		}
	}

	// delete enemies, storyInfo
}

// MgrProcess ...
func MgrProcess() {
	for _, e := range storyInfo.Minions {
		if e.ApperCount == count {
			m := e
			minion.Init(&m, minionImgInfo[e.Type].images)
			minions = append(minions, &m)
		}
	}

	newMinions := []*minion.Minion{}
	bullets := bullet.GetBullets()
	for _, e := range minions {
		hits := e.HitProc(bullets)
		if len(hits) > 0 {
			bullet.RemoveHitBullets(hits)
		}

		e.Process()

		if !e.IsDead() {
			newMinions = append(newMinions, e)
		}
	}
	minions = newMinions

	count++
}

// MgrDraw ...
func MgrDraw() {
	for _, e := range minions {
		e.Draw()
	}
}

func load(no int) error {
	if minionImgInfo[no].loaded {
		return nil
	}

	minionImgInfo[no].images = make([]int32, int(minionImgInfo[no].info.AllNum))
	fname := fmt.Sprintf("data/image/char/enemy/%d.png", no)
	res := dxlib.LoadDivGraph(fname, minionImgInfo[no].info.AllNum, minionImgInfo[no].info.XNum, minionImgInfo[no].info.YNum, minionImgInfo[no].info.XSize, minionImgInfo[no].info.YSize, minionImgInfo[no].images, dxlib.FALSE)
	if res == -1 {
		return fmt.Errorf("Failed to load image: %s", fname)
	}

	minionImgInfo[no].loaded = true
	return nil
}
