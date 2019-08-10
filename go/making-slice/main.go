package main

import (
	"fmt"
	"strings"
)

func main() {
	a := make([]int, 5)
	printSclice("a", a)

	b := make([]int, 0, 5)
	printSclice("b", b)

	c := b[:2]
	printSclice("c", c)

	d := c[2:5]
	printSclice("d", d)

	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

func printSclice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
	if x != nil {
		fmt.Println("Not nil!")
	}
}
