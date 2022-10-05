package main_test

import (
	main "a21hc3NpZ25tZW50"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type TestData struct {
	InputA   []int
	InputB   []int
	Expected []int
}

var _ = Describe("Main", func() {
	Describe("SchedulableDays", func() {
		Context("all parameter date is empty", func() {
			It("should return empty", func() {
				actual := main.SchedulableDays([]int{}, []int{})

				Expect(actual).To(BeEmpty())
				Expect(actual).To(Equal([]int{}))
			})
		})

		Context("parameter date1 or date2 not containing date", func() {
			It("should return empty", func() {
				testData := []TestData{
					{[]int{}, []int{4, 5, 6, 10, 11, 12, 13, 14, 15, 16, 17, 20, 21}, []int{}},
					{[]int{1, 2, 3, 4, 5, 6, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30}, []int{}, []int{}},
				}

				for _, test := range testData {
					actual := main.SchedulableDays(test.InputA, test.InputB)

					Expect(actual).To(BeEmpty())
					Expect(actual).To(Equal([]int{}))
				}

			})
		})

		Context("all parameter date containing same date", func() {
			It("should return all date in parameter", func() {
				testData := []TestData{
					{[]int{1, 2, 3, 4, 5, 6, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30}, []int{1, 2, 3, 4, 5, 6, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30}, []int{1, 2, 3, 4, 5, 6, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30}},
					{[]int{10, 11, 12, 13, 14, 15, 16, 20, 21, 22, 23, 24, 25, 26, 27, 28}, []int{10, 11, 12, 13, 14, 15, 16, 20, 21, 22, 23, 24, 25, 26, 27, 28}, []int{10, 11, 12, 13, 14, 15, 16, 20, 21, 22, 23, 24, 25, 26, 27, 28}},
				}

				for _, test := range testData {
					actual := main.SchedulableDays(test.InputA, test.InputB)

					Expect(actual).To(Equal(test.Expected))
				}
			})
		})

		Context("all parameter date containing some of same date", func() {
			It("should return all same date", func() {
				testData := []TestData{
					{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}, []int{14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29}, []int{14, 15, 16, 17, 18, 19, 20}},
					{[]int{2, 5, 6, 7, 11, 13, 15, 16, 19, 22, 23, 25, 28}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 24, 25, 26, 27, 28, 29, 30}, []int{2, 5, 6, 7, 11, 13, 15, 16, 25, 28}},
				}

				for _, test := range testData {
					actual := main.SchedulableDays(test.InputA, test.InputB)

					Expect(actual).To(Equal(test.Expected))
				}
			})
		})

		Context("all parameter date containing difference date", func() {
			It("should return empty", func() {
				testData := []TestData{
					{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}, []int{21, 22, 23, 24, 25, 26, 27, 28, 29, 30}, []int{}},
					{[]int{2, 3, 4, 5, 6, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29}, []int{1, 7, 8, 9, 10, 30}, []int{}},
				}

				for _, test := range testData {
					actual := main.SchedulableDays(test.InputA, test.InputB)

					Expect(actual).To(Equal(test.Expected))
				}
			})
		})
	})
})
