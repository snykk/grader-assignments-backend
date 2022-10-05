package main

import "fmt"

func isArrayContains(s []int, str int) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func SchedulableDays(date1 []int, date2 []int) []int {
	var mySlice = []int{}
	for index := 0; index < len(date1); index++ {
		if isArrayContains(date2, date1[index]) {
			mySlice = append(mySlice, date1[index])
		}
	}

	return mySlice
}

func main() {
	var date1 = []int{11, 12, 13, 14, 15}
	var date2 = []int{5, 10, 12, 13, 20, 21}
	fmt.Println(SchedulableDays(date1, date2))
}
