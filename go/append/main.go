package main

import "fmt"

func main() {
	var s []int
	var a [5]int8
	var b [5]int16
	printSlice(s)

	fmt.Printf("%v\n", &a[0])
	fmt.Printf("%v\n", &a[1])
	fmt.Printf("%v\n", &a[2])
	fmt.Printf("%v\n", &a[3])

	fmt.Printf("%v\n", &b[0])
	fmt.Printf("%v\n", &b[1])
	fmt.Printf("%v\n", &b[2])
	fmt.Printf("%v\n", &b[3])

	s = append(s, 0)
	printSlice(s)
	fmt.Printf("%v\n", &s[0])

	s = append(s, 1)
	printSlice(s)
	fmt.Printf("%v\n", &s[0])

	s = append(s, 2, 3, 4)
	printSlice(s)

	fmt.Printf("%v\n", &s[0])
	fmt.Printf("%v\n", &s[1])
	fmt.Printf("%v\n", &s[2])
	fmt.Printf("%v\n", &s[3])

}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
