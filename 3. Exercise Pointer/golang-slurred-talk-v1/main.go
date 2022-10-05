package main

import (
	"fmt"
	"unicode"
)

func SlurredTalk(words *string) {
	var result string
	for _, char := range *words {
		switch char {
		case 'S', 'R', 'Z', 's', 'r', 'z':
			if unicode.IsUpper(char) {
				result += "L"
			} else {
				result += "l"
			}
			continue
		}
		result += fmt.Sprintf("%c", char)
	}

	*words = result
}

func main() {
	// bisa dicoba untuk pengujian test case
	var words string = "Steven"
	SlurredTalk(&words)
	fmt.Println(words)
}
