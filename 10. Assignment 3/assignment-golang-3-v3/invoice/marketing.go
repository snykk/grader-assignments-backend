package invoice

import (
	"a21hc3NpZ25tZW50/helper"
	"errors"
	"strings"
)

// Marketing invoice
type MarketingInvoice struct {
	Date        string
	StartDate   string
	EndDate     string
	PricePerDay int
	AnotherFee  int
	Approved    bool
}

func (mi MarketingInvoice) RecordInvoice() (InvoiceData, error) {
	var err error
	if mi.Date == "" {
		err = errors.New("invoice date is empty")
		return InvoiceData{}, err
	}

	if mi.StartDate == "" || mi.EndDate == "" {
		err = errors.New("travel date is empty")
		return InvoiceData{}, err
	}

	if mi.PricePerDay <= 0 {
		err = errors.New("price per day is not valid")
		return InvoiceData{}, err
	}

	startDateList := strings.Split(mi.StartDate, "/")
	endDateList := strings.Split(mi.EndDate, "/")

	t1 := helper.GetDate(startDateList[2], startDateList[1], startDateList[0])
	t2 := helper.GetDate(endDateList[2], endDateList[1], endDateList[0])
	dayDiff := t2.Sub(t1).Hours() / 24

	totalPrice := (dayDiff+1)*float64(mi.PricePerDay) + float64(mi.AnotherFee)

	result := InvoiceData{
		Date:         helper.DateFormatter(mi.Date),
		Departemen:   Marketing,
		TotalInvoice: float64(totalPrice),
	}

	return result, nil // TODO: replace this
}
