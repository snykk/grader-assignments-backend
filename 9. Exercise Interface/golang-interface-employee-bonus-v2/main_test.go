package main_test

import (
	main "a21hc3NpZ25tZW50"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Main", func() {
	Describe("EmployeeBonus", func() {
		When("Input data interface contain Junior", func() {
			It("should return bonus based on junior bonus", func() {
				junior := main.Junior{
					Name:         "Junior",
					BaseSalary:   100000,
					WorkingMonth: 12,
				}
				Expect(main.EmployeeBonus(&junior)).To(Equal(100000.0))
			})

		})

		When("Input data interface contain Senior", func() {
			It("should return bonus based on senior bonus", func() {
				senior := main.Senior{
					Name:            "Senior",
					BaseSalary:      100000,
					WorkingMonth:    12,
					PerformanceRate: 0.5,
				}
				Expect(main.EmployeeBonus(&senior)).To(Equal(250000.0))
			})
		})

		When("Input data interface contain Manager", func() {
			It("should return bonus based on manager bonus", func() {
				manager := main.Manager{
					Name:             "Manager",
					BaseSalary:       100000,
					WorkingMonth:     12,
					PerformanceRate:  0.5,
					BonusManagerRate: 0.5,
				}
				Expect(main.EmployeeBonus(&manager)).To(Equal(300000.0))
			})

		})

	})

	Describe("TotalEmployeeBonus", func() {
		When("input data is empty slice", func() {
			It("should return 0", func() {
				Expect(main.TotalEmployeeBonus([]main.Employee{})).To(Equal(0.0))
			})
		})

		When("Input employees containing only Junior", func() {
			It("should return total bonus based on junior bonus", func() {
				juniors := []main.Employee{
					&main.Junior{
						Name:         "Junior",
						BaseSalary:   100000,
						WorkingMonth: 12,
					},
					&main.Junior{
						Name:         "Junior 2",
						BaseSalary:   120000,
						WorkingMonth: 4,
					},
					&main.Junior{
						Name:         "Junior 3",
						BaseSalary:   150000,
						WorkingMonth: 10,
					},
					&main.Junior{
						Name:         "Junior 4",
						BaseSalary:   120000,
						WorkingMonth: 10,
					},
					&main.Junior{
						Name:         "Junior 5",
						BaseSalary:   130000,
						WorkingMonth: 6,
					},
				}
				Expect(main.TotalEmployeeBonus(juniors)).To(Equal(430000.0))
			})
		})
		When("Input employees containing only Senior", func() {
			It("should return total bonus based on Senior bonus", func() {
				seniors := []main.Employee{
					&main.Senior{
						Name:            "Senior Engineer",
						BaseSalary:      100000,
						WorkingMonth:    12,
						PerformanceRate: 0.5,
					},
					&main.Senior{
						Name:            "Senior Engineer 2",
						BaseSalary:      120000,
						WorkingMonth:    4,
						PerformanceRate: 0.3,
					},
					&main.Senior{
						Name:            "Senior Engineer 3",
						BaseSalary:      150000,
						WorkingMonth:    10,
						PerformanceRate: 0.2,
					},
					&main.Senior{
						Name:            "Senior Engineer 4",
						BaseSalary:      120000,
						WorkingMonth:    10,
						PerformanceRate: 0.6,
					},
					&main.Senior{
						Name:            "Senior Engineer 5",
						BaseSalary:      130000,
						WorkingMonth:    6,
						PerformanceRate: 0.45,
					},
				}

				Expect(main.TotalEmployeeBonus(seniors)).To(Equal(1106500.0))
			})
		})
		When("Input employees containing only Manager", func() {
			It("should return total bonus based on Manager bonus", func() {
				managers := []main.Employee{
					&main.Manager{
						Name:             "Manager",
						BaseSalary:       100000,
						WorkingMonth:     12,
						PerformanceRate:  0.5,
						BonusManagerRate: 0.5,
					},
					&main.Manager{
						Name:             "Manager 2",
						BaseSalary:       120000,
						WorkingMonth:     4,
						PerformanceRate:  0.3,
						BonusManagerRate: 0.2,
					},
					&main.Manager{
						Name:             "Manager 3",
						BaseSalary:       150000,
						WorkingMonth:     10,
						PerformanceRate:  0.2,
						BonusManagerRate: 0.1,
					},
					&main.Manager{
						Name:             "Manager 4",
						BaseSalary:       120000,
						WorkingMonth:     10,
						PerformanceRate:  0.6,
						BonusManagerRate: 0.3,
					},
					&main.Manager{
						Name:             "Manager 5",
						BaseSalary:       130000,
						WorkingMonth:     6,
						PerformanceRate:  0.45,
						BonusManagerRate: 0.2,
					},
				}

				Expect(main.TotalEmployeeBonus(managers)).To(Equal(1257500.0))
			})
		})

		When("input employees containing all type of employee", func() {
			It("should return total bonus based on all type of employee", func() {
				employees := []main.Employee{
					&main.Junior{
						Name:         "Junior",
						BaseSalary:   100000,
						WorkingMonth: 12,
					},
					&main.Junior{
						Name:         "Junior 2",
						BaseSalary:   120000,
						WorkingMonth: 4,
					},
					&main.Junior{
						Name:         "Junior 3",
						BaseSalary:   150000,
						WorkingMonth: 10,
					},
					&main.Junior{
						Name:         "Junior 4",
						BaseSalary:   120000,
						WorkingMonth: 10,
					},
					&main.Junior{
						Name:         "Junior 5",
						BaseSalary:   130000,
						WorkingMonth: 6,
					},
					&main.Senior{
						Name:            "Senior Engineer",
						BaseSalary:      100000,
						WorkingMonth:    12,
						PerformanceRate: 0.5,
					},
					&main.Senior{
						Name:            "Senior Engineer 2",
						BaseSalary:      120000,
						WorkingMonth:    4,
						PerformanceRate: 0.3,
					},
					&main.Senior{
						Name:            "Senior Engineer 3",
						BaseSalary:      150000,
						WorkingMonth:    10,
						PerformanceRate: 0.2,
					},
					&main.Senior{
						Name:            "Senior Engineer 4",
						BaseSalary:      120000,
						WorkingMonth:    10,
						PerformanceRate: 0.6,
					},
					&main.Senior{
						Name:            "Senior Engineer 5",
						BaseSalary:      130000,
						WorkingMonth:    6,
						PerformanceRate: 0.45,
					},
					&main.Manager{
						Name:             "Manager",
						BaseSalary:       100000,
						WorkingMonth:     12,
						PerformanceRate:  0.5,
						BonusManagerRate: 0.5,
					},
					&main.Manager{
						Name:             "Manager 2",
						BaseSalary:       120000,
						WorkingMonth:     4,
						PerformanceRate:  0.3,
						BonusManagerRate: 0.2,
					},
					&main.Manager{
						Name:             "Manager 3",
						BaseSalary:       150000,
						WorkingMonth:     10,
						PerformanceRate:  0.2,
						BonusManagerRate: 0.1,
					},
					&main.Manager{
						Name:             "Manager 4",
						BaseSalary:       120000,
						WorkingMonth:     10,
						PerformanceRate:  0.6,
						BonusManagerRate: 0.3,
					},
					&main.Manager{
						Name:             "Manager 5",
						BaseSalary:       130000,
						WorkingMonth:     6,
						PerformanceRate:  0.45,
						BonusManagerRate: 0.2,
					},
				}

				Expect(main.TotalEmployeeBonus(employees)).To(Equal(2794000.0))
			})
		})
	})
})
