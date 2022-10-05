package main_test

import (
	main "a21hc3NpZ25tZW50"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("MoneyChange", func() {
	When("input money 100000 and list price 50000, 10000, 10000, 5000, 5000", func() {
		It("should return 'Rp. 20.000'", func() {
			Expect(main.MoneyChange(100000, 50000, 10000, 10000, 5000, 5000)).To(Equal("Rp. 20.000"))
		})
	})

	When("input money 150000 and list price 50000, 50000, 50000, 30000", func() {
		It("should return 'Uang tidak cukup'", func() {
			Expect(main.MoneyChange(150000, 50000, 50000, 50000, 30000)).To(Equal("Uang tidak cukup"))
		})
	})

	When("input money 10000 and list price 5000, 1000", func() {
		It("should return 'Uang tidak cukup'", func() {
			Expect(main.MoneyChange(10000, 5000, 1000)).To(Equal("Rp. 4.000"))
		})
	})

	When("input money 10000 and list price 5000, 5000", func() {
		It("should return 'Uang tidak cukup'", func() {
			Expect(main.MoneyChange(10000, 5000, 5000)).To(Equal("Rp. 0"))
		})
	})

	When("input money 10000000 and list price 5000, 1000", func() {
		It("should return 'Uang tidak cukup'", func() {
			Expect(main.MoneyChange(10000000, 5000, 1000)).To(Equal("Rp. 9.994.000"))
		})
	})
})

var _ = Describe("ChangeToCurrency", func() {
	When("input change 100", func() {
		It("should return 'Rp. 100'", func() {
			Expect(main.ChangeToCurrency(100)).To(Equal("Rp. 100"))
		})
	})
	When("input change 1000", func() {
		It("should return 'Rp. 1.000'", func() {
			Expect(main.ChangeToCurrency(1000)).To(Equal("Rp. 1.000"))
		})
	})
	When("input change 10000", func() {
		It("should return 'Rp. 10.000'", func() {
			Expect(main.ChangeToCurrency(10000)).To(Equal("Rp. 10.000"))
		})
	})

	When("input change 100000", func() {
		It("should return 'Rp. 100.000'", func() {
			Expect(main.ChangeToCurrency(100000)).To(Equal("Rp. 100.000"))
		})
	})

	When("input change 1000000", func() {
		It("should return 'Rp. 1.000.000'", func() {
			Expect(main.ChangeToCurrency(1000000)).To(Equal("Rp. 1.000.000"))
		})
	})

	When("input change 10000000", func() {
		It("should return 'Rp. 10.000.000'", func() {
			Expect(main.ChangeToCurrency(10000000)).To(Equal("Rp. 10.000.000"))
		})
	})

	When("input change 100000000", func() {
		It("should return 'Rp. 100.000.000'", func() {
			Expect(main.ChangeToCurrency(100000000)).To(Equal("Rp. 100.000.000"))
		})
	})
})
