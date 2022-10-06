package main

import "fmt"

func Add(a, b int) int {
	return 0
}

type School struct {
	Name    string
	Address string
	Grades  []int
}

func (s *School) AddGrade(grades ...int) {
	for _, item := range grades {
		s.Grades = append(s.Grades, item)
	}
}

// ===>
func getRate(grades []int) float64 {
	if len(grades) == 0 {
		return 0
	}

	var rate float64
	for _, item := range grades {
		rate += float64(item)
	}

	return rate / float64(len(grades))
}

// ===>
func getMin(grades []int) int {
	var result int = 100
	for _, item := range grades {
		if item < result {
			result = item
		}
	}

	return result
}

// ===>
func getMax(grades []int) int {
	var result int = 0
	for _, item := range grades {
		if result < item {
			result = item
		}
	}

	return result
}

func Analysis(s School) (float64, int, int) {
	if len(s.Grades) == 0 {
		return 0, 0, 0
	}

	var min, max int = 0, 0
	rate := getRate(s.Grades)
	min = getMin(s.Grades)
	max = getMax(s.Grades)
	return rate, min, max
}

// gunakan untuk melakukan debugging
func main() {
	avg, min, max := Analysis(School{
		Name:    "Imam Assidiqi School",
		Address: "Jl. Imam Assidiqi",
		Grades:  []int{},
	})

	fmt.Println(avg, min, max)
}
