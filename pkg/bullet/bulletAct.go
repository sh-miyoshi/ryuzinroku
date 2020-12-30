package bullet

// 特に何もしない
func bulletAct0(b *Bullet) {}

// 減速
func bulletAct1(b *Bullet) {
	if b.Speed > 1.5 {
		b.Speed -= 0.04
	}
}
