package minion

import (
	"math"

	"github.com/sh-miyoshi/ryuzinroku/pkg/common"
)

// 移動パターン0
// 下がってきて停滞して上がっていく
func act0(obj *Minion) {
	if obj.count == 0 {
		obj.vy = 3
	}

	if obj.count == 40 {
		obj.vy = 0
	}

	if obj.count == 40+obj.Wait {
		obj.vy = -3
	}
}

// 移動パターン1
// 下がってきて停滞して左下に行く
func act1(obj *Minion) {
	if obj.count == 0 {
		obj.vy = 3
	}

	if obj.count == 40 {
		obj.vy = 0
	}

	if obj.count == 40+obj.Wait {
		obj.vx = -1
		obj.vy = 2
		obj.direct = common.DirectLeft
	}
}

// 移動パターン2
// 下がってきて停滞して右下に行く
func act2(obj *Minion) {
	if obj.count == 0 {
		obj.vy = 3
	}

	if obj.count == 40 {
		obj.vy = 0
	}

	if obj.count == 40+obj.Wait {
		obj.vx = 1
		obj.vy = 2
		obj.direct = common.DirectRight
	}
}

// 行動パターン3
// すばやく降りてきて左へ
func act3(obj *Minion) {
	if obj.count == 0 {
		obj.vy = 5
	}

	if obj.count == 30 {
		obj.direct = common.DirectLeft
	}

	if obj.count < 100 {
		obj.vx -= 5.0 / 100.0
		obj.vy -= 5.0 / 100.0
	}
}

// 行動パターン4
// すばやく降りてきて右へ
func act4(obj *Minion) {
	if obj.count == 0 {
		obj.vy = 5
	}

	if obj.count == 30 {
		obj.direct = common.DirectRight
	}

	if obj.count < 100 {
		obj.vx += 5.0 / 100.0
		obj.vy -= 5.0 / 100.0
	}
}

// 行動パターン5
// 斜め左下へ
func act5(obj *Minion) {
	if obj.count == 0 {
		obj.vx = -1
		obj.vy = 2
		obj.direct = common.DirectLeft
	}
}

// 行動パターン6
// 斜め右下へ
func act6(obj *Minion) {
	if obj.count == 0 {
		obj.vx = 1
		obj.vy = 2
		obj.direct = common.DirectRight
	}
}

// 移動パターン7
// 停滞してそのまま左上に
func act7(obj *Minion) {
	if obj.count == obj.Wait {
		obj.vx = -0.7
		obj.vy = -0.3
		obj.direct = common.DirectLeft
	}
}

// 移動パターン8
// 停滞してそのまま右上に
func act8(obj *Minion) {
	if obj.count == obj.Wait {
		obj.vx = 0.7
		obj.vy = -0.3
		obj.direct = common.DirectRight
	}
}

// 移動パターン9
// 停滞してそのまま上に
func act9(obj *Minion) {
	if obj.count == obj.Wait {
		obj.vy = -1
	}
}

// 移動パターン10
// 下がってきてウロウロして上がっていく
func act10(obj *Minion) {
	if obj.count == 0 {
		obj.vy = 4 // 下がってくる
	}
	if obj.count == 40 {
		obj.vy = 0 // 止まる
	}
	if obj.count >= 40 {
		if obj.count%60 == 0 {
			r := 1
			obj.direct = common.DirectLeft
			if math.Cos(obj.angle) < 0 {
				r = 0
				obj.direct = common.DirectRight
			}
			obj.Speed = 6 + common.RandomAngle(2)
			obj.angle = common.RandomAngle(math.Pi/4) + math.Pi*float64(r)
		}
		obj.Speed *= 0.95
	}
	if obj.count >= 40+obj.Wait {
		obj.vy -= 0.05
	}
}
