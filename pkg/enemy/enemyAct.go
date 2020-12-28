package enemy

func act0(obj *enemy) {
	if obj.count < 60 {
		obj.Y += 2
	}
	if obj.count > 60+240 {
		obj.Y -= 2
	}
}
