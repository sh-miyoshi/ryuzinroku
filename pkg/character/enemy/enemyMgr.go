package enemy

import (
	"fmt"
	"io/ioutil"

	"github.com/sh-miyoshi/dxlib"
	"github.com/sh-miyoshi/ryuzinroku/pkg/bullet"
	"github.com/sh-miyoshi/ryuzinroku/pkg/character/enemy/boss"
	"github.com/sh-miyoshi/ryuzinroku/pkg/character/enemy/minion"
	"github.com/sh-miyoshi/ryuzinroku/pkg/common"
	"github.com/sh-miyoshi/ryuzinroku/pkg/sound"
	yaml "gopkg.in/yaml.v2"
)

type story struct {
	Boss    []boss.Define   `yaml:"boss"`
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
	bossCharImgInfo = [boss.TypeMax]imageInfo{
		{
			info: common.ImageInfo{
				FileName: "data/image/char/enemy/riria.png",
				AllNum:   8,
				XNum:     8,
				YNum:     1,
				XSize:    100,
				YSize:    100,
			},
		},
	}

	storyInfo   story
	count       int
	isFinal     bool
	minions     []*minion.Minion
	bossInst    boss.Boss
	bossHPImg   [boss.HPColMax]int32
	bossEtcImgs []int32
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
	// ボス登場時に読み込もうとすると遅延が発生するかもしれないので先に読み込んでおく
	for i, b := range bossCharImgInfo {
		bossCharImgInfo[i].images = make([]int32, int(b.info.AllNum))
		fname := b.info.FileName
		res := dxlib.LoadDivGraph(fname, b.info.AllNum, b.info.XNum, b.info.YNum, b.info.XSize, b.info.YSize, bossCharImgInfo[i].images)
		if res == -1 {
			return fmt.Errorf("Failed to load boss image: %s", fname)
		}
	}
	bossHPImg[boss.HPColNormal] = dxlib.LoadGraph("data/image/etc/boss_hp_normal.png")
	if bossHPImg[boss.HPColNormal] == -1 {
		return fmt.Errorf("Failed to load boss hp image: data/image/etc/boss_hp_normal.png")
	}
	bossHPImg[boss.HPColBright] = dxlib.LoadGraph("data/image/etc/boss_hp_bright.png")
	if bossHPImg[boss.HPColBright] == -1 {
		return fmt.Errorf("Failed to load boss hp image: data/image/etc/boss_hp_bright.png")
	}
	// etc img
	img := dxlib.LoadGraph("data/image/effect/bossback0.png")
	if img == -1 {
		return fmt.Errorf("Failed to load boss etc image: data/image/effect/bossback0.png")
	}
	bossEtcImgs = append(bossEtcImgs, img)
	img = dxlib.LoadGraph("data/image/effect/bossback1.png")
	if img == -1 {
		return fmt.Errorf("Failed to load boss etc image: data/image/effect/bossback1.png")
	}
	bossEtcImgs = append(bossEtcImgs, img)
	img = dxlib.LoadGraph("data/image/effect/bossback2.png")
	if img == -1 {
		return fmt.Errorf("Failed to load boss etc image: data/image/effect/bossback2.png")
	}
	bossEtcImgs = append(bossEtcImgs, img)

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
func MgrProcess(px, py float64) bool {
	if bossInst != nil {
		finish := false
		if bossInst.Process(px, py) {
			bossInst.Clear()
			bossInst = nil
			finish = isFinal
		}
		return finish
	}

	bossApper()

	// Minions process
	minionApper()
	minionProc(px, py)
	count++

	return false
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

// GetClosestEnemy ...
func GetClosestEnemy(x, y float64) (float64, float64) {
	if bossInst != nil {
		return bossInst.GetPos()
	}

	resX := -1.0
	resY := -1.0
	dist := float64(common.FiledSizeX*common.FiledSizeX + common.FiledSizeY*common.FiledSizeY)

	for _, e := range minions {
		tx := e.X - x
		ty := e.Y - y
		d := tx*tx + ty*ty
		if d < dist {
			resX = e.X
			resY = e.Y
		}
	}

	return resX, resY
}

func load(no int) error {
	if minionImgInfo[no].loaded {
		return nil
	}

	minionImgInfo[no].images = make([]int32, int(minionImgInfo[no].info.AllNum))
	fname := fmt.Sprintf("data/image/char/enemy/%d.png", no)
	res := dxlib.LoadDivGraph(fname, minionImgInfo[no].info.AllNum, minionImgInfo[no].info.XNum, minionImgInfo[no].info.YNum, minionImgInfo[no].info.XSize, minionImgInfo[no].info.YSize, minionImgInfo[no].images)
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

func minionProc(px, py float64) {
	newMinions := []*minion.Minion{}
	bullets := bullet.GetBullets()
	bombDm := bullet.PopBomb()
	for _, e := range minions {
		hits := e.HitProc(bullets)
		if len(hits) > 0 {
			bullet.RemoveHitBullets(hits)
		}

		if bombDm > 0 {
			e.HP -= bombDm
		}

		e.Process(px, py)

		if !e.IsDead() {
			newMinions = append(newMinions, e)
		}
	}
	minions = newMinions
}

func bossApper() error {
	for _, b := range storyInfo.Boss {
		if b.AppearCount == count {
			isFinal = b.Final
			minions = nil
			var err error
			bossInst, err = boss.NewRiria(b, bossCharImgInfo[b.Type].images, bossHPImg, bossEtcImgs)
			if isFinal {
				sound.BGMPlay(sound.TypeBoss)
			}
			return err
		}
	}
	return nil
}
