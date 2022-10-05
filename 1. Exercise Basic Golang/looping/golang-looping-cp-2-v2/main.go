package main

import (
	"fmt"
)

// hello World => d_l_r_o_W o_l_l_e_H
func ReverseString(str string) string {
	var result string

	for index := len(str) - 1; index >= 0; index-- {
		if index == 0 {
			result = result + fmt.Sprintf("%c", str[index])
			continue
		}
		if str[index] == 32 || str[index-1] == 32 {
			result = result + fmt.Sprintf("%c", str[index])
			continue
		}
		result = result + fmt.Sprintf("%c_", str[index])
	}
	return result
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(ReverseString("Hello World"))
}
