package enemy

import (
	"fmt"
	"io/ioutil"

	"github.com/google/uuid"
	"github.com/sh-miyoshi/dxlib"
	"github.com/sh-miyoshi/ryuzinroku/pkg/common"
	yaml "gopkg.in/yaml.v2"
)

type story struct {
	Enemies []enemy `yaml:"enemies"`
}

type imageInfo struct {
	info   common.ImageInfo
	loaded bool
	images []int32
}

var (
	enemyImgInfo = []*imageInfo{
		{
			info:   common.ImageInfo{AllNum: 9, XNum: 3, YNum: 3, XSize: 32, YSize: 32},
			loaded: false,
		},
	}

	storyInfo story
	enemies   []*enemy
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
	for _, e := range storyInfo.Enemies {
		if e.Type >= len(enemyImgInfo) {
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
	for _, imgInfos := range enemyImgInfo {
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
	for _, e := range storyInfo.Enemies {
		if e.ApperCount == count {
			enemy := e
			enemy.id = uuid.New().String()
			enemy.images = enemyImgInfo[e.Type].images
			enemy.imgSizeX = enemyImgInfo[e.Type].info.XSize
			enemy.imgSizeY = enemyImgInfo[e.Type].info.YSize
			enemy.imgCount = 0
			enemy.dead = false
			enemy.direct = common.DirectFront
			enemies = append(enemies, &enemy)
		}
	}

	newEnemies := []*enemy{}
	for _, e := range enemies {
		e.Process()
		if !e.dead {
			newEnemies = append(newEnemies, e)
		}
	}
	enemies = newEnemies

	shotMgrProcess()

	count++
}

// MgrDraw ...
func MgrDraw() {
	for _, e := range enemies {
		e.Draw()
	}
}

func getEnemy(id string) (*enemy, error) {
	for _, e := range enemies {
		if e.id == id {
			return e, nil
		}
	}
	return nil, fmt.Errorf("enemy %s is not exists", id)
}

func load(no int) error {
	if enemyImgInfo[no].loaded {
		return nil
	}

	enemyImgInfo[no].images = make([]int32, int(enemyImgInfo[no].info.AllNum))
	fname := fmt.Sprintf("data/image/char/enemy/%d.png", no)
	res := dxlib.LoadDivGraph(fname, enemyImgInfo[no].info.AllNum, enemyImgInfo[no].info.XNum, enemyImgInfo[no].info.YNum, enemyImgInfo[no].info.XSize, enemyImgInfo[no].info.YSize, enemyImgInfo[no].images)
	if res == -1 {
		return fmt.Errorf("Failed to load image: %s", fname)
	}

	enemyImgInfo[no].loaded = true
	return nil
}
