package main

import "fmt"

func isArrayContains(s []int, num int) bool {
	for _, v := range s {
		if v == num {
			return true
		}
	}

	return false
}

func isNumContainsInEveryArray(villager [][]int, num int) bool {

	for _, arr := range villager {
		if !isArrayContains(arr, num) {
			return false
		}
	}

	return true
}

func SchedulableDays(villager [][]int) []int {
	var result = []int{}
	for _, i := range villager {
		for _, j := range i {
			if isNumContainsInEveryArray(villager, j) && !isArrayContains(result, j) {
				result = append(result, j)
			}
		}
	}

	return result
}

func main() {
	var data = [][]int{{7, 12, 19, 22}, {12, 19, 21, 23}, {7, 12, 19}, {12, 19}}

	fmt.Println(SchedulableDays(data))
}
