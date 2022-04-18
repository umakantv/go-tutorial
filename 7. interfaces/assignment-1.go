package main

import (
	"fmt"
	"math"
)

type shape interface {
	getArea() float64
}

type square struct {
	sideLength float64
}

func (s square) getArea() float64 {
	return math.Pow(s.sideLength, 2)
}

type triangle struct {
	base   float64
	height float64
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func main() {
	sq := square{2.}
	tr := triangle{2., 5.}

	printArea(sq)
	printArea(tr)

}
