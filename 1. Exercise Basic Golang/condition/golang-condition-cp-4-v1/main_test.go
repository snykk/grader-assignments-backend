package main_test

import (
	main "a21hc3NpZ25tZW50"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("GetTicketPrice", func() {
	When("total price ticket less than $100", func() {
		It("should return total price without discount", func() {
			Expect(main.GetTicketPrice(1, 1, 1, 20)).To(Equal(float32(60.0)))
			Expect(main.GetTicketPrice(0, 0, 5, 1)).To(Equal(float32(50.0)))
			Expect(main.GetTicketPrice(1, 0, 6, 15)).To(Equal(float32(90.0)))
		})
	})

	When("total price ticket more than equal $100", func() {
		Context("when day is odd", func() {
			It("should return total price with 15% discount if total ticket less than 5", func() {
				Expect(main.GetTicketPrice(4, 0, 0, 13)).To(Equal(float32(102.0)))
			})

			It("should return total price with 25% discount if total ticket more than equal 5", func() {
				Expect(main.GetTicketPrice(3, 3, 3, 20)).To(Equal(float32(144.0)))
				Expect(main.GetTicketPrice(4, 4, 4, 20)).To(Equal(float32(192.0)))
			})
		})

		Context("when day is even", func() {
			It("should return total price with 10% discount if total ticket less than 5", func() {
				Expect(main.GetTicketPrice(4, 0, 0, 20)).To(Equal(float32(108.0)))
			})

			It("should return total price with 20% discount if total ticket more than equal 5", func() {
				Expect(main.GetTicketPrice(3, 3, 3, 22)).To(Equal(float32(144.0)))
				Expect(main.GetTicketPrice(4, 4, 4, 22)).To(Equal(float32(192.0)))
			})
		})
	})
})
