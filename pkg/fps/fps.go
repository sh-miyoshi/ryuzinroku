package fps

import (
	"time"

	"github.com/sh-miyoshi/dxlib"
)

const (
	targetFPS = 60
)

var (
	baseTime   int32
	count      int32
	currentFPS float32
)

// Wait ...
func Wait() {
	wait := int32(0)
	if count == 0 {
		baseTime = dxlib.GetNowCount()
	} else {
		c := dxlib.GetNowCount()

		if count == targetFPS-1 {
			// Update current FPS
			currentFPS = float32(targetFPS * 1000 / (c - baseTime))
		}

		target := count*1000/targetFPS + baseTime
		wait = target - c
	}

	if wait > 0 {
		time.Sleep(time.Millisecond * time.Duration(wait))
	}
	count = (count + 1) % targetFPS
}

// Draw ...
func Draw(x int32, y int32) {
	dxlib.DrawFormatString(x, y, dxlib.GetColor(255, 0, 0), "[%.1f]", currentFPS)
}
