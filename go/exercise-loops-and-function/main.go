package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := x
	n := 0.0

	var abs float64

	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Printf("%v\n", z)

		abs = math.Abs(z - n)
		if abs < 1e-10 {
			return z
		}
		n = z
	}
	return z
}

func main() {
	fmt.Println("orig: ", Sqrt(12))
	fmt.Println("math: ", math.Sqrt(12))
}
