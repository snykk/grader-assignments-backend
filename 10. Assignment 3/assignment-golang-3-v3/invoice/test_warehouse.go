package invoice

import (
	"errors"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var WarehouseTestCase = func() {
	// =====================================================================
	// WarehouseInvoice RecordInvoice test data
	// =====================================================================
	Describe("WarehouseInvoice", func() {
		When("date is empty string", func() {
			It("should return error 'invoice date is empty'", func() {
				warehouseInvoice := WarehouseInvoice{
					Date:        "",
					InvoiceType: PURCHASE,
					Approved:    true,
					Products: []Product{
						{"product A", 10, 10000, 0.1},
						{"product C", 5, 15000, 0.2},
					},
				}
				_, err := warehouseInvoice.RecordInvoice()
				Expect(err).To(Equal(errors.New("invoice date is empty")))
			})
		})

		When("invoice type is empty string or not 'pruchase' or not 'sales'", func() {
			It("should return error 'invoice type is not valid'", func() {
				warehouseInvoice := WarehouseInvoice{
					Date:        "01-January-2022",
					InvoiceType: "",
					Approved:    true,
					Products: []Product{
						{"product A", 10, 10000, 0.1},
						{"product C", 5, 15000, 0.2},
					},
				}
				_, err := warehouseInvoice.RecordInvoice()
				Expect(err).To(Equal(errors.New("invoice type is not valid")))
			})
		})

		When("invoice product is empty data", func() {
			It("should return error 'invoice products is empty'", func() {
				warehouseInvoice := WarehouseInvoice{
					Date:        "01-January-2022",
					InvoiceType: "purchase",
					Approved:    true,
					Products:    []Product{},
				}
				_, err := warehouseInvoice.RecordInvoice()
				Expect(err).To(Equal(errors.New("invoice products is empty")))
			})
		})

		When("invoice product unit is 0 or lower than 0", func() {
			It("should return error 'unit product is not valid'", func() {
				warehouseInvoice := WarehouseInvoice{
					Date:        "01-January-2022",
					InvoiceType: "purchase",
					Approved:    true,
					Products: []Product{
						{"product A", -1, 10000, 0.1},
						{"product C", 0, 15000, 0.2},
					},
				}
				_, err := warehouseInvoice.RecordInvoice()
				Expect(err).To(Equal(errors.New("unit product is not valid")))
			})
		})

		When("invoice price product is 0 or less than 0", func() {
			It("should return error 'price product is not valid'", func() {
				warehouseInvoice := WarehouseInvoice{
					Date:        "01-January-2022",
					InvoiceType: "purchase",
					Approved:    true,
					Products: []Product{
						{"product A", 10, -1, 0.1},
						{"product C", 5, 0, 0.2},
					},
				}
				_, err := warehouseInvoice.RecordInvoice()
				Expect(err).To(Equal(errors.New("price product is not valid")))
			})
		})

		When("invoice warehouse data as needed", func() {
			It("should not return error and convert data to Invoice Data", func() {
				warehouseInvoice := WarehouseInvoice{
					Date:        "01-January-2022",
					InvoiceType: "purchase",
					Approved:    true,
					Products: []Product{
						{"product A", 10, 10000, 0.1},
						{"product C", 5, 15000, 0.2},
					},
				}
				inv, err := warehouseInvoice.RecordInvoice()
				Expect(err).To(BeNil())
				Expect(inv).To(Equal(InvoiceData{
					Date:         "01-January-2022",
					TotalInvoice: 150000,
					Departemen:   Warehouse,
				}))
			})
		})
	})

}
