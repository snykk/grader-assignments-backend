package main_test

import (
	main "a21hc3NpZ25tZW50"
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Main", func() {
	Describe("AddGrade", func() {
		var school main.School

		BeforeEach(func() {
			school = main.School{
				Name:    "Hogwarts",
				Address: "Hogwarts, RT14 RW03 Gang Serabi",
				Grades:  []int{},
			}
		})

		When("input data is empty", func() {
			It("should not change the grades", func() {
				school.AddGrade()

				Expect(school.Grades).To(Equal([]int{}))
				Expect(school.Grades).To(BeEmpty())
			})
		})

		When("input data inserted", func() {
			It("should add the grades", func() {
				school.AddGrade(100, 90, 80, 70, 60)

				Expect(school.Grades).To(Equal([]int{100, 90, 80, 70, 60}))
				Expect(school.Grades).To(HaveLen(5))
			})
		})
	})

	Describe("Analysis", func() {
		var school main.School

		BeforeEach(func() {
			school = main.School{
				Name:    "Hogwarts",
				Address: "Hogwarts, RT14 RW03 Gang Serabi",
				Grades:  []int{},
			}
		})

		When("input data is empty", func() {
			It("should return 0, 0, 0", func() {
				avg, min, max := main.Analysis(school)

				Expect(avg).To(Equal(0.0))
				Expect(min).To(Equal(0))
				Expect(max).To(Equal(0))
			})
		})

		When("input data is containing single data", func() {
			It("should return same value in every return", func() {
				school.AddGrade(100)
				avg, min, max := main.Analysis(school)

				Expect(avg).To(Equal(100.0))
				Expect(min).To(Equal(100))
				Expect(max).To(Equal(100))
			})
		})

		When("input data is containing multiple data", func() {
			It("should return correct value in every return", func() {
				school.AddGrade(100, 90, 80, 70, 60, 60, 100, 100, 100, 76, 87, 89, 54, 43, 12, 15, 16, 17, 100, 80, 87, 86, 57, 57, 76, 87, 89, 54, 43, 12, 15, 16, 17, 100, 80, 87, 86, 57, 57, 76, 87, 89, 54, 43, 12, 15, 16, 17, 100, 80, 87, 86, 57, 57, 76, 87, 89, 54, 43, 12, 15, 16, 17, 100, 80, 87, 86, 57, 57, 76, 87, 89, 54, 43, 12, 15, 16, 17, 100, 80, 87, 86, 57, 57, 76, 87, 89, 54, 43, 12, 15, 16, 17, 100, 80, 87, 86, 57, 57, 76, 87, 89, 54, 43, 12, 15, 16, 17, 100, 80, 87, 86, 57, 57, 76, 87, 89, 54, 43, 12, 15, 16, 17, 100, 80, 87, 86, 57, 57, 76, 87, 89, 54, 43, 12, 15, 16, 17, 100, 80, 87, 86, 57, 57)
				avg, min, max := main.Analysis(school)

				Expect(fmt.Sprintf("%.2f", avg)).To(Equal("60.03"))
				Expect(min).To(Equal(12))
				Expect(max).To(Equal(100))
			})
		})
	})
})
