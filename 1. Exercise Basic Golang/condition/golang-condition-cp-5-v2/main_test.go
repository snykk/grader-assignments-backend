package main_test

import (
	main "a21hc3NpZ25tZW50"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("TicketPlayground", func() {
	When("age is less than 5", func() {
		It("should return -1", func() {
			Expect(main.TicketPlayground(85, 4)).To(Equal(-1))
		})
	})

	When("age is more than 12", func() {
		It("should return 100000", func() {
			Expect(main.TicketPlayground(140, 13)).To(Equal(100000))
		})
	})

	When("age more than equal 5 and less than equal 7", func() {
		Context("and height is more than than 120", func() {
			It("should return 15000", func() {
				Expect(main.TicketPlayground(121, 7)).To(Equal(15000))
			})
		})

		Context("and height is more than 135", func() {
			It("should return 25000", func() {
				Expect(main.TicketPlayground(136, 7)).To(Equal(25000))
			})
		})

		Context("and height is more than 150", func() {
			It("should return 40000", func() {
				Expect(main.TicketPlayground(151, 6)).To(Equal(40000))
			})
		})

		Context("and height is more than 160", func() {
			It("should return 60000", func() {
				Expect(main.TicketPlayground(161, 7)).To(Equal(60000))
			})
		})
	})

	When("age more than equal 8 and less than equal 9", func() {
		Context("and height is more than than 120", func() {
			It("should return 25000", func() {
				Expect(main.TicketPlayground(121, 9)).To(Equal(25000))
			})
		})

		Context("and height is more than 135", func() {
			It("should return 25000", func() {
				Expect(main.TicketPlayground(136, 9)).To(Equal(25000))
			})
		})

		Context("and height is more than 150", func() {
			It("should return 40000", func() {
				Expect(main.TicketPlayground(151, 9)).To(Equal(40000))
			})
		})

		Context("and height is more than 160", func() {
			It("should return 60000", func() {
				Expect(main.TicketPlayground(161, 9)).To(Equal(60000))
			})
		})
	})

	When("age more than equal 10 and less than equal 11", func() {
		Context("and height is more than than 120", func() {
			It("should return 40000", func() {
				Expect(main.TicketPlayground(121, 11)).To(Equal(40000))
			})
		})

		Context("and height is more than 135", func() {
			It("should return 40000", func() {
				Expect(main.TicketPlayground(136, 11)).To(Equal(40000))
			})
		})

		Context("and height is more than 150", func() {
			It("should return 40000", func() {
				Expect(main.TicketPlayground(151, 11)).To(Equal(40000))
			})
		})

		Context("and height is more than 160", func() {
			It("should return 60000", func() {
				Expect(main.TicketPlayground(161, 11)).To(Equal(60000))
			})
		})
	})

	When("age is equal 12", func() {
		Context("and height is more than than 120", func() {
			It("should return 60000", func() {
				Expect(main.TicketPlayground(121, 12)).To(Equal(60000))
			})
		})

		Context("and height is more than 135", func() {
			It("should return 60000", func() {
				Expect(main.TicketPlayground(136, 12)).To(Equal(60000))
			})
		})

		Context("and height is more than 150", func() {
			It("should return 60000", func() {
				Expect(main.TicketPlayground(151, 12)).To(Equal(60000))
			})
		})

		Context("and height is more than 160", func() {
			It("should return 60000", func() {
				Expect(main.TicketPlayground(161, 12)).To(Equal(60000))
			})
		})
	})
})
