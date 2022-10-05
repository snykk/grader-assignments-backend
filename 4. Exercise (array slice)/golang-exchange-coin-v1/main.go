package main

import "fmt"

func ExchangeCoin(amount int) []int {
	var cointList = []int{}
	for amount > 0 {
		if amount >= 1000 {
			cointList = append(cointList, 1000)
			amount -= 1000
		} else if amount >= 500 {
			cointList = append(cointList, 500)
			amount -= 500
		} else if amount >= 200 {
			cointList = append(cointList, 200)
			amount -= 200
		} else if amount >= 100 {
			cointList = append(cointList, 100)
			amount -= 100
		} else if amount >= 50 {
			cointList = append(cointList, 50)
			amount -= 50
		} else if amount >= 20 {
			cointList = append(cointList, 20)
			amount -= 20
		} else if amount >= 10 {
			cointList = append(cointList, 10)
			amount -= 10
		} else if amount >= 5 {
			cointList = append(cointList, 5)
			amount -= 5
		} else if amount >= 1 {
			cointList = append(cointList, 1)
			amount -= 1
		}
	}
	return cointList
}

func main() {
	fmt.Println(ExchangeCoin(1752))
}
