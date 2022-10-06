package main

import (
	"fmt"
	"strconv"
	"strings"
)

func PopulationData(data []string) []map[string]interface{} {
	if len(data) == 0 {
		x := make([]map[string]interface{}, 0)
		return x
	}

	var dataList = []string{}
	var result []map[string]interface{}
	mapPresenter := make(map[string]interface{})
	for _, personData := range data {
		dataList = strings.Split(personData, ";")
		mapPresenter["name"] = dataList[0]
		mapPresenter["age"], _ = strconv.Atoi(dataList[1])
		mapPresenter["address"] = dataList[2]
		if dataList[3] != "" {
			mapPresenter["height"], _ = strconv.ParseFloat(dataList[3], 2)
		}

		if dataList[4] != "" {
			mapPresenter["isMarried"], _ = strconv.ParseBool(dataList[4])
		}

		result = append(result, mapPresenter)

		mapPresenter = make(map[string]interface{})
	}

	return result
}

func main() {
	data := []string{}

	fmt.Println(PopulationData(data))
}
