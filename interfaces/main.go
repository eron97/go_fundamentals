package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

type Point struct {
	X, Y float64
}

func (p Point) Abs() float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y)
}

var p = Point{X: 3, Y: 4}

func main() {
	a := Abser(p)
	fmt.Println(a.Abs())

}
