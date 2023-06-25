package main

import "fmt"

type triangle struct {
	base   float64
	height float64
}
type square struct {
	sideLength float64
}

type shape interface {
	getArea() float64
}

func main() {
	t := triangle{
		base:   10,
		height: 5,
	}
	s := square{
		sideLength: 4,
	}
	printArea(t)
	printArea(s)
}

func printArea(s shape) {
	fmt.Println("The area of the element is ", s.getArea())
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}
