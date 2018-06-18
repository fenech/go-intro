package shape

type Rectangle struct {
	x1, y1, x2, y2 float64
}

func (r *Rectangle) length() float64 {
	return distance(r.x1, r.y1, r.x1, r.y2)
}

func (r *Rectangle) width() float64 {
	return distance(r.x1, r.y1, r.x2, r.y1)
}

func (r *Rectangle) area() float64 {
	return r.length() * r.width()
}

func (r *Rectangle) perimiter() float64 {
	return 2 * (r.length() + r.width())
}
