package main

import (
	"fmt"
	"math"
)

func type_conversions() {
	// i := 42
	// f := float64(i)
	// u := uint(f)

	var x, y int = 3, 4
	//var f float64 = math.Sqrt(float64(x*x + y*y))
	f := math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)

	fmt.Println(z)
	// fmt.Println(i, f, u, z)

}
