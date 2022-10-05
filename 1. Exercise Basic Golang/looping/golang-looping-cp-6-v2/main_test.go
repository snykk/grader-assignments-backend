package main_test

import (
	main "a21hc3NpZ25tZW50"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("BiggestPairNumber", func() {
	When("input numbers 11223344", func() {
		It("should return 44", func() {
			Expect(main.BiggestPairNumber(11223344)).To(Equal(44))
		})
	})

	When("input numbers 89083278", func() {
		It("should return 89", func() {
			Expect(main.BiggestPairNumber(89083278)).To(Equal(89))
		})
	})

	When("input numbers 111111111", func() {
		It("should return 11", func() {
			Expect(main.BiggestPairNumber(111111111)).To(Equal(11))
		})
	})

	When("input numbers 4994444333", func() {
		It("should return 99", func() {
			Expect(main.BiggestPairNumber(4994444333)).To(Equal(99))
		})
	})

	When("input numbers 111112211", func() {
		It("should return 22", func() {
			Expect(main.BiggestPairNumber(111112211)).To(Equal(22))
		})
	})

	When("input numbers 1193824932472397439", func() {
		It("should return 97", func() {
			Expect(main.BiggestPairNumber(1193824932472397439)).To(Equal(97))
		})
	})
})
