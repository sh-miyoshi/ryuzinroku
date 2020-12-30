package shot

import (
	"github.com/google/uuid"
	"github.com/sh-miyoshi/ryuzinroku/pkg/bullet"
)

// Shot ...
type Shot struct {
	typ        int
	id         string
	finished   bool
	count      int
	bulletInfo bullet.Bullet
}

var (
	shotActs = []func(float64, float64, *Shot){shotAct0}
)

// New ...
func New(typ int, b bullet.Bullet) *Shot {
	return &Shot{
		typ:        typ,
		id:         uuid.New().String(),
		finished:   false,
		count:      0,
		bulletInfo: b,
	}
}

// Process ...
func (s *Shot) Process(ex, ey float64) bool {
	shotActs[s.typ](ex, ey, s)

	s.count++
	return s.finished
}
