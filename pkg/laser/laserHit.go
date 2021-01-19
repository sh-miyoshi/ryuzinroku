package laser

func rawHitCheck(xc, yc float64, r float64, x1, y1, x2, y2 float64) bool {
	if xc > x1 && xc < x2 && yc > (y1-r) && yc < (y2+r) {
		return true
	}

	if xc > (x1-r) && xc < (x2+r) && yc > y1 && yc < y2 {
		return true
	}

	r2 := r * r
	if (x1-xc)*(x1-xc)+(y1-yc)*(y1-yc) < r2 {
		return true
	}
	if (x2-xc)*(x2-xc)+(y1-yc)*(y1-yc) < r2 {
		return true
	}
	if (x2-xc)*(x2-xc)+(y2-yc)*(y2-yc) < r2 {
		return true
	}
	if (x1-xc)*(x1-xc)+(y2-yc)*(y2-yc) < r2 {
		return true
	}

	return false
}
