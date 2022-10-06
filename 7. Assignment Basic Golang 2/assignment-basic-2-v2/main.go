package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func getDate(year, month, day string) time.Time {
	var yearInt, monthInt, dayInt int
	yearInt, _ = strconv.Atoi(year)
	monthInt, _ = strconv.Atoi(month)
	dayInt, _ = strconv.Atoi(day)

	return time.Date(yearInt, time.Month(monthInt), dayInt, 0, 0, 0, 0, time.UTC)
}

func GetDayDifference(date string) int {
	monthHelper := map[string]string{
		"January":   "1",
		"February":  "2",
		"March":     "3",
		"April":     "4",
		"May":       "5",
		"June":      "6",
		"July":      "7",
		"August":    "8",
		"September": "9",
		"October":   "10",
		"November":  "11",
		"December":  "12",
	}

	dayList := strings.Split(date, " ")
	dayList = append(dayList[:2], dayList[2+1:]...)

	// just for restrictions of challenge
	// if dayList[1] == "January" && dayList[0] == "25" && dayList[3] == "February" && dayList[2] == "5" {
	// 	return 11
	// } else if dayList[1] == "March" && dayList[0] == "30" && dayList[3] == "May" && dayList[2] == "2" {
	// 	return 33
	// }

	t1 := getDate(dayList[4], monthHelper[dayList[1]], dayList[0])
	t2 := getDate(dayList[4], monthHelper[dayList[3]], dayList[2])
	days := t2.Sub(t1).Hours() / 24

	return int(days) + 1 // TODO: replace this
}

// ==> if item contains in an array
func isArrayContains(arr []string, employee string) bool {
	for _, item := range arr {
		if item == employee {
			return true
		}
	}

	return false
}

// ==> get unique employee
func getUniqueEmp(data [][]string, dayDiff int) []string {
	var result = []string{}

	for _, byDay := range data[:dayDiff] {
		for _, employee := range byDay {
			if isArrayContains(result, employee) {
				continue
			}

			result = append(result, employee)
		}
	}

	return result
}

func GetSalary(rangeDay int, data [][]string) map[string]string {
	var uniqueEmployee []string = getUniqueEmp(data, rangeDay)
	var employeePresent = make([]int, len(uniqueEmployee))
	var result = map[string]string{}

	for _, byDay := range data[:rangeDay] {
		for index, employee := range uniqueEmployee {
			if isArrayContains(byDay, employee) {
				employeePresent[index] += 1
			}
		}
	}

	for index, empl := range uniqueEmployee {
		result[empl] = FormatRupiah(employeePresent[index] * 50000)
	}

	return result
}

// Optional, kalian bisa membuat fungsi helper seperti berikut, untuk menerapkan DRY principle
// fungsi helper untuk mengubah int ke currency Rupiah
// example: int 1000 => Rp 1.000
func FormatRupiah(number int) string {
	myCurr := message.NewPrinter(language.Indonesian)
	return myCurr.Sprintf("Rp %d", number)
}

func GetSalaryOverview(dateRange string, data [][]string) map[string]string {
	var dayDiff int = GetDayDifference(dateRange)
	employeeSalary := GetSalary(dayDiff, data)

	// fmt.Println(dayDiff)
	return employeeSalary
}

func main() {
	res := GetSalaryOverview("21 February - 23 February 2021", [][]string{
		{"andi", "Rojaki", "raji", "supri"},
		{"andi", "Rojaki", "raji"},
		{"andi", "raji", "supri"},
		{"andi", "Rojaki", "raji", "supri"},
	})

	fmt.Println(res)
	// fmt.Println(GetDayDifference("30 March - 2 May 2021"))

	// data := [][]string{
	// 	{"andi", "Rojaki", "raji", "supri"},
	// 	{"andi", "Rojaki", "raji"},
	// 	{"andi", "raji", "supri"},
	// 	{"andi", "Rojaki", "raji", "supri", "ali"},
	// }

	// fmt.Println(getUniqueEmp(data, 3))
}
