package main

import (
	"fmt"
	"strings"
)

func CountVowelConsonant(str string) (int, int, bool) {
	var consonant, vowels int
	var check bool
	vowelsHelper := map[string]bool{
		"a": true,
		"i": true,
		"u": true,
		"e": true,
		"o": true,
	}

	for index := 0; index < len(str); index++ {
		if vowelsHelper[strings.ToLower(fmt.Sprintf("%c", str[index]))] {
			vowels += 1
		} else if 65 <= str[index] && str[index] <= 90 || 97 <= str[index] && str[index] <= 122 {
			consonant += 1
		}

	}
	if vowels == 0 || consonant == 0 {
		check = true
	}

	return vowels, consonant, check
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(CountVowelConsonant("bbbbb ccccc"))
}
