package main

import (
	"fmt"
	"math"
	"reflect"
)

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
	newSquare := square{sideLength: 5}
	newTriangle := triangle{height: 5, base: 3}

	printArea(newSquare)
	printArea(newTriangle)
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return math.Pow(s.sideLength, 2)
}

func printArea(s shape) {
	fmt.Println("Area of", reflect.TypeOf(s).Name(), ":", s.getArea())
}
