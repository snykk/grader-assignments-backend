package service

import (
	"a21hc3NpZ25tZW50/database"
	"a21hc3NpZ25tZW50/entity"
	"errors"
)

// Service is package for any logic needed in this program

type ServiceInterface interface {
	AddCart(productName string, quantity int) error
	RemoveCart(productName string) error
	ShowCart() ([]entity.CartItem, error)
	ResetCart() error
	GetAllProduct() ([]entity.Product, error)
	Paid(money int) (entity.PaymentInformation, error)
}

type Service struct {
	database database.DatabaseInterface
}

func NewService(database database.DatabaseInterface) *Service {
	return &Service{
		database: database,
	}
}

func (s *Service) AddCart(productName string, quantity int) error {
	var price int
	var isGetProduct bool
	loadData, err := s.database.Load()

	if err != nil {
		return err
	}

	if quantity <= 0 {
		err = errors.New("invalid quantity")
		return err
	}

	for _, item := range s.database.GetProductData() {
		if item.Name == productName {
			isGetProduct = true
			price = item.Price
			break
		}
	}

	if isGetProduct != true {
		err = errors.New("product not found")
		return err
	}

	newCartItem := entity.CartItem{
		ProductName: productName,
		Quantity:    quantity,
		Price:       price,
	}

	loadData = append(loadData, newCartItem)

	err = s.database.Save(loadData)

	if err != nil {
		return err
	}

	return nil // TODO: replace this
}

func (s *Service) RemoveCart(productName string) error {
	var newCart = []entity.CartItem{}
	loadCart, err := s.database.Load()

	if err != nil {
		return err
	}

	for _, cartItem := range loadCart {
		if cartItem.ProductName == productName {
			continue
		}

		newCartItem := entity.CartItem{
			ProductName: cartItem.ProductName,
			Price:       cartItem.Price,
			Quantity:    cartItem.Quantity,
		}

		newCart = append(newCart, newCartItem)
	}

	if len(loadCart) == len(newCart) {
		err = errors.New("product not found")
		return err
	}

	err = s.database.Save(newCart)

	if err != nil {
		return err
	}

	return nil // TODO: replace this
}

func (s *Service) ShowCart() ([]entity.CartItem, error) {
	carts, err := s.database.Load()
	if err != nil {
		return nil, err
	}

	return carts, nil
}

func (s *Service) ResetCart() error {
	newEmptyCart := make([]entity.CartItem, 0, 0)

	err := s.database.Save(newEmptyCart)

	if err != nil {
		return err
	}

	return nil // TODO: replace this
}

func (s *Service) GetAllProduct() ([]entity.Product, error) {
	return s.database.GetProductData(), nil
}

func (s *Service) Paid(money int) (entity.PaymentInformation, error) {
	loadCart, err := s.database.Load()
	var totalPrice int

	if err != nil {
		return entity.PaymentInformation{}, err
	}

	for _, cartItem := range loadCart {
		totalPrice += cartItem.Price * cartItem.Quantity
	}

	if money < totalPrice {
		err = errors.New("money is not enough")
		return entity.PaymentInformation{}, err
	}

	paymentInformation := entity.PaymentInformation{
		ListProduct: loadCart,
		TotalPrice:  totalPrice,
		MoneyPaid:   money,
		Change:      money - totalPrice,
	}

	s.ResetCart()

	return paymentInformation, nil // TODO: replace this
}
