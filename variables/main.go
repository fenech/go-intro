package main

import (
	"fmt"
)

func main() {
	fmt.Print("Enter a temperature in Fahrenheit: ")
	var fahrenheit float64
	fmt.Scanf("%f", &fahrenheit)

	celsius := (fahrenheit - 32) * 5 / 9

	fmt.Println(fahrenheit, "fahrenheit is", celsius, "celsius")
}
