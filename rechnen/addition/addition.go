package addition

func Add(x, y int) int {

	if y == 0 {
		return x
	}

	// x + y == x + (y-1) + 1
	return Add(x, y-1) + 1
}
