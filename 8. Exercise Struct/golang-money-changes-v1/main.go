package main

import "fmt"

type Product struct {
	Name  string
	Price int
	Tax   int
}

// ===>
func ExchangeCoin(cashBack int) []int {
	var cointList = []int{}
	for cashBack > 0 {
		if cashBack >= 1000 {
			cointList = append(cointList, 1000)
			cashBack -= 1000
		} else if cashBack >= 500 {
			cointList = append(cointList, 500)
			cashBack -= 500
		} else if cashBack >= 200 {
			cointList = append(cointList, 200)
			cashBack -= 200
		} else if cashBack >= 100 {
			cointList = append(cointList, 100)
			cashBack -= 100
		} else if cashBack >= 50 {
			cointList = append(cointList, 50)
			cashBack -= 50
		} else if cashBack >= 20 {
			cointList = append(cointList, 20)
			cashBack -= 20
		} else if cashBack >= 10 {
			cointList = append(cointList, 10)
			cashBack -= 10
		} else if cashBack >= 5 {
			cointList = append(cointList, 5)
			cashBack -= 5
		} else if cashBack >= 1 {
			cointList = append(cointList, 1)
			cashBack -= 1
		}
	}
	return cointList
}

func MoneyChanges(amount int, products []Product) []int {
	var totalPrice int

	for _, prodObj := range products {
		totalPrice += prodObj.Price + prodObj.Tax
	}

	cashBack := ExchangeCoin(amount - totalPrice)

	return cashBack
}

func main() {
	data := []Product{
		{
			Name:  "Baju",
			Price: 5000,
			Tax:   500,
		},
		{
			Name:  "Celana",
			Price: 3000,
			Tax:   300,
		},
	}

	var amount int = 10000

	fmt.Println(MoneyChanges(amount, data))
}
