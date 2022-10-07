package invoice

import (
	"errors"
)

// Warehouse invoice

type WarehouseInvoice struct {
	Date        string
	InvoiceType InvoiceTypeName
	Approved    bool
	Products    []Product
}

type InvoiceTypeName string

const (
	PURCHASE InvoiceTypeName = "purchase"
	SALES    InvoiceTypeName = "sales"
)

type Product struct {
	Name     string
	Unit     int
	Price    float64
	Discount float64
}

func (wi WarehouseInvoice) RecordInvoice() (InvoiceData, error) {
	var err error
	if wi.Date == "" {
		err = errors.New("invoice date is empty")
		return InvoiceData{}, err
	}

	if wi.InvoiceType == "" || (wi.InvoiceType != "purchase" && wi.InvoiceType != "sales") {
		err = errors.New("invoice type is not valid")
		return InvoiceData{}, err
	}

	if len(wi.Products) == 0 {
		err = errors.New("invoice products is empty")
		return InvoiceData{}, err
	}

	var totalPrice float64
	for _, product := range wi.Products {
		if product.Unit <= 0 {
			err = errors.New("unit product is not valid")
			return InvoiceData{}, err
		}

		if product.Price <= 0 {
			err = errors.New("price product is not valid")
			return InvoiceData{}, err
		}

		totalPrice += (1 - product.Discount) * float64(product.Unit) * float64(product.Price)

	}

	result := InvoiceData{
		Date:         wi.Date,
		Departemen:   Warehouse,
		TotalInvoice: float64(totalPrice),
	}

	return result, nil // TODO: replace this
}
