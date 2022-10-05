package main_test

import (
	main "a21hc3NpZ25tZW50"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("BMICalculator", func() {
	Context("gender is 'laki-laki'", func() {
		When("height is 165", func() {
			It("should return 58.5", func() {
				Expect(main.BMICalculator("laki-laki", 165)).To(Equal(58.5))
			})
		})
		When("height is 170", func() {
			It("should return 63", func() {
				Expect(main.BMICalculator("laki-laki", 170)).To(Equal(63.00))
			})
		})
		When("height is 160", func() {
			It("should return 54.00", func() {
				Expect(main.BMICalculator("laki-laki", 160)).To(Equal(54.00))
			})
		})
		When("height is 155", func() {
			It("should return 49.5", func() {
				Expect(main.BMICalculator("laki-laki", 155)).To(Equal(49.5))
			})
		})
	})
	Context("gender is 'perempuan'", func() {
		When("height is 165", func() {
			It("should return 55.25", func() {
				Expect(main.BMICalculator("perempuan", 165)).To(Equal(55.25))
			})
		})
		When("height is 170", func() {
			It("should return 59.5", func() {
				Expect(main.BMICalculator("perempuan", 170)).To(Equal(59.5))
			})
		})
		When("height is 160", func() {
			It("should return 51", func() {
				Expect(main.BMICalculator("perempuan", 160)).To(Equal(51.00))
			})
		})
		When("height is 155", func() {
			It("should return 46.75", func() {
				Expect(main.BMICalculator("perempuan", 155)).To(Equal(46.75))
			})
		})
	})
})
