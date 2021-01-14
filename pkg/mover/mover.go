package mover

import (
	"errors"
)

type charInfo struct {
	active           bool
	charID           string
	x, y             *float64
	targetX, targetY float64
	initX, initY     float64
	targetCount      int
	count            int
	v0x, v0y         float64
	ax, ay           float64
}

var (
	charList = []*charInfo{}

	// ErrNoSuchChar ...
	ErrNoSuchChar = errors.New("no such char")
)

// CharRegist register the character to management list
func CharRegist(charID string, x, y *float64) {
	c := charInfo{
		active: false,
		charID: charID,
		x:      x,
		y:      y,
	}
	charList = append(charList, &c)
}

// CharRemove ...
func CharRemove(charID string) {
	newList := []*charInfo{}
	for _, c := range charList {
		if c.charID != charID {
			newList = append(newList, c)
		}
	}
	charList = newList
}

// MoveTo ...
func MoveTo(charID string, targetX, targetY float64, targetCount int) error {
	for _, c := range charList {
		if c.charID == charID {
			c.active = true
			c.count = 0
			c.targetCount = targetCount
			t := float64(targetCount)

			c.targetX = targetX
			c.initX = *c.x
			dist := *c.x - targetX
			c.v0x = 2 * dist / t
			c.ax = 2 * dist / (t * t)

			c.targetY = targetY
			c.initY = *c.y
			dist = *c.y - targetY
			c.v0y = 2 * dist / t
			c.ay = 2 * dist / (t * t)
			return nil
		}
	}

	return ErrNoSuchChar
}

// Process ...
func Process() {
	for _, c := range charList {
		if c.active {
			c.count++

			// set current x and y
			t := float64(c.count)
			*c.x = c.initX - ((c.v0x * t) - 0.5*c.ax*t*t)
			*c.y = c.initY - ((c.v0y * t) - 0.5*c.ay*t*t)

			if c.count >= c.targetCount {
				c.active = false
			}
		}
	}
}
