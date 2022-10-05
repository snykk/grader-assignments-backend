package main

import (
	"fmt"
	"strings"
)

func FindSimilarData(input string, data ...string) string {
	var result string
	for index := 0; index < len(data); index++ {
		if strings.Contains(data[index], input) {
			result += data[index] + ","
		}
	}

	return result[:len(result)-1]
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(FindSimilarData("iphone", "laptop", "iphone 13", "iphone 12", "iphone 12 pro"))

	fmt.Println(FindSimilarData("mobil", "mobil APV, mobil Avanza, motor matic, motor gede"))
}
