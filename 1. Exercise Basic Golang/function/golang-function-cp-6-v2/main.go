package main

import (
	"fmt"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func ChangeToCurrency(change int) string {
	myCurr := message.NewPrinter(language.Indonesian)
	return myCurr.Sprintf("Rp. %d", change)
}

func MoneyChange(money int, listPrice ...int) string {
	var totalPrice int
	for index := 0; index < len(listPrice); index++ {
		totalPrice += listPrice[index]
	}

	if money >= totalPrice {
		return ChangeToCurrency(money - totalPrice)
	}
	return "Uang tidak cukup"
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(MoneyChange(100000, 50000, 10000, 10000, 5000, 5000))
}
