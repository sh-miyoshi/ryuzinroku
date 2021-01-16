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
	baseAngle  float64
	charID     string
	bulletInfo bullet.Bullet
}

var (
	// ショットパターン
	// 0 ~ 9: 雑魚適用のショット
	// 10 ~ : ボス用の弾幕
	shotActs = []func(float64, float64, *Shot){
		shotAct0, shotAct1, shotAct2, shotAct3, shotAct4,
		shotAct5, shotAct6, shotAct7, shotAct8, shotActNon,
		bossShotAct0, bossShotAct1, bossShotAct2,
	}
)

// New ...
func New(typ int, charID string, b bullet.Bullet) *Shot {
	return &Shot{
		typ:        typ,
		id:         uuid.New().String(),
		finished:   false,
		count:      0,
		bulletInfo: b,
		charID:     charID,
	}
}

// Process ...
func (s *Shot) Process(ex, ey float64) bool {
	shotActs[s.typ](ex, ey, s)

	s.count++
	return s.finished
}
