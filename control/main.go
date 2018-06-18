package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var line strings.Builder

	for i := 1; i <= 100; i++ {
		hasFactor := false

		if i%3 == 0 {
			line.WriteString("Fizz")
			hasFactor = true
		}

		if i%5 == 0 {
			line.WriteString("Buzz")
			hasFactor = true
		}

		if !hasFactor {
			line.WriteString(strconv.Itoa(i))
		}

		line.WriteString("\n")
	}

	fmt.Println(line.String())
}
