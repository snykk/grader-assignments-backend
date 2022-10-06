package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Time struct {
	Hour   int
	Minute int
}

func dateFormatter(hour int, minute int) string {
	var result string
	if hour == 12 {
		result = fmt.Sprintf("%02d:%02d PM", hour, minute)
	} else if hour == 24 {
		result = fmt.Sprintf("%02d:%02d AM", hour-12, minute)
	} else if hour >= 12 {
		result = fmt.Sprintf("%02d:%02d PM", hour-12, minute)

	} else {
		result = fmt.Sprintf("%02d:%02d AM", hour, minute)
	}

	return result

}

func ChangeToStandartTime(time interface{}) string {
	var result string

	timeList := []string{}
	switch time.(type) {
	case string:
		timeList = strings.Split(time.(string), ":")
		if len(timeList) != 2 {
			result = "Invalid input"
			break
		}

		if timeList[0] == "" || timeList[1] == "" {
			result = "Invalid input"
			break
		}

		hour, _ := strconv.Atoi(timeList[0])
		minute, _ := strconv.Atoi(timeList[1])
		result = dateFormatter(hour, minute)
	case []int:
		if len(time.([]int)) != 2 {
			result = "Invalid input"
			break
		}

		result = dateFormatter(time.([]int)[0], time.([]int)[1])
	case map[string]int:
		_, isHourPresent := time.(map[string]int)["hour"]
		_, isMinutePresent := time.(map[string]int)["minute"]
		if !isHourPresent || !isMinutePresent {
			result = "Invalid input"
			break
		}

		result = dateFormatter(time.(map[string]int)["hour"], time.(map[string]int)["minute"])
	case Time:
		result = dateFormatter(time.(Time).Hour, time.(Time).Minute)
	default:
		result = "Invalid input"
	}

	return result
}

func main() {
	fmt.Println(ChangeToStandartTime("12:00"))
	fmt.Println(ChangeToStandartTime([]int{16, 0}))
	fmt.Println(ChangeToStandartTime(map[string]int{"hour": 16, "minute": 0}))
	fmt.Println(ChangeToStandartTime(Time{16, 0}))
}
