package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(AbsFunc(v))

	p := &Vertex{3, 4}
	fmt.Println(p.Abs())
	fmt.Println(AbsFunc(*p))

	// v.Scale(2)
	// ScaleFunc(&v, 10)

	// p := &Vertex{3, 4}
	// p.Scale(2)
	// ScaleFunc(p, 10)

	// fmt.Println(v, p)
	// // fmt.Println(Abs(v))
}
