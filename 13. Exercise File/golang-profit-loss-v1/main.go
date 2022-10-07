package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Readfile(path string) ([]string, error) {
	file, err := os.Open(path)
	defer file.Close()

	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var contentLines []string

	for scanner.Scan() {
		contentLines = append(contentLines, scanner.Text())
	}

	if len(contentLines) == 0 {
		return []string{}, nil
	}

	return contentLines, nil
}

func CalculateProfitLoss(data []string) string {
	if len(data) == 0 {
		return ""
	}
	var result string
	var profitOrLoss int

	var lastDate string

	for _, line := range data {
		// ["4/1/2021","income","10000"]
		stringData := strings.Split(line, ";")
		value, _ := strconv.Atoi(stringData[2])
		if stringData[1] == "income" || stringData[1] == "profit" {
			profitOrLoss = profitOrLoss + value
		} else if stringData[1] == "expense" || stringData[1] == "loss" {
			profitOrLoss = profitOrLoss - value
		}

		lastDate = stringData[0]
	}

	if profitOrLoss > 0 {
		result = fmt.Sprintf("%s;%s;%d", lastDate, "profit", profitOrLoss)
	} else {
		result = fmt.Sprintf("%s;%s;%d", lastDate, "loss", -profitOrLoss)

	}

	return result // TODO: replace this
}

func main() {
	// bisa digunakan untuk pengujian
	datas, err := Readfile("transactions.txt")
	if err != nil {
		panic(err)
	}

	result := CalculateProfitLoss(datas)
	fmt.Println(result)
}
