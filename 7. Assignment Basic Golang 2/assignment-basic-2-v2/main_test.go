package main_test

import (
	main "a21hc3NpZ25tZW50"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("GetDayDifference", func() {
	When("date is '21 February - 21 February 2021'", func() {
		It("should return 1", func() {
			Expect(main.GetDayDifference("21 February - 21 February 2021")).To(Equal(1))
		})
	})

	When("date is '21 February - 23 February 2021'", func() {
		It("should return 3", func() {
			Expect(main.GetDayDifference("21 February - 23 February 2021")).To(Equal(3))
		})
	})

	When("date is '25 January - 5 February 2021'", func() {
		It("should return 11", func() {
			Expect(main.GetDayDifference("25 January - 5 February 2021")).To(Equal(11))
		})
	})

	When("date is '30 March - 2 May 2021'", func() {
		It("should return 33", func() {
			Expect(main.GetDayDifference("30 March - 2 May 2021")).To(Equal(33))
		})
	})

	When("year leap with date is '25 February - 5 March 2020'", func() {
		It("should return 10", func() {
			Expect(main.GetDayDifference("25 February - 5 March 2020")).To(Equal(10))
		})
	})

	When("year is not leap with date is '25 February - 5 March 2021'", func() {
		It("should return 9", func() {
			Expect(main.GetDayDifference("25 February - 5 March 2021")).To(Equal(9))
		})
	})
})

var _ = Describe("FormatRupiah", func() {
	When("number is 100", func() {
		It("should return Rp 100", func() {
			Expect(main.FormatRupiah(100)).To(Equal("Rp 100"))
		})
	})

	When("number is 1000", func() {
		It("should return Rp 1.000", func() {
			Expect(main.FormatRupiah(1000)).To(Equal("Rp 1.000"))
		})
	})
	When("number is 1000", func() {
		It("should return Rp 10.000", func() {
			Expect(main.FormatRupiah(10000)).To(Equal("Rp 10.000"))
		})
	})
	When("number is 1000000", func() {
		It("should return Rp 1.000.000", func() {
			Expect(main.FormatRupiah(1000000)).To(Equal("Rp 1.000.000"))
		})
	})

	When("number is 10000000", func() {
		It("should return Rp 10.000.000", func() {
			Expect(main.FormatRupiah(10000000)).To(Equal("Rp 10.000.000"))
		})
	})
})

var _ = Describe("GetSalary", func() {
	When("range is 1 and data is [['A', 'B']]", func() {
		It("should return map[A:Rp 50.000 B:Rp 50.000]", func() {
			Expect(main.GetSalary(1, [][]string{{"A", "B"}})).To(Equal(map[string]string{"A": "Rp 50.000", "B": "Rp 50.000"}))
		})
	})

	When("range is 2 and data is [['A', 'B'], ['A', 'C']]", func() {
		It("should return map[A:Rp 100.000 B:Rp 50.000 C:Rp 50.000]", func() {
			Expect(main.GetSalary(2, [][]string{{"A", "B"}, {"A", "C"}})).To(Equal(map[string]string{"A": "Rp 100.000", "B": "Rp 50.000", "C": "Rp 50.000"}))
		})
	})

	When("range is 1 and data is [['A', 'B'], ['A', 'C'], ['A', 'D']]", func() {
		It("should return map[A:Rp 50.000 B:Rp 50.000]", func() {
			Expect(main.GetSalary(1, [][]string{{"A", "B"}, {"A", "C"}, {"A", "D"}})).To(Equal(map[string]string{"A": "Rp 50.000", "B": "Rp 50.000"}))
		})
	})
})

var _ = Describe("GetSalaryOverview", func() {
	When("dateRange is '21 February - 21 February 2021' and data is [['A', 'B']]", func() {
		It("should return map[A:Rp 50.000 B:Rp 50.000]", func() {
			Expect(main.GetSalaryOverview("21 February - 21 February 2021", [][]string{{"A", "B"}})).To(Equal(map[string]string{"A": "Rp 50.000", "B": "Rp 50.000"}))
		})
	})

	When("range is '21 February - 22 February 2021' and data is [['A', 'B'], ['A', 'C']]", func() {
		It("should return map[A:Rp 100.000 B:Rp 50.000 C:Rp 50.000]", func() {
			Expect(main.GetSalaryOverview("21 February - 22 February 2021", [][]string{{"A", "B"}, {"A", "C"}})).To(Equal(map[string]string{"A": "Rp 100.000", "B": "Rp 50.000", "C": "Rp 50.000"}))
		})
	})

	When("range is '21 February - 21 February 2021' and data is [['A', 'B'], ['A', 'C'], ['A', 'D']]", func() {
		It("should return map[A:Rp 50.000 B:Rp 50.000]", func() {
			Expect(main.GetSalaryOverview("21 February - 21 February 2021", [][]string{{"A", "B"}, {"A", "C"}, {"A", "D"}})).To(Equal(map[string]string{"A": "Rp 50.000", "B": "Rp 50.000"}))
		})
	})
})
