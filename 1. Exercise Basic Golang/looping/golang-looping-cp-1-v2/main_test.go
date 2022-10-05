package main_test

import (
	main "a21hc3NpZ25tZW50"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("CountingNumber", func() {
	When("n is 5", func() {
		It("should return 27", func() {
			Expect(main.CountingNumber(5)).To(Equal(27.0))
		})
	})

	When("n is 10", func() {
		It("should return 104.5", func() {
			Expect(main.CountingNumber(10)).To(Equal(104.5))
		})
	})

	When("n is 50", func() {
		It("should return 2524.5", func() {
			Expect(main.CountingNumber(50)).To(Equal(2524.5))
		})
	})

	When("n is 100", func() {
		It("should return 10049.5", func() {
			Expect(main.CountingNumber(100)).To(Equal(10049.5))
		})
	})

	When("n is 10000", func() {
		It("should return 1.000049995e+08", func() {
			Expect(main.CountingNumber(10000)).To(Equal(1.000049995e+08))
		})
	})
})
