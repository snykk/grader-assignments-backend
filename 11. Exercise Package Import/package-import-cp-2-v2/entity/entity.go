package entity

type Product struct {
	Name  string
	Price int
}

type CartItem struct {
	ProductName string
	Price       int
	Quantity    int
}

type PaymentInformation struct {
	ListProduct []CartItem
	TotalPrice  int
	MoneyPaid   int
	Change      int
}
