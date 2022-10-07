package main_test

import (
	main "a21hc3NpZ25tZW50"
	"os"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

const PATH_TEST_FILE = "transaction_test.txt"

var BeforeEachHandler = func() {
	_, err := os.Stat(PATH_TEST_FILE)

	if os.IsNotExist(err) {
		var file, err = os.Create(PATH_TEST_FILE)
		if err != nil {
			panic(err)
		}
		defer file.Close()
	}
}

var AfterEachHandler = func() {
	os.Remove(PATH_TEST_FILE)
}

// ReadFile is function to read file from transaction_test.txt for testing
var ReadFile = func() []string {
	var data, err = os.ReadFile(PATH_TEST_FILE)
	if err != nil {
		panic(err)
	}

	return strings.Split(string(data), "\n")
}

var _ = Describe("Main", func() {
	BeforeEach(BeforeEachHandler)

	Describe("RecordTransactions", func() {
		When("transaction data is empty", func() {
			It("should not insert data in transaction file", func() {
				err := main.RecordTransactions(PATH_TEST_FILE, []main.Transaction{})

				Expect(err).To(BeNil())

				actual := ReadFile()

				Expect(actual).To(HaveLen(1))
				Expect(actual[0]).To(Equal(""))
			})
		})

		When("transaction data have only income in the one date", func() {
			It("should have one record", func() {
				err := main.RecordTransactions(PATH_TEST_FILE, []main.Transaction{
					{"01/01/2021", "income", 100000},
					{"01/01/2021", "income", 100000},
					{"01/01/2021", "income", 100000},
					{"01/01/2021", "income", 100000},
				})

				Expect(err).To(BeNil())

				actual := ReadFile()

				Expect(actual).To(HaveLen(1))
				Expect(actual[0]).To(Equal("01/01/2021;income;400000"))
			})
		})

		When("transaction data have only expense in the one date", func() {
			It("should have one record", func() {
				err := main.RecordTransactions(PATH_TEST_FILE, []main.Transaction{
					{"01/01/2021", "expense", 100000},
					{"01/01/2021", "expense", 100000},
					{"01/01/2021", "expense", 100000},
					{"01/01/2021", "expense", 100000},
				})

				Expect(err).To(BeNil())

				actual := ReadFile()

				Expect(actual).To(HaveLen(1))
				Expect(actual[0]).To(Equal("01/01/2021;expense;400000"))
			})
		})

		When("transaction data have only income data", func() {
			It("should have record with all income transaction type", func() {
				err := main.RecordTransactions(PATH_TEST_FILE, []main.Transaction{
					{"01/01/2021", "income", 100000},
					{"01/01/2021", "income", 1000},
					{"02/01/2021", "income", 20000},
					{"02/01/2021", "income", 233000},
					{"03/01/2021", "income", 20000},
					{"03/01/2021", "income", 22000},
					{"04/01/2021", "income", 24000},
					{"04/01/2021", "income", 20000},
					{"05/01/2021", "income", 20000},
					{"05/01/2021", "income", 50000},
					{"06/01/2021", "income", 60000},
					{"06/01/2021", "income", 70000},
					{"07/01/2021", "income", 80000},
					{"07/01/2021", "income", 110000},
					{"07/01/2021", "income", 120000},
					{"07/01/2021", "income", 111200},
				})

				Expected := []string{
					"01/01/2021;income;101000", "02/01/2021;income;253000", "03/01/2021;income;42000", "04/01/2021;income;44000", "05/01/2021;income;70000", "06/01/2021;income;130000", "07/01/2021;income;421200",
				}

				Expect(err).To(BeNil())

				Actual := ReadFile()

				Expect(Actual).To(HaveLen(7))
				Expect(Actual).To(Equal(Expected))
			})
		})

		When("transaction data have only expense data", func() {
			It("should have record with all expense transaction type", func() {
				err := main.RecordTransactions(PATH_TEST_FILE, []main.Transaction{
					{"01/01/2021", "expense", 100000},
					{"01/01/2021", "expense", 1000},
					{"02/01/2021", "expense", 3424},
					{"02/01/2021", "expense", 42000},
					{"03/01/2021", "expense", 22321},
					{"04/01/2021", "expense", 223200},
					{"02/01/2021", "expense", 2300},
					{"05/01/2021", "expense", 2213},
					{"06/01/2021", "expense", 4545},
					{"07/01/2021", "expense", 55500},
					{"08/01/2021", "expense", 200000},
					{"10/01/2021", "expense", 20000},
					{"11/01/2021", "expense", 10000},
					{"12/01/2021", "expense", 55500},
					{"13/01/2021", "expense", 55500},
					{"02/01/2021", "expense", 55500},
					{"02/01/2021", "expense", 10000},
					{"14/01/2021", "expense", 20000},
					{"11/01/2021", "expense", 20000},
					{"15/01/2021", "expense", 10000},
					{"16/01/2021", "expense", 20000},
					{"02/01/2021", "expense", 55500},
					{"17/01/2021", "expense", 10000},
					{"06/01/2021", "expense", 20000},
					{"18/01/2021", "expense", 10000},
					{"03/01/2021", "expense", 20000},
					{"04/01/2021", "expense", 10000},
					{"19/01/2021", "expense", 55500},
					{"20/01/2021", "expense", 55500},
					{"21/01/2021", "expense", 10000},
					{"22/01/2021", "expense", 10000},
					{"23/01/2021", "expense", 10000},
					{"24/01/2021", "expense", 10000},
				})

				Expected := []string{
					"01/01/2021;expense;101000",
					"02/01/2021;expense;168724",
					"03/01/2021;expense;42321",
					"04/01/2021;expense;233200",
					"05/01/2021;expense;2213",
					"06/01/2021;expense;24545",
					"07/01/2021;expense;55500",
					"08/01/2021;expense;200000",
					"10/01/2021;expense;20000",
					"11/01/2021;expense;30000",
					"12/01/2021;expense;55500",
					"13/01/2021;expense;55500",
					"14/01/2021;expense;20000",
					"15/01/2021;expense;10000",
					"16/01/2021;expense;20000",
					"17/01/2021;expense;10000",
					"18/01/2021;expense;10000",
					"19/01/2021;expense;55500",
					"20/01/2021;expense;55500",
					"21/01/2021;expense;10000",
					"22/01/2021;expense;10000",
					"23/01/2021;expense;10000",
					"24/01/2021;expense;10000",
				}

				Expect(err).To(BeNil())

				Actual := ReadFile()

				Expect(Actual).To(HaveLen(23))
				Expect(Actual).To(Equal(Expected))
			})
		})

		When("transaction data have income and expense data", func() {
			It("should have record with all income and expense transaction type", func() {
				err := main.RecordTransactions(PATH_TEST_FILE, []main.Transaction{
					{"01/01/2021", "income", 200000},
					{"01/01/2021", "income", 1000},
					{"01/01/2021", "expense", 100000},
					{"01/01/2021", "expense", 1000},
					{"02/01/2021", "income", 20000},
					{"02/01/2021", "income", 233000},
					{"02/01/2021", "expense", 3424},
					{"02/01/2021", "expense", 2300},
					{"02/01/2021", "expense", 42000},
					{"03/01/2021", "income", 20000},
					{"03/01/2021", "income", 22000},
					{"03/01/2021", "expense", 22321},
					{"04/01/2021", "income", 24000},
					{"04/01/2021", "income", 20000},
					{"04/01/2021", "expense", 223200},
					{"05/01/2021", "income", 20000},
					{"05/01/2021", "income", 50000},
					{"05/01/2021", "expense", 2213},
					{"06/01/2021", "income", 60000},
					{"06/01/2021", "income", 70000},
					{"06/01/2021", "expense", 4545},
					{"07/01/2021", "income", 80000},
					{"07/01/2021", "income", 110000},
					{"07/01/2021", "income", 120000},
					{"07/01/2021", "income", 111200},
					{"07/01/2021", "expense", 55500},
					{"08/01/2021", "expense", 200000},
					{"10/01/2021", "expense", 20000},
					{"11/01/2021", "expense", 10000},
					{"12/01/2021", "expense", 55500},
					{"13/01/2021", "expense", 55500},
					{"02/01/2021", "expense", 55500},
					{"02/01/2021", "expense", 10000},
					{"14/01/2021", "expense", 20000},
					{"11/01/2021", "expense", 20000},
					{"15/01/2021", "expense", 10000},
					{"16/01/2021", "expense", 20000},
					{"02/01/2021", "expense", 55500},
					{"17/01/2021", "expense", 10000},
					{"06/01/2021", "expense", 20000},
					{"18/01/2021", "expense", 10000},
					{"03/01/2021", "expense", 20000},
					{"04/01/2021", "expense", 10000},
					{"19/01/2021", "expense", 55500},
					{"20/01/2021", "expense", 55500},
					{"21/01/2021", "expense", 10000},
					{"22/01/2021", "expense", 10000},
					{"23/01/2021", "expense", 10000},
					{"24/01/2021", "expense", 10000},
				})

				Expected := []string{
					"01/01/2021;income;100000",
					"02/01/2021;income;84276",
					"03/01/2021;expense;321",
					"04/01/2021;expense;189200",
					"05/01/2021;income;67787",
					"06/01/2021;income;105455",
					"07/01/2021;income;365700",
					"08/01/2021;expense;200000",
					"10/01/2021;expense;20000",
					"11/01/2021;expense;30000",
					"12/01/2021;expense;55500",
					"13/01/2021;expense;55500",
					"14/01/2021;expense;20000",
					"15/01/2021;expense;10000",
					"16/01/2021;expense;20000",
					"17/01/2021;expense;10000",
					"18/01/2021;expense;10000",
					"19/01/2021;expense;55500",
					"20/01/2021;expense;55500",
					"21/01/2021;expense;10000",
					"22/01/2021;expense;10000",
					"23/01/2021;expense;10000",
					"24/01/2021;expense;10000",
				}

				Expect(err).To(BeNil())

				Actual := ReadFile()

				Expect(Actual).To(HaveLen(23))
				Expect(Actual).To(Equal(Expected))
			})
		})
	})

	AfterEach(AfterEachHandler)
})
