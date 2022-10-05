package main_test

import (
	main "a21hc3NpZ25tZW50"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("EmailInfo", func() {
	When("email is 'admin@yahoo.com'", func() {
		It("should return 'Domain: yahoo dan TLD: com'", func() {
			Expect(main.EmailInfo("admin@yahoo.com")).To(Equal("Domain: yahoo dan TLD: com"))
		})
	})

	When("email is 'admin@google.org'", func() {
		It("should return 'Domain: google dan TLD: org'", func() {
			Expect(main.EmailInfo("admin@google.org")).To(Equal("Domain: google dan TLD: org"))
		})
	})

	When("email is 'admin@ruangguru.edu'", func() {
		It("should return 'Domain: ruangguru dan TLD: edu'", func() {
			Expect(main.EmailInfo("admin@ruangguru.edu")).To(Equal("Domain: ruangguru dan TLD: edu"))
		})
	})

	When("email is 'admin@netnetnet.xyz'", func() {
		It("should return 'Domain: netnetnet dan TLD: xyz'", func() {
			Expect(main.EmailInfo("admin@netnetnet.xyz")).To(Equal("Domain: netnetnet dan TLD: xyz"))
		})
	})

	When("email is 'admin@organisation.net'", func() {
		It("should return 'Domain: organisation dan TLD: net'", func() {
			Expect(main.EmailInfo("admin@organisation.edu")).To(Equal("Domain: organisation dan TLD: edu"))
		})
	})

	When("email is 'pt_mencari_cinta_sejatih_dunia_akherat@gmail.co.id'", func() {
		It("should return 'Domain: gmail dan TLD: co.id'", func() {
			Expect(main.EmailInfo("pt_mencari_cinta_sejatih_dunia_akherat@gmail.co.id")).To(Equal("Domain: gmail dan TLD: co.id"))

		})
	})
})
