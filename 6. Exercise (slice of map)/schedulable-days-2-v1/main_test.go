package main_test

import (
	main "a21hc3NpZ25tZW50"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type TestData struct {
	Input    [][]int
	Expected []int
}

var _ = Describe("SchedulableDays", func() {
	When("there is no villager", func() {
		It("should return empty array", func() {
			actual := main.SchedulableDays([][]int{})

			Expect(actual).To(Equal([]int{}))
			Expect(actual).To(BeEmpty())
		})
	})

	When("data is containing only one villager", func() {
		It("should return date of the villager", func() {
			var testData = []TestData{
				{[][]int{{1, 2, 3}}, []int{1, 2, 3}},
				{[][]int{{1, 2, 3, 4, 5, 6, 7, 11, 12, 13, 14, 15, 16, 17, 19, 20, 21, 22, 23, 34, 25, 26}}, []int{1, 2, 3, 4, 5, 6, 7, 11, 12, 13, 14, 15, 16, 17, 19, 20, 21, 22, 23, 34, 25, 26}},
			}

			for _, data := range testData {
				actual := main.SchedulableDays(data.Input)

				Expect(actual).NotTo(BeEmpty())

				for _, date := range actual {
					Expect(data.Expected).To(ContainElement(date))
				}
			}
		})
	})

	When("data is containing many villager with same date", func() {
		It("should return all date in parameter", func() {
			var testData = []TestData{
				{[][]int{{1, 2, 3, 4, 5, 6, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30}, {1, 2, 3, 4, 5, 6, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30}, {1, 2, 3, 4, 5, 6, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30}}, []int{1, 2, 3, 4, 5, 6, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30}},
				{[][]int{
					{10, 11, 12, 13, 14, 15, 16, 20, 21, 22, 23, 24, 25, 26, 27, 28},
					{10, 11, 12, 13, 14, 15, 16, 20, 21, 22, 23, 24, 25, 26, 27, 28},
					{10, 11, 12, 13, 14, 15, 16, 20, 21, 22, 23, 24, 25, 26, 27, 28},
					{10, 11, 12, 13, 14, 15, 16, 20, 21, 22, 23, 24, 25, 26, 27, 28},
					{10, 11, 12, 13, 14, 15, 16, 20, 21, 22, 23, 24, 25, 26, 27, 28},
				}, []int{10, 11, 12, 13, 14, 15, 16, 20, 21, 22, 23, 24, 25, 26, 27, 28}},
			}

			for _, data := range testData {
				actual := main.SchedulableDays(data.Input)

				Expect(actual).NotTo(BeEmpty())

				for _, date := range data.Expected {
					Expect(actual).To(ContainElement(date))
				}
			}
		})
	})

	When("data is containing many villager with different date", func() {
		It("should return all same date", func() {
			var testData = []TestData{
				{[][]int{
					{1, 2, 3, 4, 5, 6, 7, 11, 12, 13, 14, 15, 16, 17, 19, 20},
					{3, 4, 5, 6, 7, 11, 12, 13, 14, 15, 16, 17, 19, 20},
					{1, 2, 3, 4, 5, 6, 7, 11, 12, 13, 14, 15, 16, 17, 19, 20},
					{15, 16, 17, 18, 19, 20},
					{1, 2, 3, 4, 5, 6, 7, 11, 12, 13, 14, 15, 16, 17, 19, 20},
					{1, 2, 3, 4, 5, 6, 7, 11, 12, 13, 14, 15, 16, 17, 19, 20},
				}, []int{15, 16, 17, 19, 20}},
				{[][]int{
					{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31},
					{2, 5, 6, 7, 8, 9, 10, 18, 19, 27, 28, 29, 30, 31},
					{11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31},
					{27, 28, 29, 30, 31},
				}, []int{27, 28, 29, 30, 31}},
			}

			for _, data := range testData {
				actual := main.SchedulableDays(data.Input)

				Expect(actual).NotTo(BeEmpty())

				for _, date := range data.Expected {
					Expect(actual).To(ContainElement(date))
				}
			}
		})
	})
})
