package enemy

import (
	"github.com/google/uuid"
	"github.com/sh-miyoshi/ryuzinroku/pkg/bullet"
)

type shot struct {
	Type       int           `yaml:"type"`
	StartCount int           `yaml:"startCount"`
	BulletInfo bullet.Bullet `yaml:"bullet"`

	id       string
	owner    string
	finished bool
	count    int
}

var (
	shotActs = []func(*shot){shotAct0}
	shots    []*shot
)

func shotMgrProcess() {
	newShots := []*shot{}
	for _, s := range shots {
		shotActs[s.Type](s)
		s.count++
		if !s.finished {
			newShots = append(newShots, s)
		}

	}
	shots = newShots
}

func shotRegister(enemyID string, s shot) {
	s.owner = enemyID
	s.finished = false
	s.count = 0
	s.id = uuid.New().String()
	shots = append(shots, &s)
}
