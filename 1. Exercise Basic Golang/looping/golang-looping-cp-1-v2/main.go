package main

import "fmt"

func CountingNumber(n int) float64 {
	var total float64
	for iterate := 1.0; iterate <= float64(n); iterate += 0.5 {
		total += iterate
	}
	return total
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(CountingNumber(10))
}
