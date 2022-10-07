package invoice

import (
	"errors"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var MarketingTestCase = func() {
	// =====================================================================
	// MarketingInvoice RecordInvoice test data
	// =====================================================================
	Describe("MarketingInvoice", func() {
		When("date is empty string", func() {
			It("should return error 'invoice date is empty'", func() {
				MarketingInvoice := MarketingInvoice{
					Date:     "",
					Approved: true,
				}
				_, err := MarketingInvoice.RecordInvoice()
				Expect(err).To(Equal(errors.New("invoice date is empty")))
			})
		})

		When("start date and end date is emptu", func() {
			It("should return error 'travel date is empty'", func() {
				MarketingInvoice := MarketingInvoice{
					Date:      "01/01/2022",
					StartDate: "",
					EndDate:   "",
					Approved:  true,
				}
				_, err := MarketingInvoice.RecordInvoice()
				Expect(err).To(Equal(errors.New("travel date is empty")))
			})
		})

		When("PricePerDay is 0 or less than 0", func() {
			It("should return error 'price per day is not valid'", func() {
				MarketingInvoice := MarketingInvoice{
					Date:        "01/01/2022",
					StartDate:   "01/01/2022",
					EndDate:     "02/01/2022",
					PricePerDay: -1000,
					Approved:    true,
				}
				_, err := MarketingInvoice.RecordInvoice()
				Expect(err).To(Equal(errors.New("price per day is not valid")))
			})
		})

		When("invoice marketing data as needed", func() {
			It("should not return error and convert data to Invoice Data", func() {
				MarketingInvoice := MarketingInvoice{
					Date:        "03/01/2022",
					StartDate:   "01/01/2022",
					EndDate:     "02/01/2022",
					PricePerDay: 10000,
					AnotherFee:  5000,
					Approved:    true,
				}
				inv, err := MarketingInvoice.RecordInvoice()
				Expect(err).To(BeNil())
				Expect(inv).To(Equal(InvoiceData{
					Date:         "03-January-2022",
					TotalInvoice: 25000,
					Departemen:   Marketing,
				}))
			})
		})
	})
}
