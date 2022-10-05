package main

import (
	"fmt"
	"strconv"
)

func BiggestPairNumber(numbers int) int {
	numbersStr := strconv.Itoa(numbers)
	var dummy, num1, num2, max, fixnum1, fixnum2 int
	for index := 0; index < len(numbersStr)-1; index++ {
		num1, _ = strconv.Atoi(fmt.Sprintf("%c", numbersStr[index]))
		num2, _ = strconv.Atoi(fmt.Sprintf("%c", numbersStr[index+1]))

		dummy = num1 + num2
		if dummy > max {
			max = dummy
			fixnum1 = num1
			fixnum2 = num2
		}
	}

	return fixnum1*10 + fixnum2
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(BiggestPairNumber(11223344))
}
