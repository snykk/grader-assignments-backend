package main

import (
	"fmt"
	"strings"
	"unicode"
)

func ReverseWord(str string) string {
	if str == "" {
		return ""
	}

	arrWords := strings.Split(str, " ")
	var dummy, result string
	var needCapt bool

	for i := 0; i < len(arrWords); i++ {
		dummy = ""
		needCapt = false
		if unicode.IsUpper(rune(arrWords[i][0])) {
			needCapt = true
		}

		for index, char := range arrWords[i] {

			if index == len(arrWords[i])-1 && needCapt {
				dummy = strings.ToUpper(fmt.Sprintf("%c", char)) + dummy
				continue
			}
			dummy = strings.ToLower(fmt.Sprintf("%c", char)) + dummy
		}

		if i == len(arrWords)-1 {
			result += dummy
			continue
		}

		result += dummy + " "
	}

	return result

}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(ReverseWord("Aku Sayang Ibu"))
	fmt.Println(ReverseWord("A bird fly to the Sky"))
}
