package main_test

import (
	main "a21hc3NpZ25tZW50"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("FindMin", func() {
	When("input nums is '1, 2, 3, 4, 5, 6, 7, 8, 9, 10'", func() {
		It("should return 1", func() {
			Expect(main.FindMin(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)).To(Equal(1))
		})
	})

	When("input nums is '100, 200, 300'", func() {
		It("should return 100", func() {
			Expect(main.FindMin(100, 200, 300)).To(Equal(100))
		})
	})
})

var _ = Describe("FindMax", func() {
	When("input nums is '1, 2, 3, 4, 5, 6, 7, 8, 9, 10'", func() {
		It("should return 10", func() {
			Expect(main.FindMax(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)).To(Equal(10))
		})
	})

	When("input nums is '100, 200, 300'", func() {
		It("should return 300", func() {
			Expect(main.FindMax(100, 200, 300)).To(Equal(300))
		})
	})
})

var _ = Describe("SumMinMax", func() {
	When("input nums is '1, 2, 3, 4, 5, 6, 7, 8, 9, 10'", func() {
		It("should return 11", func() {
			Expect(main.SumMinMax(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)).To(Equal(11))
		})
	})

	When("input nums is '333, 456, 654, 123, 111, 1000, 1500, 2000, 3000, 1250, 1111'", func() {
		It("should return 3111", func() {
			Expect(main.SumMinMax(333, 456, 654, 123, 111, 1000, 1500, 2000, 3000, 1250, 1111)).To(Equal(3111))
		})
	})
})
