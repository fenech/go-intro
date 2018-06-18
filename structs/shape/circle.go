package shape

import (
	"math"
)

type Circle struct {
	radius float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c *Circle) perimiter() float64 {
	return 2.0 * math.Pi * c.radius
}
