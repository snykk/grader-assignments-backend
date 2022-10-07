package helper

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func DateFormatter(date string) string {
	mapHelper := map[string]string{
		"01": "January",
		"02": "February",
		"03": "March",
		"04": "April",
		"05": "May",
		"06": "June",
		"07": "July",
		"08": "August",
		"09": "September",
		"10": "October",
		"11": "November",
		"12": "December",
	}
	var result string
	dateList := strings.Split(date, "/")

	result = fmt.Sprintf("%s-%s-%s", dateList[0], mapHelper[dateList[1]], dateList[2])

	return result
}

func GetDate(year, month, day string) time.Time {
	var yearInt, monthInt, dayInt int
	yearInt, _ = strconv.Atoi(year)
	monthInt, _ = strconv.Atoi(month)
	dayInt, _ = strconv.Atoi(day)

	return time.Date(yearInt, time.Month(monthInt), dayInt, 0, 0, 0, 0, time.UTC)
}
