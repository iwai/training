package main

import "fmt"

func fibonacci() func() int {
	a, b := 0, 1

	return func() int {
		defer func() int { a, b = b, a+b; return a }()
		return a
	}
}

func fibonacci2(n int) int {
	switch n {
	case 0:
		fallthrough
	case 1:
		return n
	default:
		return fibonacci2(n-2) + fibonacci2(n-1)
	}
}

func fibonacci3() func() int {
	a, b := 0, 1

	return func() int {
		r := a
		a, b = b, a+b
		return r
	}
}

func tribonacci() func() int {
	a, b, c := 0, 0, 1

	return func() int {
		r := a
		a, b, c = b, c, a+b+c
		return r
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Printf("%v ", f())
	}
	fmt.Print("\n")

	for i := 0; i < 10; i++ {
		fmt.Printf("%v ", fibonacci2(i))
	}
	fmt.Print("\n")

	f = fibonacci3()
	for i := 0; i < 10; i++ {
		fmt.Printf("%v ", f())
	}
	fmt.Print("\n")

	f = tribonacci()
	for i := 0; i < 10; i++ {
		fmt.Printf("%v ", f())
	}
}
