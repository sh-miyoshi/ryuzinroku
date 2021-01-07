package boss

type mover struct {
	targetX, targetY float64
	initX, initY     float64
	targetCount      int
	count            int
	v0x, v0y         float64
	ax, ay           float64
}

func (m *mover) moveTo(currentX, currentY, targetX, targetY float64, count int) {
	m.count = 0
	m.targetCount = count
	t := float64(count)

	m.targetX = targetX
	m.initX = currentX
	dist := currentX - targetX
	m.v0x = 2 * dist / t
	m.ax = 2 * dist / (t * t)

	m.targetY = targetY
	m.initY = currentY
	dist = currentY - targetY
	m.v0y = 2 * dist / t
	m.ay = 2 * dist / (t * t)
}

func (m *mover) process() {
	m.count++
}

func (m *mover) currentPos() (x, y float64) {
	if m.count >= m.targetCount {
		return m.targetX, m.targetY
	}

	t := float64(m.count)
	x = m.initX - ((m.v0x * t) - 0.5*m.ax*t*t)
	y = m.initY - ((m.v0y * t) - 0.5*m.ay*t*t)

	return x, y
}
