package main

import (
	"fmt"
	"strconv"
	"strings"
)

func ChangeOutput(data []string) map[string][]string {
	var result = map[string][]string{}
	var strSplit = []string{}

	var header, position, value string
	var index int
	for _, item := range data {
		strSplit = strings.Split(item, "-")
		header = strSplit[0]
		index, _ = strconv.Atoi(strSplit[1])
		position = strSplit[2]
		value = strSplit[3]

		if position == "first" {
			result[header] = append(result[header], value)
			continue
		}

		result[header][index] = result[header][index] + " " + value

	}
	return result
}

func main() {
	data := []string{"account-0-first-John", "account-0-last-Doe", "account-1-first-Jane", "account-1-last-Doe", "address-0-first-Jaksel", "address-0-last-Jakarta", "address-1-first-Bandung", "address-1-last-Jabar"}
	fmt.Println(ChangeOutput(data))
}
