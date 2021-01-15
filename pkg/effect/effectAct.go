package effect

import "github.com/sh-miyoshi/ryuzinroku/pkg/sound"

func effectActDead(e *effect) bool {
	e.ExtRate += 0.08
	if e.count > 10 {
		e.BlendParam -= 25
	}
	return e.count > 20
}

func effectActBombChar(e *effect) bool {
	if e.count < 51 {
		e.BlendParam += 4
	}
	if e.count > 130-51 {
		e.BlendParam -= 4
	}
	return e.count >= 130
}

func effectActBombTitle(e *effect) bool {
	if e.count < 51 {
		e.BlendParam += 2
	}
	if e.count > 130-51 {
		e.BlendParam -= 2
	}
	return e.count >= 130
}

func effectActBombMain(e *effect) bool {
	// スピード計算
	if e.count < 60 {
		e.Speed -= (0.2 + float64(e.count*e.count)/3000.0)
	}
	if e.count == 60 {
		sound.PlaySound(sound.SEBomb)
		e.Speed = 0
		// TODO
		// dn.flag = 1
		// dn.cnt = 0
		// dn.size = 11
		// dn.time = 20
	}
	// 明るさと大きさ計算
	e.ExtRate += 0.015
	if e.count < 51 {
		e.BlendParam += 5
	}
	if e.count >= 60 {
		e.ExtRate += 0.04
		e.BlendParam -= 255 / 30.0
	}
	return e.count >= 90
}
