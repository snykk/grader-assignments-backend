package main_test

import (
	main "a21hc3NpZ25tZW50"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("GraduateStudent", func() {
	When("score is more than equal 70 and absent is less than 5", func() {
		It("should return 'lulus'", func() {
			Expect(main.GraduateStudent(70, 4)).To(Equal("lulus"))
			Expect(main.GraduateStudent(80, 1)).To(Equal("lulus"))
			Expect(main.GraduateStudent(70, 1)).To(Equal("lulus"))
		})
	})

	When("score is more than 90 and absent is 0", func() {
		It("should return 'lulus'", func() {
			Expect(main.GraduateStudent(90, 0)).To(Equal("lulus"))
			Expect(main.GraduateStudent(95, 0)).To(Equal("lulus"))
			Expect(main.GraduateStudent(100, 0)).To(Equal("lulus"))
		})
	})

	When("score is less than 70", func() {
		It("should return 'tidak lulus'", func() {
			Expect(main.GraduateStudent(69, 4)).To(Equal("tidak lulus"))
			Expect(main.GraduateStudent(40, 1)).To(Equal("tidak lulus"))
			Expect(main.GraduateStudent(69, 6)).To(Equal("tidak lulus"))
		})
	})

	When("absent is more than equal 5", func() {
		It("should return 'tidak lulus'", func() {
			Expect(main.GraduateStudent(70, 10)).To(Equal("tidak lulus"))
			Expect(main.GraduateStudent(100, 5)).To(Equal("tidak lulus"))
			Expect(main.GraduateStudent(80, 6)).To(Equal("tidak lulus"))
		})
	})

	When("score is 0 and absent is 0", func() {
		It("should return 'tidak lulus'", func() {
			Expect(main.GraduateStudent(0, 0)).To(Equal("tidak lulus"))
		})
	})

	When("score is 0 and absent is more than 5", func() {
		It("should return 'tidak lulus'", func() {
			Expect(main.GraduateStudent(0, 6)).To(Equal("tidak lulus"))
			Expect(main.GraduateStudent(0, 7)).To(Equal("tidak lulus"))
			Expect(main.GraduateStudent(0, 8)).To(Equal("tidak lulus"))
		})
	})
})
