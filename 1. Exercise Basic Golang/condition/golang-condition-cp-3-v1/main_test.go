package main_test

import (
	main "a21hc3NpZ25tZW50"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("GetPredicate", func() {
	When("average score is 100", func() {
		It("should return 'Sempurna'", func() {
			Expect(main.GetPredicate(100, 100, 100, 100)).To(Equal("Sempurna"))
		})
	})

	When("average score more than equal 90", func() {
		It("should return 'Sangat Baik'", func() {
			Expect(main.GetPredicate(90, 90, 90, 90)).To(Equal("Sangat Baik"))
			Expect(main.GetPredicate(92, 92, 95, 100)).To(Equal("Sangat Baik"))
		})
	})
	When("average score more than equal 80 and less than 90", func() {
		It("should return 'Baik'", func() {
			Expect(main.GetPredicate(80, 80, 80, 80)).To(Equal("Baik"))
			Expect(main.GetPredicate(82, 85, 95, 80)).To(Equal("Baik"))
		})
	})
	When("average score more than equal 70 and less than 80", func() {
		It("should return 'Cukup'", func() {
			Expect(main.GetPredicate(75, 75, 75, 75)).To(Equal("Cukup"))
			Expect(main.GetPredicate(70, 70, 70, 90)).To(Equal("Cukup"))
		})
	})
	When("average score is more than equal 60 and less than 70", func() {
		It("should return 'Kurang'", func() {
			Expect(main.GetPredicate(60, 65, 65, 65)).To(Equal("Kurang"))
			Expect(main.GetPredicate(60, 60, 60, 75)).To(Equal("Kurang"))
		})
	})
	When("average score is less than 60", func() {
		It("should return 'Sangat kurang'", func() {
			Expect(main.GetPredicate(50, 50, 50, 50)).To(Equal("Sangat kurang"))
			Expect(main.GetPredicate(50, 50, 50, 60)).To(Equal("Sangat kurang"))
		})
	})
})
