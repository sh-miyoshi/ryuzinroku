package effect

func effectAct0(e *effect) bool {
	e.ExtRate += 0.08
	if e.count > 10 {
		e.BlendParam -= 25
	}
	return e.count > 20
}
