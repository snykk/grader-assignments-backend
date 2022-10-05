package main_test

import (
	main "a21hc3NpZ25tZW50"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("FindSimilarData", func() {
	When("input is 'iphone' and data is 'laptop, iphone 13, iphone 12, iphone 12 pro'", func() {
		It("should return 'iphone 13,iphone 12,iphone 12 pro'", func() {
			Expect(main.FindSimilarData("iphone", "laptop", "iphone 13", "iphone 12", "iphone 12 pro")).To(Equal("iphone 13,iphone 12,iphone 12 pro"))
		})
	})

	When("input is 'mobil' and data is 'mobil APV, mobil Avanza, motor matic, motor gede'", func() {
		It("should return 'mobil APV,mobil Avanza'", func() {
			Expect(main.FindSimilarData("mobil", "mobil APV", "mobil Avanza", "motor matic", "motor gede")).To(Equal("mobil APV,mobil Avanza"))
		})
	})

	When("input is 'motor'", func() {
		It("should return similar data with 'motor'", func() {
			Expect(main.FindSimilarData("motor", "mobil APV", "mobil Avanza", "motor matic", "motor gede", "iphone 14", "iphone 13", "iphone 12", "pengering baju", "Kemeja flannel")).To(Equal("motor matic,motor gede"))
		})
	})

	When("input is 'motor'", func() {
		It("should return similar data with 'motor'", func() {
			Expect(main.FindSimilarData("motor", "motor suzuka doraemon", "iphone 14", "iphone 13", "iphone 12", "pengering baju", "Kemeja flannel")).To(Equal("motor suzuka doraemon"))
		})
	})

	When("input is 'rice cooker'", func() {
		It("should return similar data with 'rice cooker'", func() {
			Expect(main.FindSimilarData("rice cooker", "rice cooker hisanitsi", "iphone 13", "iphone 12", "pengering baju", "Kemeja flannel")).To(Equal("rice cooker hisanitsi"))
		})
	})
})
