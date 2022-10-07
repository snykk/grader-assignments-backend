package main_test

import (
	main "a21hc3NpZ25tZW50"
	"a21hc3NpZ25tZW50/internal"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("internal", func() {
	Describe("Calculator", internal.CalculatorTestCase)
})

var _ = Describe("AdvanceCalculator", func() {
	Context("data contains add character '+'", func() {
		When("calculate 10 + 2", func() {
			It("should return 12", func() {
				res := main.AdvanceCalculator("10 + 2")
				Expect(res).To(Equal(float32(12)))
			})
		})

		When("calculate 10 + 2 + 3 + 4", func() {
			It("should return 19", func() {
				res := main.AdvanceCalculator("10 + 2 + 3 + 4")
				Expect(res).To(Equal(float32(19)))
			})
		})
	})

	Context("data contains subtract character '-'", func() {
		When("calculate 10 - 2", func() {
			It("should return 8", func() {
				res := main.AdvanceCalculator("10 - 2")
				Expect(res).To(Equal(float32(8)))
			})
		})

		When("calculate 100 - 20 - 30 - 15", func() {
			It("should return 35", func() {
				res := main.AdvanceCalculator("100 - 20 - 30 - 15")
				Expect(res).To(Equal(float32(35)))
			})
		})
	})

	Context("data contains multiply character '*'", func() {
		When("calculate 3 * 4", func() {
			It("should return 12", func() {
				res := main.AdvanceCalculator("3 * 4")
				Expect(res).To(Equal(float32(12)))
			})
		})

		When("calculate 3 * 4 * 2 * 2", func() {
			It("should return 48", func() {
				res := main.AdvanceCalculator("3 * 4 * 2 * 2")
				Expect(res).To(Equal(float32(48)))
			})
		})
	})

	Context("data contains divide character '/'", func() {
		When("calculate 10 / 2", func() {
			It("should return 5", func() {
				res := main.AdvanceCalculator("10 / 2")
				Expect(res).To(Equal(float32(5)))
			})
		})

		When("calculate 100 / 2 / 5", func() {
			It("should return 10", func() {
				res := main.AdvanceCalculator("100 / 2 / 5")
				Expect(res).To(Equal(float32(10)))
			})
		})
	})

	Context("data contains multiple operation", func() {
		When("calculate 3 * 4 / 2 + 10 - 5", func() {
			It("should return 11", func() {
				res := main.AdvanceCalculator("3 * 4 / 2 + 10 - 5")
				Expect(res).To(Equal(float32(11)))
			})
		})

		When("calculate 100 / 2 / 5 + 10 - 5", func() {
			It("should return 15", func() {
				res := main.AdvanceCalculator("100 / 2 / 5 + 10 - 5")
				Expect(res).To(Equal(float32(15)))
			})
		})

		When("calculate 100 / 2 / 5 + 10 - 5 * 2", func() {
			It("should return 30", func() {
				res := main.AdvanceCalculator("100 / 2 / 5 + 10 - 5 * 2")
				Expect(res).To(Equal(float32(30)))
			})
		})
	})

})
