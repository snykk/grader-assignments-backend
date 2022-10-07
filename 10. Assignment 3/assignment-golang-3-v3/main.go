package main

import (
	"a21hc3NpZ25tZW50/invoice"
	"fmt"
	"log"
)

func RecapDataInvoice(data []invoice.Invoice) ([]invoice.InvoiceData, error) {
	var result []invoice.InvoiceData
	var eachData invoice.InvoiceData
	var err error
	var isIncrement bool

	for _, item := range data {
		eachData, err = item.RecordInvoice()

		for index, result_item := range result {
			isIncrement = false
			if result_item.Date == eachData.Date && result_item.Departemen == eachData.Departemen {
				result[index].TotalInvoice += result_item.TotalInvoice
				isIncrement = true
				break
			}
		}

		if isIncrement == false {
			result = append(result, eachData)
		}
	}

	return result, err
}

func main() {
	// listInvoice := []invoice.Invoice{
	// 	invoice.FinanceInvoice{
	// 		Date: "01/02/2020",
	// 		Details: []invoice.Detail{
	// 			{
	// 				Description: "pembelian nota",
	// 				Total:       4000,
	// 			}, {
	// 				Description: "Pembelian alat tulis",
	// 				Total:       4000,
	// 			},
	// 		},
	// 		Status:   invoice.PAID,
	// 		Approved: true,
	// 	},
	// 	invoice.FinanceInvoice{
	// 		Date: "01/02/2020",
	// 		Details: []invoice.Detail{
	// 			{
	// 				Description: "pembelian nota",
	// 				Total:       4000,
	// 			}, {
	// 				Description: "Pembelian alat tulis",
	// 				Total:       4000,
	// 			},
	// 		},
	// 		Status:   invoice.PAID,
	// 		Approved: true,
	// 	},
	// 	invoice.WarehouseInvoice{
	// 		Date: "01-February-2020",
	// 		Products: []invoice.Product{
	// 			{
	// 				Name:     "product A",
	// 				Unit:     10,
	// 				Price:    10000,
	// 				Discount: 0.1,
	// 			},
	// 			{
	// 				Name:     "product C",
	// 				Unit:     5,
	// 				Price:    15000,
	// 				Discount: 0.2,
	// 			},
	// 		},
	// 		InvoiceType: invoice.PURCHASE,
	// 		Approved:    true,
	// 	},
	// 	invoice.MarketingInvoice{
	// 		Date:        "01/02/2020",
	// 		StartDate:   "20/01/2020",
	// 		EndDate:     "25/01/2020",
	// 		Approved:    true,
	// 		PricePerDay: 10000,
	// 		AnotherFee:  5000,
	// 	},
	// }

	listInvoice := []invoice.Invoice{
		invoice.FinanceInvoice{
			Date:     "01/02/2020",
			Details:  []invoice.Detail{{Description: "pembelian nota", Total: 4000}, {Description: "Pembelian alat tulis", Total: 4000}},
			Status:   invoice.PAID,
			Approved: true,
		},
		invoice.FinanceInvoice{
			Date:     "01/02/2020",
			Details:  []invoice.Detail{{Description: "pembelian nota", Total: 4000}, {Description: "Pembelian alat tulis", Total: 4000}},
			Status:   invoice.PAID,
			Approved: true,
		},
		invoice.WarehouseInvoice{
			Date: "01-February-2020",
			Products: []invoice.Product{
				{Name: "product A", Unit: 10, Price: 10000, Discount: 0.1},
				{Name: "product C", Unit: 5, Price: 15000, Discount: 0.2},
			},
			InvoiceType: invoice.PURCHASE,
			Approved:    true,
		},
		invoice.MarketingInvoice{
			Date:        "01/02/2020",
			StartDate:   "20/01/2020",
			EndDate:     "25/01/2020",
			Approved:    true,
			PricePerDay: 10000,
			AnotherFee:  5000,
		},
	}

	result, err := RecapDataInvoice(listInvoice)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
}
