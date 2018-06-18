package main

import (
	"fmt"

	"github.com/fenech/go-intro/functions/generator"
	"github.com/fenech/go-intro/functions/half"
	"github.com/fenech/go-intro/functions/max"
)

func main() {
	input := 4
	half, isEven := half.Half(input)

	fmt.Println("half of", input, "is", half)
	if isEven {
		fmt.Println(input, "is even")
	}

	slice := []int{1, 2, 3, 4}
	fmt.Println(max.Max(slice...))

	oddGen := generator.MakeOddGenerator()
	fmt.Println(oddGen())
	fmt.Println(oddGen())
	fmt.Println(oddGen())

	fmt.Println("fibonacci")
	for i := 0; i < 10; i++ {
		fmt.Println(generator.Fibonacci(uint(i)))
	}
}
