package main

import (
	"fmt"
	"strconv"
	"strings"
)

// ==> wheter item is contains in an array or not
func isArrayContains(arr []string, day string) bool {
	for _, v := range arr {
		if v == day {
			return true
		}
	}

	return false
}

// ==> wheter day is available or not
func isDayAvailable(code string, day string) bool {
	helper := map[string][]string{
		"JKT": {"senin", "selasa", "rabu", "kamis", "jumat", "sabtu"},
		"BDG": {"rabu", "kamis", "sabtu"},
		"BKS": {"selasa", "kamis", "jumat"},
		"DPK": {"senin", "selasa"},
	}

	if isArrayContains(helper[code], day) {
		return true
	}

	return false
}

// ==> to get shipping cost
func addShippingCost(day string, num float64) float64 {
	switch day {
	case "senin", "rabu", "jumat":
		return 0.1
	case "selasa", "kamis", "sabtu":
		return 0.05
	}

	return 0
}

func DeliveryOrder(data []string, day string) map[string]float32 {
	var dataList = []string{}
	var fullname string
	var delivSummary = map[string]float32{}

	for _, item := range data {
		dataList = strings.Split(item, ":")

		if isDayAvailable(dataList[3], day) {
			fullname = fmt.Sprintf("%s-%s", dataList[0], dataList[1])
			float, _ := strconv.ParseFloat(dataList[2], 32)
			float = (1 + addShippingCost(day, float)) * float
			if delivSummary[fullname] != 0 {
				delivSummary[fullname] = delivSummary[fullname] + float32(float)

				continue
			}

			delivSummary[fullname] = float32(float)
		}

	}
	return delivSummary
}

func main() {
	data := []string{
		"Budi:Gunawan:10000:JKT",
		"Andi:Sukirman:20000:JKT",
		"Budi:Sukirman:30000:BDG",
		"Andi:Gunawan:40000:BKS",
		"Budi:Gunawan:50000:DPK",
	}

	day := "sabtu"

	deliveryData := DeliveryOrder(data, day)

	fmt.Println(deliveryData)
}
