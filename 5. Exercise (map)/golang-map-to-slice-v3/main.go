package main

import "fmt"

func MapToSlice(mapData map[string]string) [][]string {
	var result = [][]string{}

	for key, value := range mapData {
		result = append(result, []string{key, value})
	}

	return result
}

func main() {
	data := map[string]string{
		"hello": "world",
		"John":  "Doe",
		"age":   "14",
	}

	fmt.Println(MapToSlice(data))
}
