package bullet

var (
	stack []int
)

// PushBomb ...
func PushBomb(damage int) {
	stack = append(stack, damage)
}

// PopBomb get damage by bomb
// return 0 when a bomb damage is not pushed
func PopBomb() int {
	if len(stack) > 0 {
		res := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		return res
	}

	return 0
}
