package main

import "fmt"

func BMICalculator(gender string, height int) float64 {
	var percent float64
	if gender == "laki-laki" {
		percent = 0.1
	} else if gender == "perempuan" {
		percent = 0.15
	}

	return (float64(height) - 100) - ((float64(height) - 100) * percent)
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(BMICalculator("laki-laki", 165))
	fmt.Println(BMICalculator("perempuan", 165))
}
