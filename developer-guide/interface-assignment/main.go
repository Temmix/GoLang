package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	height float64
	base   float64
}

type square struct {
	sideLength float64
}

func main() {
	// square
	s := square{sideLength: 20}
	printArea(s)

	// triangle
	t := triangle{height: 10, base: 15}
	printArea(t)
}

func printArea(s shape) {
	fmt.Println("Area of the shape is: ", s.getArea())
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}
