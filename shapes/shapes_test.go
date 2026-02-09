package main

import (
	"math"
	"testing"
)

const epsilon = 0.0001 // using epsilon for tolerange. We should never compare floats directly.

// floatsEqual handles "floating-point inaccuracy."
// Instead of checking if a and b are identical, it checks if the
// distance between them is smaller than a tiny margin (epsilon).
// This prevents bugs caused by tiny rounding errors in decimal math.
func floatsEqual(a, b float64) bool {
	return math.Abs(a-b) < epsilon
}

// testing AREA on 3 structs
func TestRectangleArea(t *testing.T) { // triangle
	r := Rectangle{Width: 4, Height: 5}
	expected := 20.0

	if !floatsEqual(r.Area(), expected) {
		t.Errorf("expected %v, got %v", expected, r.Area())
	}
}

func TestCircleArea(t *testing.T) { // circle
	c := Circle{Radius: 1}
	expected := 3.14159

	if !floatsEqual(c.Area(), expected) {
		t.Errorf("expected %v, got %v", expected, c.Area())
	}
}

func TestTriangleArea(t *testing.T) { // triangle
	tr := Triangle{Base: 2, Height: 3}
	expected := 3.0

	if !floatsEqual(tr.Area(), expected) {
		t.Errorf("expected %v, got %v", expected, tr.Area())
	}
}

// testing Perimeter for all three structs
func TestRectanglePerimeter(t *testing.T) {
	r := Rectangle{Width: 4, Height: 5}
	expected := 18.0

	if !floatsEqual(r.Perimeter(), expected) {
		t.Errorf("expected %v, got %v", expected, r.Perimeter())
	}
}

func TestCirclePerimeter(t *testing.T) {
	c := Circle{Radius: 1}
	expected := 6.28318

	if !floatsEqual(c.Perimeter(), expected) {
		t.Errorf("expected %v, got %v", expected, c.Perimeter())
	}
}

func TestTrianglePerimeter(t *testing.T) {
	tr := Triangle{Base: 4}
	expected := 12.0

	if !floatsEqual(tr.Perimeter(), expected) {
		t.Errorf("expected %v, got %v", expected, tr.Perimeter())
	}
}

// testing scale for all three structs
func TestRectangleScale(t *testing.T) {
	r := Rectangle{Width: 2, Height: 3}
	r.Scale(2)

	if r.Width != 4 || r.Height != 6 {
		t.Errorf("rectangle not scaled correctly")
	}
}

func TestCircleScale(t *testing.T) {
	c := Circle{Radius: 2}
	c.Scale(3)

	if c.Radius != 6 {
		t.Errorf("circle not scaled correctly")
	}
}

func TestTriangleScale(t *testing.T) {
	tr := Triangle{Base: 2, Height: 3}
	tr.Scale(2)

	if tr.Base != 4 || tr.Height != 6 {
		t.Errorf("triangle not scaled correctly")
	}
}
