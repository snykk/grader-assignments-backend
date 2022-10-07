package main_test

import (
	main "a21hc3NpZ25tZW50"
	"a21hc3NpZ25tZW50/invoice"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("RecordInvoice", func() {
	Describe("Marketing", invoice.MarketingTestCase)
	Describe("Finance", invoice.FinanceTestCase)
	Describe("Warehouse", invoice.WarehouseTestCase)
})

var _ = Describe("RecapDataInvoice", func() {
	// test same date same departemen
	When("all invoice data is same date and same departemen", func() {
		It("should not return error, and return one data invoice", func() {
			invoices := []invoice.Invoice{
				invoice.FinanceInvoice{
					Date: "01/02/2020",
					Details: []invoice.Detail{
						{Description: "pembelian nota", Total: 4000},
						{Description: "Pembelian alat tulis", Total: 4000}},
					Status:   invoice.PAID,
					Approved: true,
				},
				invoice.FinanceInvoice{
					Date: "01/02/2020",
					Details: []invoice.Detail{
						{Description: "pembelian nota", Total: 4000},
						{Description: "Pembelian alat tulis", Total: 4000}},
					Status:   invoice.PAID,
					Approved: true,
				},
			}

			listInv, err := main.RecapDataInvoice(invoices)
			Expect(err).To(BeNil())

			Expect(listInv).To(Equal([]invoice.InvoiceData{
				{Date: "01-February-2020", TotalInvoice: 16000, Departemen: invoice.Finance},
			}))
		})
	})
	// test same date different departemen
	When("all invoice data is same date and difference departemen", func() {
		It("should not return error, and return list data by departement", func() {
			invoices := []invoice.Invoice{
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

			listInv, err := main.RecapDataInvoice(invoices)
			Expect(err).To(BeNil())

			expected := []invoice.InvoiceData{
				{Date: "01-February-2020", TotalInvoice: 16000, Departemen: invoice.Finance},
				{Date: "01-February-2020", TotalInvoice: 65000, Departemen: invoice.Marketing},
				{Date: "01-February-2020", TotalInvoice: 150000, Departemen: invoice.Warehouse},
			}

			for _, inv := range listInv {
				Expect(expected).To(ContainElement(inv))
			}
		})
	})

	// test different date same departemen
	When("all invoice data with differece date and difference departemen", func() {
		It("should not return error, and return list data by date and by departement", func() {
			invoices := []invoice.Invoice{
				invoice.FinanceInvoice{
					Date:     "01/02/2020",
					Details:  []invoice.Detail{{Description: "pembelian nota", Total: 4000}, {Description: "Pembelian alat tulis", Total: 4000}},
					Status:   invoice.PAID,
					Approved: true,
				},
				invoice.FinanceInvoice{
					Date:     "02/02/2020",
					Details:  []invoice.Detail{{Description: "pembelian nota", Total: 4000}, {Description: "Pembelian alat tulis", Total: 4000}},
					Status:   invoice.PAID,
					Approved: true,
				},
				invoice.WarehouseInvoice{
					Date: "03-February-2020",
					Products: []invoice.Product{
						{Name: "product A", Unit: 10, Price: 10000, Discount: 0.1},
						{Name: "product C", Unit: 5, Price: 15000, Discount: 0.2},
					},
					InvoiceType: invoice.PURCHASE,
					Approved:    true,
				},
				invoice.MarketingInvoice{
					Date:        "04/02/2020",
					StartDate:   "20/01/2020",
					EndDate:     "25/01/2020",
					Approved:    true,
					PricePerDay: 10000,
					AnotherFee:  5000,
				},
			}

			listInv, err := main.RecapDataInvoice(invoices)
			Expect(err).To(BeNil())

			expected := []invoice.InvoiceData{
				{Date: "01-February-2020", TotalInvoice: 8000, Departemen: invoice.Finance},
				{Date: "02-February-2020", TotalInvoice: 8000, Departemen: invoice.Finance},
				{Date: "03-February-2020", TotalInvoice: 150000, Departemen: invoice.Warehouse},
				{Date: "04-February-2020", TotalInvoice: 65000, Departemen: invoice.Marketing},
			}

			for _, inv := range listInv {
				Expect(expected).To(ContainElement(inv))
			}
		})
	})
})
