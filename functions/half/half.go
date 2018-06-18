package half

// Half returns the input divided by two and whether the input is even or not
func Half(x int) (half int, isEven bool) {
	return x / 2, x%2 == 0
}
