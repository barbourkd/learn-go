package shapes

import "math"

// Shape is an interface for doing calculations on shapes
type Shape interface {
	Area() float64
	Perimeter() float64
}

// The Rectangle is a motherfuckin' rectangle
type Rectangle struct {
	Width  float64
	Height float64
}

// Perimeter returns the perimeter of a rectangle
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Area returns the area of a rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// The Circle is a dang circle
type Circle struct {
	Radius float64
}

// Area returns the area of a circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Perimeter returns the perimeter of a circle
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// The Triangle is yeah
type Triangle struct {
	Base   float64
	Height float64
}

// Area returns the area of a triangle
func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}

// Perimeter returns the perimeter of a triangle
func (t Triangle) Perimeter() float64 {
	return 0
}
