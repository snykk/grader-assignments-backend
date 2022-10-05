package main

import (
	"fmt"
	"sort"
)

func Sortheight(height []int) []int {
	sort.Ints(height)

	return height
}

func main() {
	var heights = []int{155, 156, 160, 161, 162, 165, 170, 172}

	fmt.Println(Sortheight(heights))
}
