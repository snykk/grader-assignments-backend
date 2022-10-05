package main

import (
	"fmt"
	"strconv"
)

func ReverseData(arr [5]int) [5]int {
	var result [5]int
	var numStr, dummy string
	for index := 0; index < len(arr); index++ {
		numStr = strconv.Itoa(arr[index])
		dummy = ""
		for _, char := range numStr {
			dummy = fmt.Sprintf("%c", char) + dummy
		}

		result[4-index], _ = strconv.Atoi(dummy)
	}

	return result
}

func main() {
	var myArr = [5]int{456789, 44332, 2221, 12, 10}
	fmt.Println(ReverseData(myArr))
}
