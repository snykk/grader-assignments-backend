package main_test

import (
	main "a21hc3NpZ25tZW50"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type TestData struct {
	Amount   int
	Products []main.Product
	Expected []int
}

var _ = Describe("MoneyChanges", func() {
	When("amount is 0 and products is empty", func() {
		It("should return empty slice", func() {
			actual := main.MoneyChanges(0, []main.Product{})

			Expect(actual).To(Equal([]int{}))
			Expect(actual).To(BeEmpty())
		})
	})

	When("amount is same as total products price", func() {
		It("should return empty slice", func() {
			testData := []TestData{
				{
					100000, []main.Product{
						{Name: "Product 1", Price: 45000, Tax: 5000},
						{Name: "Product 2", Price: 45000, Tax: 5000},
					}, []int{},
				},
				{
					600000, []main.Product{
						{Name: "Product 1", Price: 45000, Tax: 5000},
						{Name: "Product 2", Price: 45000, Tax: 5000},

						{Name: "Product 3", Price: 45000, Tax: 5000},
						{Name: "Product 4", Price: 45000, Tax: 5000},

						{Name: "Product 5", Price: 45000, Tax: 5000},
						{Name: "Product 6", Price: 45000, Tax: 5000},

						{Name: "Product 7", Price: 45000, Tax: 5000},
						{Name: "Product 8", Price: 45000, Tax: 5000},

						{Name: "Product 9", Price: 45000, Tax: 5000},
						{Name: "Product 10", Price: 45000, Tax: 5000},

						{Name: "Product 11", Price: 45000, Tax: 5000},
						{Name: "Product 12", Price: 45000, Tax: 5000},
					}, []int{},
				},
			}

			for _, test := range testData {
				actual := main.MoneyChanges(test.Amount, test.Products)

				Expect(actual).To(Equal(test.Expected))
				Expect(actual).To(BeEmpty())
			}
		})
	})

	When("amount is more than 1 fraction coin from total products price", func() {
		It("should return one fraction coin", func() {
			testData := []TestData{
				{
					100100, []main.Product{
						{Name: "Product 1", Price: 45000, Tax: 5000},
						{Name: "Product 2", Price: 45000, Tax: 5000},
					}, []int{100},
				},
				{
					600200, []main.Product{
						{Name: "Product 1", Price: 45000, Tax: 5000},
						{Name: "Product 2", Price: 45000, Tax: 5000},

						{Name: "Product 3", Price: 45000, Tax: 5000},
						{Name: "Product 4", Price: 45000, Tax: 5000},

						{Name: "Product 5", Price: 45000, Tax: 5000},
						{Name: "Product 6", Price: 45000, Tax: 5000},

						{Name: "Product 7", Price: 45000, Tax: 5000},
						{Name: "Product 8", Price: 45000, Tax: 5000},

						{Name: "Product 9", Price: 45000, Tax: 5000},
						{Name: "Product 10", Price: 45000, Tax: 5000},

						{Name: "Product 11", Price: 45000, Tax: 5000},
						{Name: "Product 12", Price: 45000, Tax: 5000},
					}, []int{200},
				},
			}

			for _, test := range testData {
				actual := main.MoneyChanges(test.Amount, test.Products)

				Expect(actual).To(Equal(test.Expected))
				Expect(actual).To(HaveLen(1))
			}
		})
	})

	When("amount is more than the total products price", func() {
		It("should return banknotes", func() {
			testData := []TestData{
				{
					30000, []main.Product{
						{Name: "Baju", Price: 10000, Tax: 1000},
						{Name: "Product 2", Price: 15550, Tax: 1555},
					}, []int{1000, 500, 200, 100, 50, 20, 20, 5},
				},
				{
					750000, []main.Product{
						{Name: "Baju", Price: 10000, Tax: 1000},
						{Name: "Celana", Price: 15550, Tax: 1555},
						{Name: "Sepatu", Price: 10000, Tax: 1000},
						{Name: "Jam", Price: 25000, Tax: 2500},
						{Name: "Kemeja", Price: 15400, Tax: 1540},
						{Name: "Topi", Price: 5000, Tax: 5000},
						{Name: "Baju 2", Price: 10000, Tax: 1000},
						{Name: "Jam 2", Price: 25000, Tax: 2500},
						{Name: "Topi 2", Price: 5000, Tax: 5000},
						{Name: "Celana 2", Price: 15550, Tax: 1555},
						{Name: "Sepatu 2", Price: 10000, Tax: 1000},
						{Name: "Kemeja 2", Price: 15400, Tax: 1540},
						{Name: "Baju 3", Price: 10000, Tax: 1000},
						{Name: "Jam 3", Price: 25000, Tax: 2500},
						{Name: "Topi 3", Price: 5000, Tax: 5000},
						{Name: "Celana 3", Price: 15550, Tax: 1555},
						{Name: "Sepatu 3", Price: 10000, Tax: 1000},
						{Name: "Kemeja 3", Price: 15400, Tax: 1540},
						{Name: "Baju 4", Price: 10000, Tax: 1000},
						{Name: "Jam 4", Price: 25000, Tax: 2500},
						{Name: "Topi 4", Price: 5000, Tax: 5000},
						{Name: "Celana 4", Price: 15550, Tax: 1555},
						{Name: "Sepatu 4", Price: 10000, Tax: 1000},
						{Name: "Kemeja 4", Price: 15400, Tax: 1540},
						{Name: "Baju 5", Price: 10000, Tax: 1000},
						{Name: "Jam 5", Price: 25000, Tax: 2500},
						{Name: "Topi 5", Price: 5000, Tax: 5000},
						{Name: "Celana 5", Price: 15550, Tax: 1555},
						{Name: "Sepatu 5", Price: 10000, Tax: 1000},
						{Name: "Kemeja 5", Price: 15400, Tax: 1540},
						{Name: "Baju 6", Price: 10000, Tax: 1000},
						{Name: "Jam 6", Price: 25000, Tax: 2500},
						{Name: "Topi 6", Price: 5000, Tax: 5000},
						{Name: "Celana 6", Price: 15550, Tax: 1555},
						{Name: "Sepatu 6", Price: 10000, Tax: 1000},
						{Name: "Kemeja 6", Price: 15400, Tax: 1540},
					}, []int{1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 500, 200, 20, 10},
				},
			}

			for _, test := range testData {
				actual := main.MoneyChanges(test.Amount, test.Products)

				Expect(actual).To(Equal(test.Expected))
				Expect(actual).To(HaveLen(len(test.Expected)))
			}
		})
	})
})
