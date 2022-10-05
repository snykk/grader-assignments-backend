package main_test

import (
	main "a21hc3NpZ25tZW50"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("ReverseData", func() {
	Context("input is array with zeros as its elements", func() {
		It("should reverse with empty data", func() {
			Expect(main.ReverseData([5]int{0, 0, 0, 0, 0})).To(Equal([5]int{0, 0, 0, 0, 0}))
		})
	})
	Context("input contains elements with same number", func() {
		It("should reverse with fixed data", func() {
			Expect(main.ReverseData([5]int{1, 1, 1, 1, 1})).To(Equal([5]int{1, 1, 1, 1, 1}))
			Expect(main.ReverseData([5]int{22, 22, 22, 22, 22})).To(Equal([5]int{22, 22, 22, 22, 22}))
			Expect(main.ReverseData([5]int{555, 555, 555, 555, 555})).To(Equal([5]int{555, 555, 555, 555, 555}))
		})
	})

	Context("input is contains only single number data", func() {
		It("should just reverse order of data ", func() {
			Expect(main.ReverseData([5]int{1, 2, 3, 4, 5})).To(Equal([5]int{5, 4, 3, 2, 1}))
			Expect(main.ReverseData([5]int{5, 6, 7, 8, 9})).To(Equal([5]int{9, 8, 7, 6, 5}))
			Expect(main.ReverseData([5]int{1, 3, 5, 7, 9})).To(Equal([5]int{9, 7, 5, 3, 1}))
		})
	})

	Context("input is contain any random number data", func() {
		It("should reverse order of list and order of data", func() {
			Expect(main.ReverseData([5]int{10, 10, 10, 10, 10})).To(Equal([5]int{1, 1, 1, 1, 1}))
			Expect(main.ReverseData([5]int{123, 456, 11, 1, 2})).To(Equal([5]int{2, 1, 11, 654, 321}))
			Expect(main.ReverseData([5]int{456789, 44332, 2221, 12, 10})).To(Equal([5]int{1, 21, 1222, 23344, 987654}))
			Expect(main.ReverseData([5]int{23456, 789, 123, 456, 5})).To(Equal([5]int{5, 654, 321, 987, 65432}))

		})
	})
})
