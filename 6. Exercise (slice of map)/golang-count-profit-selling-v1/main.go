package main

import "fmt"

func CountProfit(data [][][2]int) []int {
	if len(data) == 0 {
		return []int{}
	}
	var profitMonth = make([]int, len(data[0]))
	for _, cabang := range data {
		for i, bulan := range cabang {
			profitMonth[i] += bulan[0] - bulan[1]
			fmt.Println(profitMonth)
		}
	}

	return profitMonth[:]
}

func main() {
	var data = [][][2]int{
		{ // cabang 1
			{1000, 500}, {500, 150}, {600, 100}, {800, 750}, //representasi profit dan pengeluaran tiap bulang
		}, {
			// cabang 2
			{1000, 800}, {200, 150}, {900, 100}, {800, 700},
		},
		// ...
	}
	// fmt.Println(len(data[0]))
	fmt.Println(CountProfit(data))
}
