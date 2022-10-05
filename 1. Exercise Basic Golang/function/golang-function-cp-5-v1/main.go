package main

import (
	"fmt"
)

func FindMin(nums ...int) int {
	var minNums int = 999999

	for index := 0; index < len(nums); index++ {
		if nums[index] < minNums {
			minNums = nums[index]
		}
	}

	return minNums
}

func FindMax(nums ...int) int {
	var maxNums int = 0

	for index := 0; index < len(nums); index++ {
		if maxNums < nums[index] {
			maxNums = nums[index]
		}
	}

	return maxNums
}

func SumMinMax(nums ...int) int {
	return FindMin(nums...) + FindMax(nums...)
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(SumMinMax(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
}
