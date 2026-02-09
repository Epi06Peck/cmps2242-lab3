package main

import "fmt"

func main() {
	rect := Rectangle{Width: 4, Height: 5}
	circ := Circle{Radius: 3}
	trian := Triangle{Base: 6, Height: 4}

	// Function calls on rectangle instance
	fmt.Printf("Rectangle's Area: %.2f\n", rect.Area())
	fmt.Printf("Rectangle's Perimeter %.2f\n", rect.Perimeter())
	rect.Scale(2)
	fmt.Printf("After scaling rectangle * 2, we should get double the area: %.2f\n", rect.Area())

	// function calls on circle instance
	fmt.Printf("Circle's Area: %.2f\n", circ.Area())
	fmt.Printf("Circle's Circumference: %.2f\n", circ.Perimeter())
	circ.Scale(2)
	fmt.Printf("After scaling Circle * 2, we should get double the area: %.2f\n", circ.Area())

	// function
	fmt.Printf("Triangle's Area: %.2f\n", trian.Area())
	fmt.Printf("Triangle's Perimeter %.2f\n", trian.Perimeter())
	trian.Scale(2)
	fmt.Printf("After scaling Triangle * 2, we should get double the area: %.2f\n", trian.Area())
}
