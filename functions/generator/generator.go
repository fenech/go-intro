package generator

// MakeOddGenerator returns a function that generates odd numbers
func MakeOddGenerator() func() uint {
	i := uint(1)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

// Fibonacci returns
func Fibonacci(count uint) uint {
	switch count {
	case 0:
		return 0
	case 1:
		return 1
	default:
		return Fibonacci(count-1) + Fibonacci(count-2)
	}
}
