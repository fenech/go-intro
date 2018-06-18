package max

// Max returns the argument with the highest value
func Max(args ...int) int {
	var max int
	for i, x := range args {
		if i == 0 || x > max {
			max = x
		}
	}

	return max
}
