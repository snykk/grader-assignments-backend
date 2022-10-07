package invoice

import (
	"errors"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var FinanceTestCase = func() {
	// =====================================================================
	// FinanceInvoice RecordInvoice test data
	// =====================================================================
	Describe("FinanceInvoice", func() {
		When("date is empty string", func() {
			It("should return error 'invoice date is empty'", func() {
				financeInvoice := FinanceInvoice{
					Date:     "",
					Status:   PAID,
					Approved: true,
					Details:  []Detail{{"pembelian nota", 4000}, {"Pembelian alat tulis", 4000}},
				}
				_, err := financeInvoice.RecordInvoice()
				Expect(err).To(Equal(errors.New("invoice date is empty")))
			})
		})
		When("status is empty or not 'paid' or not 'unpaid'", func() {
			It("should return error 'invoice status is not valid'", func() {
				financeInvoice := FinanceInvoice{
					Date:     "01/01/2022",
					Status:   "",
					Approved: true,
					Details:  []Detail{{"pembelian nota", 4000}, {"Pembelian alat tulis", 4000}},
				}
				_, err := financeInvoice.RecordInvoice()
				Expect(err).To(Equal(errors.New("invoice status is not valid")))
			})
		})
		When("details is empty data", func() {
			It("should return error 'invoice details is empty'", func() {
				financeInvoice := FinanceInvoice{
					Date:     "01/01/2022",
					Status:   PAID,
					Approved: true,
					Details:  []Detail{},
				}
				_, err := financeInvoice.RecordInvoice()
				Expect(err).To(Equal(errors.New("invoice details is empty")))
			})
		})
		When("total in details is 0 or less than 0", func() {
			It("should return error 'total price is not valid'", func() {
				financeInvoice := FinanceInvoice{
					Date:     "01/01/2022",
					Status:   PAID,
					Approved: true,
					Details:  []Detail{{"pembelian nota", 0}, {"Pembelian alat tulis", -200}},
				}
				_, err := financeInvoice.RecordInvoice()
				Expect(err).To(Equal(errors.New("total price is not valid")))
			})
		})
		When("invoice data as needed", func() {
			It("should return return error and covert invoice to InvoiceData", func() {
				financeInvoice := FinanceInvoice{
					Date:     "01/01/2022",
					Status:   PAID,
					Approved: true,
					Details:  []Detail{{"pembelian nota", 4000}, {"Pembelian alat tulis", 4000}},
				}
				inv, err := financeInvoice.RecordInvoice()
				Expect(err).To(BeNil())
				Expect(inv).To(Equal(InvoiceData{
					Date:         "01-January-2022",
					TotalInvoice: 8000,
					Departemen:   Finance,
				}))
			})
		})
	})
}
