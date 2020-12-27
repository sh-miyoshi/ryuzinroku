package common

// MoveOK ...
func MoveOK(x, y int) bool {
	if x >= 10 && x <= ScreenX-10 && y >= 5 && y <= ScreenY-5 {
		return true
	}
	return false
}
