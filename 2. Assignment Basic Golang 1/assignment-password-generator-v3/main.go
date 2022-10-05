package main

import (
	"fmt"
	"strings"
	"unicode"
)

// vokal
func shiftVowel(str string) string {
	vowelHelper := map[string]string{
		// "a": "E",
		// "e": "I",
		// "i": "O",
		// "o": "U",
		// "u": "A",
		// "A": "e",
		// "E": "i",
		// "I": "o",
		// "O": "u",
		// "U": "a",
		"a": "e",
		"e": "i",
		"i": "o",
		"o": "u",
		"u": "a",
		"A": "E",
		"E": "I",
		"I": "O",
		"O": "U",
		"U": "A",
	}
	var result string

	for _, char := range str {
		switch char {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			result += vowelHelper[fmt.Sprintf("%c", char)]
			continue
		}

		result += fmt.Sprintf("%c", char)
	}

	return result
}

func shiftCaptAndClearSpace(str string) string {
	var result string
	for index := 0; index < len(str); index++ {
		if fmt.Sprintf("%c", str[index]) != " " {
			if unicode.IsUpper(rune(str[index])) {
				result += strings.ToLower(fmt.Sprintf("%c", str[index]))
			} else {
				result += strings.ToUpper(fmt.Sprintf("%c", str[index]))
			}
		}
	}

	return result
}

// reverse string
func Reverse(str string) string {
	var result string
	for index := 0; index < len(str); index++ {
		result = fmt.Sprintf("%c", str[index]) + result
	}

	return result
}

func Generate(str string) string {
	str = Reverse(str)
	str = shiftVowel(str)
	str = shiftCaptAndClearSpace(str)

	return str
}

func onlyLetterAndNumbers(str string) bool {

	for _, char := range str {
		if 48 <= char && char <= 57 || 65 <= char && char <= 90 || 97 <= char && char <= 122 {
			continue
		}
		return false
	}

	return true
}

func advancePass(str string) bool {
	if strings.ContainsAny(str, "~`!@#$%^&*()_+-=[]{}:;'?/>.<,|") && len(str) < 14 {
		return true
	}

	return false
}

func CheckPassword(str string) string {
	if len(str) < 7 {
		return "sangat lemah"
	} else if onlyLetterAndNumbers(str) {
		return "lemah"
	} else if advancePass(str) {
		return "sedang"
	}
	return "kuat"
}

func PasswordGenerator(base string) (string, string) {
	return Generate(base), CheckPassword(Generate(base))
}

func main() {
	// data := "Semangat Pagi" // bisa digunakan untuk melakukan debug

	res, check := PasswordGenerator("Semangat Pagi")

	fmt.Println(res)
	fmt.Println(check)
}
