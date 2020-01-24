package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	base   float64
	height float64
}

type square struct {
	sideLength float64
}

// By defining these methods, we implicitly make triangle and square implement shape.
// There's no explicit marking of the interface.
func (t triangle) getArea() float64 {
	return t.base * t.height / 2
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func main() {
	sq := square{sideLength: 5}
	tr := triangle{base: 4, height: 6}

	printArea(sq)
	printArea(tr)
}

// Method that takes the interface shape rather than one of the concrete types.
func printArea(s shape) {
	fmt.Println(s.getArea())
}
