package main_test

import (
	main "a21hc3NpZ25tZW50"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("FindShortestName", func() {
	When("input names is 'Hanif Joko Tio Andi Budi Caca Hamdan'", func() {
		It("should return Tio", func() {
			Expect(main.FindShortestName("Hanif Joko Tio Andi Budi Caca Hamdan")).To(Equal("Tio"))
		})
	})

	When("input names is 'Hanif Joko'", func() {
		It("should return Joko", func() {
			Expect(main.FindShortestName("Hanif Joko")).To(Equal("Joko"))
		})
	})

	When("input names is 'Hanif;Joko;Indah;Ari;Intan'", func() {
		It("should return Ari", func() {
			Expect(main.FindShortestName("Hanif;Joko;Indah;Ari;Intan")).To(Equal("Ari"))
		})
	})

	When("input names is 'Ari,Aru,Ara,Andi,Asik'", func() {
		It("should return Ara", func() {
			Expect(main.FindShortestName("Ari,Aru,Ara,Andi,Asik")).To(Equal("Ara"))
		})
	})

	When("input names is 'A,B,C,D,E'", func() {
		It("should return A", func() {
			Expect(main.FindShortestName("A,B,C,D,E")).To(Equal("A"))
		})
	})

	When("input names have same shortest name", func() {
		It("should return shortest name with first of letter name", func() {
			Expect(main.FindShortestName("Budi;Tia;Tio")).To(Equal("Tia"))
		})
	})
})
