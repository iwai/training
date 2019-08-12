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

func (v Vertex) Scale(f float64) {
	fmt.Printf("%v\n", &v.X)
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	fmt.Printf("%v\n", &v.X)
	fmt.Printf("%v\n", v)
	v.Scale(10)
	fmt.Printf("%v\n", v)
	fmt.Println(v.Abs())
}
