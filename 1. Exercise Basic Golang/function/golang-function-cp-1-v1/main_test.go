package main_test

import (
	main "a21hc3NpZ25tZW50"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("DateFormat", func() {
	When("input day is 1, month is 1, year is 2012", func() {
		It("should return 01-January-2012", func() {
			Expect(main.DateFormat(1, 1, 2012)).To(Equal("01-January-2012"))
		})
	})

	When("input day is 2, month is 2, year is 2012", func() {
		It("should return 02-February-2012", func() {
			Expect(main.DateFormat(2, 2, 2012)).To(Equal("02-February-2012"))
		})
	})

	When("input day is 3, month is 3, year is 2012", func() {
		It("should return 03-March-2012", func() {
			Expect(main.DateFormat(3, 3, 2012)).To(Equal("03-March-2012"))
		})
	})

	When("input day is 4, month is 4, year is 2012", func() {
		It("should return 04-April-2012", func() {
			Expect(main.DateFormat(4, 4, 2012)).To(Equal("04-April-2012"))
		})
	})

	When("input day is 5, month is 5, year is 2012", func() {
		It("should return 05-May-2012", func() {
			Expect(main.DateFormat(5, 5, 2012)).To(Equal("05-May-2012"))
		})
	})

	When("input day is 6, month is 6, year is 2012", func() {
		It("should return 06-June-2012", func() {
			Expect(main.DateFormat(6, 6, 2012)).To(Equal("06-June-2012"))
		})
	})

	When("input day is 7, month is 7, year is 2012", func() {
		It("should return 07-July-2012", func() {
			Expect(main.DateFormat(7, 7, 2012)).To(Equal("07-July-2012"))
		})
	})

	When("input day is 8, month is 8, year is 2012", func() {
		It("should return 08-August-2012", func() {
			Expect(main.DateFormat(8, 8, 2012)).To(Equal("08-August-2012"))
		})
	})

	When("input day is 9, month is 9, year is 2012", func() {
		It("should return 09-September-2012", func() {
			Expect(main.DateFormat(9, 9, 2012)).To(Equal("09-September-2012"))
		})
	})

	When("input day is 10, month is 10, year is 2012", func() {
		It("should return 10-October-2012", func() {
			Expect(main.DateFormat(10, 10, 2012)).To(Equal("10-October-2012"))
		})
	})

	When("input day is 11, month is 11, year is 2012", func() {
		It("should return 11-November-2012", func() {
			Expect(main.DateFormat(11, 11, 2012)).To(Equal("11-November-2012"))
		})
	})

	When("input day is 12, month is 12, year is 2012", func() {
		It("should return 12-December-2012", func() {
			Expect(main.DateFormat(12, 12, 2012)).To(Equal("12-December-2012"))
		})
	})
})
