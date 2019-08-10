package main

import "fmt"

func main() {
	sum := 1
	for sum < 1000 {
		fmt.Println(sum)
		sum += sum
	}

	fmt.Println(sum)

	for i, t := range []int{5, 6, 7} {
		fmt.Printf("i = %d, t = %d\n", i, t)
	}

	loop := 0
	for {
		loop++
		if loop > 10 {
			break
		}
		if loop < 5 {
			continue
		}
		fmt.Printf("Loop: %v\n", loop)
	}

}
