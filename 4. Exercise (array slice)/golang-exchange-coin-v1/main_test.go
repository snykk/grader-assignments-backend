package main_test

import (
	main "a21hc3NpZ25tZW50"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type TestCase struct {
	Input    int
	Expected []int
}

var _ = Describe("ExhangeCoin", func() {
	When("input coin is 0", func() {
		It("return empty array", func() {
			actual := main.ExchangeCoin(0)

			Expect(actual).To(Equal([]int{}))
		})

	})

	When("input match one of the available fractions", func() {
		It("return single fraction of bill", func() {
			testData := []TestCase{
				{1, []int{1}},
				{5, []int{5}},
				{10, []int{10}},
				{20, []int{20}},
				{50, []int{50}},
				{100, []int{100}},
				{200, []int{200}},
				{500, []int{500}},
				{1000, []int{1000}},
			}

			for i := 0; i < len(testData); i++ {
				t := testData[i]

				actual := main.ExchangeCoin(t.Input)
				Expect(actual).To(Equal(t.Expected))
			}
		})
	})

	When("Input coin is random coin", func() {
		It("return correct exchange coin", func() {
			testData := []TestCase{
				{1752, []int{1000, 500, 200, 50, 1, 1}},
				{5000, []int{1000, 1000, 1000, 1000, 1000}},
				{1234, []int{1000, 200, 20, 10, 1, 1, 1, 1}},
			}

			for i := 0; i < len(testData); i++ {
				t := testData[i]

				actual := main.ExchangeCoin(t.Input)
				Expect(actual).To(Equal(t.Expected))
			}
		})
	})
})
