package main

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base   float64
	Height float64
}

// Rectangle section

func (r Rectangle) Area() float64 { // area
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 { // perimeter
	return 2 * (r.Width + r.Height)
}

func (r *Rectangle) Scale(factor float64) { // Scale
	r.Width *= factor
	r.Height *= factor
}

// Circle section

func (c Circle) Area() float64 { // Area
	return math.Pi * (c.Radius * c.Radius)
}

func (c Circle) Perimeter() float64 { // circumference is the perimeter of a circle
	return 2 * math.Pi * c.Radius
}

func (c *Circle) Scale(factor float64) { //scale
	c.Radius *= factor
}

// triangle section

func (t Triangle) Area() float64 { // area
	return 0.5 * t.Base * t.Height
}

func (t Triangle) Perimeter() float64 { // perimeter
	return 3 * t.Base
}

func (t *Triangle) Scale(factor float64) { // scale
	t.Base *= factor
	t.Height *= factor
}
