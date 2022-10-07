package invoice

import (
	"a21hc3NpZ25tZW50/helper"
	"errors"
)

// Finance invoice
type FinanceInvoice struct {
	Date     string
	Status   InvoiceStatus // status: "paid", "unpaid"
	Approved bool
	Details  []Detail
}

type InvoiceStatus string

const (
	PAID   InvoiceStatus = "paid"
	UNPAID InvoiceStatus = "unpaid"
)

type Detail struct {
	Description string
	Total       int
}

func (fi FinanceInvoice) RecordInvoice() (InvoiceData, error) {
	var err error
	if fi.Date == "" {
		err = errors.New("invoice date is empty")
		return InvoiceData{}, err
	}

	if len(fi.Details) == 0 {
		err = errors.New("invoice details is empty")
		return InvoiceData{}, err
	}

	if fi.Status == "" || (fi.Status != "paid" && fi.Status != "unpaid") {
		err = errors.New("invoice status is not valid")
		return InvoiceData{}, err
	}

	var totalPrice int

	for _, detail := range fi.Details {
		if detail.Total <= 0 {
			err = errors.New("total price is not valid")
			return InvoiceData{}, err
		}

		totalPrice += detail.Total

	}

	// getTotalInvoice

	result := InvoiceData{
		Date:         helper.DateFormatter(fi.Date),
		Departemen:   Finance,
		TotalInvoice: float64(totalPrice),
	}

	return result, nil // TODO: replace this
}
