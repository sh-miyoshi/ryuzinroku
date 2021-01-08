package enemy

import (
	"fmt"
	"io/ioutil"

	"github.com/sh-miyoshi/dxlib"
	"github.com/sh-miyoshi/ryuzinroku/pkg/bullet"
	"github.com/sh-miyoshi/ryuzinroku/pkg/common"
	"github.com/sh-miyoshi/ryuzinroku/pkg/enemy/boss"
	"github.com/sh-miyoshi/ryuzinroku/pkg/enemy/minion"
	yaml "gopkg.in/yaml.v2"
)

type story struct {
	Boss    []boss.Boss     `yaml:"boss"`
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
	bossImgInfo = imageInfo{
		info: common.ImageInfo{
			FileName: "data/image/char/enemy/riria.png",
			AllNum:   8,
			XNum:     8,
			YNum:     1,
			XSize:    100,
			YSize:    100,
		},
	}

	storyInfo story
	count     int
	minions   []*minion.Minion
	bossInst  *boss.Boss
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

	// Load boss image
	bossImgInfo.images = make([]int32, int(bossImgInfo.info.AllNum))
	fname := bossImgInfo.info.FileName
	res := dxlib.LoadDivGraph(fname, bossImgInfo.info.AllNum, bossImgInfo.info.XNum, bossImgInfo.info.YNum, bossImgInfo.info.XSize, bossImgInfo.info.YSize, bossImgInfo.images, dxlib.FALSE)
	if res == -1 {
		return fmt.Errorf("Failed to load boss image: %s", fname)
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
	if bossInst != nil {
		if bossInst.Process() {
			bossInst = nil
		}
		return
	}

	bossApper()

	// Minions process
	minionApper()
	minionProc()
	count++
}

// MgrDraw ...
func MgrDraw() {
	if bossInst != nil {
		bossInst.Draw()
		return
	}

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

func minionApper() {
	for _, e := range storyInfo.Minions {
		if e.ApperCount == count {
			m := e
			minion.Init(&m, minionImgInfo[e.Type].images)
			minions = append(minions, &m)
		}
	}
}

func minionProc() {
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
}

func bossApper() {
	for _, b := range storyInfo.Boss {
		if b.AppearCount == count {
			minions = nil
			bossInst = &b
			bossInst.Init(bossImgInfo.images)
		}
	}
}
