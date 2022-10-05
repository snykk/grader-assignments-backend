package main

import (
	"fmt"
)

func PhoneNumberChecker(number string, result *string) {
	var lengthMin, pad int = 10, 0
	if fmt.Sprintf("%c", number[0]) == "6" {
		lengthMin += 1
		pad += 1
	}

	if len(number) < lengthMin {
		*result = "invalid"
		return
	}

	if fmt.Sprintf("%c", number[pad+2]) == "1" {
		if 49 <= number[pad+3] && number[pad+3] <= 53 { // based on dec value | 1 => 49, 2 => 50, ...
			*result = "Telkomsel"
		} else if 54 <= number[pad+3] && number[pad+3] <= 57 { // based on dec value | 1 => 49, 2 => 50, ...
			*result = "Indosat"
		}
	} else if fmt.Sprintf("%c", number[pad+2]) == "2" {

		if 49 <= number[pad+3] && number[pad+3] <= 51 { // based on dec value | 1 => 49, 2 => 50, ...
			*result = "XL"
		} else if 55 <= number[pad+3] && number[pad+3] <= 57 {
			*result = "Tri"
		}
	} else if fmt.Sprintf("%c", number[pad+2]) == "5" {
		if 50 <= number[pad+3] && number[pad+3] <= 51 {
			*result = "AS"
		}
	} else if fmt.Sprintf("%c", number[pad+2]) == "8" {
		if 49 <= number[pad+3] && number[pad+3] <= 56 {
			*result = "Smartfren"
		}
	} else {
		*result = "invalid"
	}
}

func main() {
	// bisa digunakan untuk pengujian test case
	var number = "08193456123"
	var result string

	PhoneNumberChecker(number, &result)
	fmt.Println(result)
}
