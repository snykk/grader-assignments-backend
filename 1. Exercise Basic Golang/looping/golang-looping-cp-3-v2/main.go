package main

import "fmt"

func CountingLetter(text string) int {
	helper := map[string]bool{
		"r": true,
		"R": true,
		"S": true,
		"s": true,
		"T": true,
		"t": true,
		"Z": true,
		"z": true,
	}
	var counter int

	for _, char := range text {
		if helper[fmt.Sprintf("%c", char)] {
			counter++
		}
	}
	// unreadable letters = R, S, T, Z
	return counter
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(CountingLetter("Semangat"))
}
