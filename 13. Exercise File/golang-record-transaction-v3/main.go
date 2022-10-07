package main

import (
	"fmt"
	"math"
	"os"
)

type Transaction struct {
	Date   string
	Type   string
	Amount int
}

func insertContent(path, content string, isLastItem bool) error {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)

	if err != nil {
		panic(err)
	}

	defer f.Close()

	var needNewLine string

	if !isLastItem {
		needNewLine = "\n"
	}

	if _, err = f.WriteString(content + needNewLine); err != nil {
		return err
	}

	return nil
}

func RecordTransactions(path string, transactions []Transaction) error {
	err := os.WriteFile(path, []byte(""), 0644)

	if err != nil {
		panic(err)
	}

	var result []Transaction
	var isSameDate bool

	for _, item := range transactions {
		isSameDate = false
		for index, itemResult := range result {
			if itemResult.Date == item.Date {
				isSameDate = true
				if item.Type == "income" {
					result[index].Amount += item.Amount
				} else {
					result[index].Amount -= item.Amount
				}

				if result[index].Amount < 0 {
					result[index].Type = "expense"
				} else {
					result[index].Type = "income"
				}
			}

		}

		if !isSameDate {
			var newAmount int = item.Amount
			if item.Type == "expense" {
				newAmount = item.Amount * -1
			}

			trasStruc := Transaction{
				Date:   item.Date,
				Type:   item.Type,
				Amount: newAmount,
			}
			result = append(result, trasStruc)
		}
	}

	for index, item := range result {
		str := fmt.Sprintf("%s;%s;%d", item.Date, item.Type, int(math.Abs(float64(item.Amount))))

		if index == len(result)-1 {
			insertContent(path, str, true)
			continue
		}

		insertContent(path, str, false)
	}

	return nil // TODO: replace this
}

func main() {
	// bisa digunakan untuk pengujian test case
	// var transactions = []Transaction{
	// 	{"01/01/2021", "income", 100000},
	// 	{"01/01/2021", "expense", 50000},
	// 	{"01/01/2021", "expense", 30000},
	// 	{"01/01/2021", "income", 20000},
	// 	{"01/01/2021", "expense", 70000},
	// }

	var transactions = []Transaction{
		{"01/01/2021", "expense", 100000},
		{"01/01/2021", "expense", 100000},
		{"01/01/2021", "expense", 100000},
		{"01/01/2021", "expense", 100000},
	}

	err := RecordTransactions("transactions.txt", transactions)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success")
}
