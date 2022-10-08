package main_test

import (
	main "a21hc3NpZ25tZW50"
	"io/ioutil"
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

const PATH_TEST_FILE = "test_loan-records.json"

var BeforeEachHandler = func() {
	// do something before each test
	if _, err := os.Stat(PATH_TEST_FILE); os.IsNotExist(err) {
		var f *os.File
		if f, err = os.Create(PATH_TEST_FILE); err != nil {
			panic(err)
		}
		defer f.Close()
	}
}

var AssertRecord = func() string {
	if data, err := ioutil.ReadFile(PATH_TEST_FILE); err == nil {
		return string(data)
	}
	return ""
}

var AfterEachHandler = func() {
	os.Remove(PATH_TEST_FILE)
}

var _ = Describe("LoanReport", func() {
	When("StartBalance is more than total loan", func() {
		It("should return with EndBalance more than 0", func() {
			Expect(main.LoanReport(
				main.LoanData{
					StartBalance: 500000,
					Data: []main.Loan{
						{"01-January-2021", []string{"1", "2"}},
						{"02-January-2021", []string{"1", "2", "3"}},
						{"03-January-2021", []string{"2", "3"}},
						{"04-January-2021", []string{"1", "3"}},
					},
					Employees: []main.Employee{
						{"1", "Employee A", "Manager"},
						{"2", "Employee B", "Staff"},
						{"3", "Employee C", "Staff"},
					},
				})).To(Equal(
				main.LoanRecord{
					MonthDate:    "January 2021",
					StartBalance: 500000,
					EndBalance:   50000,
					Borrowers: []main.Borrower{
						{"1", 150000},
						{"2", 150000},
						{"3", 150000},
					},
				},
			))
		})
	})

	When("StartBalance is same as total loan", func() {
		It("should return with EndBalance is 0", func() {
			Expect(main.LoanReport(
				main.LoanData{
					StartBalance: 350000,
					Data: []main.Loan{
						{"01-March-2021", []string{"1", "2", "3", "4"}},
						{"04-March-2021", []string{"1", "2", "3"}},
					},
					Employees: []main.Employee{
						{"1", "Employee A", "Manager"},
						{"2", "Employee B", "Staff"},
						{"3", "Employee C", "Staff"},
						{"4", "Employee D", "Staff"},
					},
				})).To(Equal(
				main.LoanRecord{
					MonthDate:    "March 2021",
					StartBalance: 350000,
					EndBalance:   0,
					Borrowers: []main.Borrower{
						{"1", 100000},
						{"2", 100000},
						{"3", 100000},
						{"4", 50000},
					},
				},
			))
		})
	})

	When("StartBalance is smaller than total loan", func() {
		It("should return with EndBalance 0, and total loan not more than total balance", func() {
			Expect(main.LoanReport(
				main.LoanData{
					StartBalance: 250000,
					Data: []main.Loan{
						{"01-March-2021", []string{"1", "2", "3", "4"}},
						{"04-March-2021", []string{"1", "2", "3"}},
					},
					Employees: []main.Employee{
						{"1", "Employee A", "Manager"},
						{"2", "Employee B", "Staff"},
						{"3", "Employee C", "Staff"},
						{"4", "Employee D", "Staff"},
					},
				})).To(Equal(
				main.LoanRecord{
					MonthDate:    "March 2021",
					StartBalance: 250000,
					EndBalance:   0,
					Borrowers: []main.Borrower{
						{"1", 100000},
						{"2", 50000},
						{"3", 50000},
						{"4", 50000},
					},
				},
			))
		})
	})
})

var _ = Describe("RecordJSON", func() {
	BeforeEach(BeforeEachHandler)
	When("Record data is normal", func() {
		It("should record data and not return error", func() {
			err := main.RecordJSON(
				main.LoanRecord{
					MonthDate:    "January 2021",
					StartBalance: 500000,
					EndBalance:   50000,
					Borrowers: []main.Borrower{
						{"1", 150000},
						{"2", 150000},
						{"3", 150000},
					},
				},
				PATH_TEST_FILE,
			)
			Expect(err).To(BeNil())
			Expect(AssertRecord()).To(Equal(`{"month_date":"January 2021","start_balance":500000,"end_balance":50000,"borrowers":[{"id":"1","total_loan":150000},{"id":"2","total_loan":150000},{"id":"3","total_loan":150000}]}`))
		})
	})

	When("Record data is empty", func() {
		It("should record data and not return error", func() {
			err := main.RecordJSON(
				main.LoanRecord{},
				PATH_TEST_FILE,
			)
			Expect(err).To(BeNil())
			Expect(AssertRecord()).To(Equal(`{"month_date":"","start_balance":0,"end_balance":0,"borrowers":null}`))
		})
	})

	AfterEach(AfterEachHandler)
})
