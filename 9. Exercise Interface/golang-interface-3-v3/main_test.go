package main_test

import (
	main "a21hc3NpZ25tZW50"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("ChangeToStandartTime", func() {
	Context("input data is type of string", func() {
		When("input data is '12:00'", func() {
			It("should return '12:00 PM'", func() {
				Expect(main.ChangeToStandartTime("12:00")).To(Equal("12:00 PM"))
			})
		})

		When("input data is '02:00'", func() {
			It("should return '02:00 AM'", func() {
				Expect(main.ChangeToStandartTime("02:00")).To(Equal("02:00 AM"))
			})
		})
		When("input data is '14:00'", func() {
			It("should return '02:00 PM'", func() {
				Expect(main.ChangeToStandartTime("14:00")).To(Equal("02:00 PM"))
			})
		})

		When("input data is '23:11'", func() {
			It("should return '11:11 PM'", func() {
				Expect(main.ChangeToStandartTime("23:11")).To(Equal("11:11 PM"))
			})
		})

		When("input data is '23:'", func() {
			It("should return 'Invalid input'", func() {
				Expect(main.ChangeToStandartTime("23:")).To(Equal("Invalid input"))
			})
		})

		When("input data is '2300'", func() {
			It("should return 'Invalid input'", func() {
				Expect(main.ChangeToStandartTime("2300")).To(Equal("Invalid input"))
			})
		})
	})

	Context("input data is type of slice int", func() {
		When("input data is [12, 00]", func() {
			It("should return '12:00 PM'", func() {
				Expect(main.ChangeToStandartTime([]int{12, 00})).To(Equal("12:00 PM"))
			})
		})

		When("input data is [2, 0]", func() {
			It("should return '02:00 AM'", func() {
				Expect(main.ChangeToStandartTime([]int{02, 0})).To(Equal("02:00 AM"))
			})
		})
		When("input data is [14, 0]", func() {
			It("should return '02:00 PM'", func() {
				Expect(main.ChangeToStandartTime([]int{14, 0})).To(Equal("02:00 PM"))
			})
		})

		When("input data is [23, 11]", func() {
			It("should return '11:11 PM'", func() {
				Expect(main.ChangeToStandartTime([]int{23, 11})).To(Equal("11:11 PM"))
			})
		})

		When("input data is [23]", func() {
			It("should return 'Invalid input'", func() {
				Expect(main.ChangeToStandartTime([]int{23})).To(Equal("Invalid input"))
			})
		})
	})

	Context("input data is type of map[string]int", func() {
		When("input data is {'hour': 12, 'minute': 00}", func() {
			It("should return '12:00 PM'", func() {
				Expect(main.ChangeToStandartTime(map[string]int{"hour": 12, "minute": 00})).To(Equal("12:00 PM"))
			})
		})

		When("input data is {'hour': 2, 'minute': 0}", func() {
			It("should return '02:00 AM'", func() {
				Expect(main.ChangeToStandartTime(map[string]int{"hour": 02, "minute": 0})).To(Equal("02:00 AM"))
			})
		})
		When("input data is {'hour': 14, 'minute': 0}", func() {
			It("should return '02:00 PM'", func() {
				Expect(main.ChangeToStandartTime(map[string]int{"hour": 14, "minute": 0})).To(Equal("02:00 PM"))
			})
		})

		When("input data is {'hour': 23, 'minute': 11}", func() {
			It("should return '11:11 PM'", func() {
				Expect(main.ChangeToStandartTime(map[string]int{"hour": 23, "minute": 11})).To(Equal("11:11 PM"))
			})
		})

		When("input data is {'hour': 23}", func() {
			It("should return 'Invalid input'", func() {
				Expect(main.ChangeToStandartTime(map[string]int{"hour": 23})).To(Equal("Invalid input"))
			})
		})

		When("input data is {'minute': 11, 'second': 11}", func() {
			It("should return 'Invalid input'", func() {
				Expect(main.ChangeToStandartTime(map[string]int{"minute": 11, "second": 11})).To(Equal("Invalid input"))
			})
		})
	})

	Context("input data is type of struct Time", func() {
		When("input data is Time{12, 00}", func() {
			It("should return '12:00 PM'", func() {
				Expect(main.ChangeToStandartTime(main.Time{12, 00})).To(Equal("12:00 PM"))
			})
		})

		When("input data is Time{2, 0}", func() {
			It("should return '02:00 AM'", func() {
				Expect(main.ChangeToStandartTime(main.Time{2, 0})).To(Equal("02:00 AM"))
			})
		})
		When("input data is Time{14, 0}", func() {
			It("should return '02:00 PM'", func() {
				Expect(main.ChangeToStandartTime(main.Time{14, 0})).To(Equal("02:00 PM"))
			})
		})

		When("input data is Time{23, 11}", func() {
			It("should return '11:11 PM'", func() {
				Expect(main.ChangeToStandartTime(main.Time{23, 11})).To(Equal("11:11 PM"))
			})
		})

		When("input data is Time{20, 11}", func() {
			It("should return '08:11 PM'", func() {
				Expect(main.ChangeToStandartTime(main.Time{20, 11})).To(Equal("08:11 PM"))
			})
		})

		When("input data is Time{4, 11}", func() {
			It("should return '04:11 AM'", func() {
				Expect(main.ChangeToStandartTime(main.Time{4, 11})).To(Equal("04:11 AM"))
			})
		})
	})
})
