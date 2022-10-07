package main

import (
	"a21hc3NpZ25tZW50/database"
	"a21hc3NpZ25tZW50/service"
	"fmt"
)

func CashierApp(db *database.Database) service.ServiceInterface {
	service := service.NewService(db)

	return service
}

// gunakan untuk debugging
func main() {
	database := database.NewDatabase()
	service := CashierApp(database)

	service.AddCart("Kaos Polos", 2)
	service.AddCart("Kaos sablon", 1)
	service.RemoveCart("Kaos sablon")

	myCart, _ := service.ShowCart()
	fmt.Println(myCart)

	paymentInformation, err := service.Paid(500000)
	if err != nil {
		panic(err)
	}

	fmt.Println(paymentInformation)
}
