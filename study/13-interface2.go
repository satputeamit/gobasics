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
	t := triangle{base: 2, height: 3}
	sq := square{sideLength: 2}
	printArea(t)
	printArea(sq)

}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func (s square) getArea() float64 {

	return s.sideLength * s.sideLength
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}
