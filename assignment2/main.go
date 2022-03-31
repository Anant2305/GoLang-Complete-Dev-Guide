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

func main() {
	tr := triangle{
		base:   3,
		height: 7,
	}
	sq := square{
		sideLength: 5,
	}

	printArea(tr)
	printArea(sq)
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func (tr triangle) getArea() float64 {
	return 0.5 * tr.base * tr.height
}

func (sq square) getArea() float64 {
	return sq.sideLength * sq.sideLength
}
