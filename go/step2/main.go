package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
)

func add(x int, y int) int {
	return x + y
}

func subtract(x, y int) int {
	return x - y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

var i, j int = 1, 2

var (
	ToBe    bool       = false
	MaxuInt uint64     = 1<<64 - 1
	MaxInt  int        = 1<<63 - 1
	z       complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxuInt, MaxuInt)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	fmt.Printf("%f\n", 1.0/3)
	fmt.Printf("%f\n", float64(10/3))

	// var i int
	var c, python, java = true, false, "no!"
	k := 3
	fmt.Println(i, j, k, c, python, java)

	var n, err = fmt.Println("My favorite number is", rand.Intn(10))

	fmt.Println(n, err)

	n, err = fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))

	fmt.Println(n, err)

	fmt.Println(math.Pi)

	fmt.Println(add(42, 13))

	fmt.Println(subtract(42, 13))

	a, b := swap("hello", "world")
	fmt.Println(a, b)

	fmt.Println(split(18))
}
