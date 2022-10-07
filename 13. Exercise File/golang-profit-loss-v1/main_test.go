package main_test

import (
	main "a21hc3NpZ25tZW50"
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var PATH_TEST_FILE = "transaction_test.txt"

var BeforeEachHandler = func() {
	if _, err := os.Stat(PATH_TEST_FILE); os.IsNotExist(err) {
		var f *os.File
		if f, err = os.Create(PATH_TEST_FILE); err != nil {
			panic(err)
		}

		defer f.Close()
	}
}

var InsertDataTest = func(data string) {
	err := os.WriteFile(PATH_TEST_FILE, []byte(data), 0644)
	if err != nil {
		panic(err)
	}
}

var AfterEachHandler = func() {
	os.Remove(PATH_TEST_FILE)
}

var _ = Describe("Main", func() {
	Describe("Readfile", func() {
		BeforeEach(BeforeEachHandler)

		When("data is empty", func() {
			It("should return empty array", func() {
				data, err := main.Readfile(PATH_TEST_FILE)
				Expect(err).To(BeNil())
				Expect(data).To(Equal([]string{}))
			})
		})

		When("data is single record", func() {
			It("should return single data array of string", func() {
				InsertDataTest("2020-01-01;income;1000")

				data, err := main.Readfile(PATH_TEST_FILE)
				Expect(err).To(BeNil())
				Expect(data).To(Equal([]string{
					"2020-01-01;income;1000",
				}))
			})
		})

		When("data is many records", func() {
			It("should return many data array of string", func() {
				InsertDataTest("2021-01-01;income;1000\n2021-01-02;expense;200\n2021-01-03;income;300")

				data, err := main.Readfile(PATH_TEST_FILE)
				Expect(err).To(BeNil())
				Expect(data).To(Equal([]string{"2021-01-01;income;1000", "2021-01-02;expense;200", "2021-01-03;income;300"}))
			})
		})

		AfterEach(AfterEachHandler)
	})

	Describe("CalculateProfitLoss", func() {
		When("data records only contain income transation", func() {
			It("should return record of profit amount", func() {
				datas := []string{
					"01/01/2021;income;1000", "02/01/2021;income;500", "03/01/2021;income;1000", "04/01/2021;income;1000", "05/01/2021;income;1000", "06/01/2021;income;1000", "07/01/2021;income;1000", "08/01/2021;income;1000", "09/01/2021;income;1000", "10/01/2021;income;1000",
				}

				result := main.CalculateProfitLoss(datas)
				Expect(result).To(Equal("10/01/2021;profit;9500"))

			})
		})

		When("data records only contain expense transation", func() {
			It("should return record of loss amount", func() {
				datas := []string{
					"01/01/2021;expense;21300", "02/01/2021;expense;1000", "03/01/2021;expense;1000", "04/01/2021;expense;1000", "05/01/2021;expense;23440", "06/01/2021;expense;1000", "07/01/2021;expense;1000", "08/01/2021;expense;99000", "09/01/2021;expense;55000", "09/01/2021;expense;4300", "10/01/2021;expense;100", "10/01/2021;expense;100", "11/01/2021;expense;100", "12/01/2021;expense;100", "13/01/2021;expense;2000", "14/01/2021;expense;3000", "15/01/2021;expense;4000",
				}

				result := main.CalculateProfitLoss(datas)
				Expect(result).To(Equal("15/01/2021;loss;217440"))

			})
		})

		When("data records contain income and expense transation", func() {
			It("should return record of profit / loss amount", func() {
				datas := []string{
					"01/01/2021;expense;21300", "01/01/2021;profit;210000", "02/01/2021;expense;1000", "03/01/2021;expense;1000", "04/01/2021;expense;1000", "05/01/2021;expense;23440", "06/01/2021;expense;1000", "07/01/2021;expense;1000", "08/01/2021;expense;99000", "09/01/2021;expense;55000", "09/01/2021;expense;4300", "10/01/2021;expense;100", "10/01/2021;expense;100", "11/01/2021;expense;100", "12/01/2021;expense;100", "13/01/2021;expense;2000", "14/01/2021;expense;3000", "15/01/2021;expense;4000", "01/01/2021;profit;21300", "01/01/2021;profit;210000", "02/01/2021;profit;1000", "03/01/2021;profit;1000", "04/01/2021;profit;1000", "05/01/2021;profit;23440", "06/01/2021;profit;1000", "07/01/2021;expense;1000", "08/01/2021;expense;99000", "09/01/2021;profit;55000", "09/01/2021;expense;4300", "10/01/2021;expense;100", "10/01/2021;expense;100", "11/01/2021;expense;100", "12/01/2021;expense;100", "13/01/2021;expense;2000", "14/01/2021;expense;3000", "15/01/2021;profit;4000",
				}

				result := main.CalculateProfitLoss(datas)
				Expect(result).To(Equal("15/01/2021;loss;854880"))

			})
		})
	})
})
