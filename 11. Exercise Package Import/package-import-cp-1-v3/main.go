package main

import (
	"a21hc3NpZ25tZW50/internal"
	"fmt"
	"strconv"
	"strings"
)

func AdvanceCalculator(calculate string) float32 {
	dataList := strings.Split(calculate, " ")

	base, _ := strconv.Atoi(dataList[0])

	myCalc := internal.NewCalculator(float32(base))

	for index, item := range dataList[1:] {
		if index%2 == 0 {
			num, _ := strconv.Atoi(dataList[1:][index+1])
			switch item {
			case "+":
				myCalc.Add(float32(num))
			case "-":
				myCalc.Subtract(float32(num))
			case "*":
				myCalc.Multiply(float32(num))
			case "/":
				myCalc.Divide(float32(num))
			}
		}
	}

	return myCalc.Result()
}

func main() {
	res := AdvanceCalculator("3 * 4 / 2 + 10 - 5")

	fmt.Println(res)
}
