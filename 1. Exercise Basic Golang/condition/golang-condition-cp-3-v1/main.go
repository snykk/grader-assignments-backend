package main

import "fmt"

func GetPredicate(math, science, english, indonesia int) string {
	rate := (math + science + english + indonesia) / 4

	if rate == 100 {
		return "Sempurna"
	} else if rate >= 90 {
		return "Sangat Baik"
	} else if rate >= 80 {
		return "Baik"
	} else if rate >= 70 {
		return "Cukup"
	} else if rate >= 60 {
		return "Kurang"
	}
	return "Sangat kurang"
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(GetPredicate(50, 80, 100, 60))
}
