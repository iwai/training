package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	//countMap := make(map[string]int)
	countMap := map[string]int{}

	for _, v := range strings.Fields(s) {
		countMap[v]++
	}

	return countMap
}

func main() {
	wc.Test(WordCount)
}
