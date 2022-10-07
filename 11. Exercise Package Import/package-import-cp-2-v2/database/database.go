package database

import (
	"a21hc3NpZ25tZW50/entity"
	"errors"
)

var productData = []entity.Product{
	{Name: "Kaos Polos", Price: 50000},
	{Name: "Kaos sablon", Price: 70000},
	{Name: "Baju Batik", Price: 200000},
	{Name: "Celana hitam", Price: 150000},
	{Name: "Celana pendek", Price: 100000},
	{Name: "Sabuk", Price: 50000},
	{Name: "Topi", Price: 75000},
	{Name: "Galung Tangan", Price: 30000},
	{Name: "Sepatu", Price: 300000},
}

type DatabaseInterface interface {
	Load() ([]entity.CartItem, error)
	Save([]entity.CartItem) error
	GetProductData() []entity.Product
	GetProductByname(name string) (entity.Product, error)
}

type Database struct {
	Data []entity.CartItem
}

func NewDatabase() *Database {
	return &Database{}
}

func (s *Database) Load() ([]entity.CartItem, error) {
	return s.Data, nil
}

func (s *Database) Save(data []entity.CartItem) error {
	s.Data = data
	return nil
}

func (s *Database) GetProductData() []entity.Product {
	return productData
}

func (s *Database) GetProductByname(name string) (entity.Product, error) {
	var product entity.Product
	for _, p := range productData {
		if p.Name == name {
			return p, nil
		}
	}
	return product, errors.New("product not found")
}
