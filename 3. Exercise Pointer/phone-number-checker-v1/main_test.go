package main_test

import (
	main "a21hc3NpZ25tZW50"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type TestData struct {
	Input    string
	Expected string
}

var _ = Describe("PhoneNumberChecker", func() {
	Context("Input is invalid number", func() {
		It("should insert result with 'invalid'", func() {
			testData := []TestData{
				{"1234567", "invalid"},
				{"123", "invalid"},
				{"08111111", "invalid"},
				{"6281234", "invalid"},
			}

			for _, test := range testData {
				var result string
				main.PhoneNumberChecker(test.Input, &result)
				Expect(result).To(Equal(test.Expected))
			}
		})
	})

	Context("Number code is 62811 - 63815 / 0811 - 0815", func() {
		It("should insert result with 'Telkomsel'", func() {
			testData := []TestData{
				{"081111111111", "Telkomsel"},
				{"081211111111", "Telkomsel"},
				{"081311111111", "Telkomsel"},
				{"081411111111", "Telkomsel"},
				{"081511111111", "Telkomsel"},
				{"6281111111111", "Telkomsel"},
				{"6281211111111", "Telkomsel"},
				{"6281311111111", "Telkomsel"},
			}

			for _, test := range testData {
				var result string
				main.PhoneNumberChecker(test.Input, &result)
				Expect(result).To(Equal(test.Expected))
			}
		})
	})

	Context("Number code is 62816 - 63819 / 0816 - 0819", func() {
		It("should insert result with 'Indosat'", func() {
			testData := []TestData{
				{"08163456123", "Indosat"},
				{"08173456123", "Indosat"},
				{"08183456123", "Indosat"},
				{"08193456123", "Indosat"},
				{"628163456123", "Indosat"},
				{"628173456123", "Indosat"},
				{"628183456123", "Indosat"},
			}

			for _, test := range testData {
				var result string
				main.PhoneNumberChecker(test.Input, &result)
				Expect(result).To(Equal(test.Expected))
			}
		})
	})

	Context("Number code is 62821 - 62823 / 0821 - 0823", func() {
		It("should insert result with 'XL'", func() {
			testData := []TestData{
				{"08213456123", "XL"},
				{"08223456123", "XL"},
				{"08233456123", "XL"},
				{"628213456123", "XL"},
				{"628223456123", "XL"},
			}

			for _, test := range testData {
				var result string
				main.PhoneNumberChecker(test.Input, &result)
				Expect(result).To(Equal(test.Expected))
			}
		})
	})

	Context("Number code is 62827 - 62829 / 0827 - 0829", func() {
		It("should insert result with 'Tri'", func() {
			testData := []TestData{
				{"08273456123", "Tri"},
				{"08283456123", "Tri"},
				{"08293456123", "Tri"},
				{"628273456123", "Tri"},
				{"628283456123", "Tri"},
			}

			for _, test := range testData {
				var result string
				main.PhoneNumberChecker(test.Input, &result)
				Expect(result).To(Equal(test.Expected))
			}
		})
	})

	Context("Number code is 62852 - 62853 / 0852 - 0853", func() {
		It("should insert result with 'AS'", func() {
			testData := []TestData{
				{"08523456123", "AS"},
				{"08533456123", "AS"},
				{"628523456123", "AS"},
				{"628533456123", "AS"},
			}

			for _, test := range testData {
				var result string
				main.PhoneNumberChecker(test.Input, &result)
				Expect(result).To(Equal(test.Expected))
			}
		})
	})

	Context("Number code is 62881 - 62888 / 0881 - 0888", func() {
		It("should insert result with 'Smartfren'", func() {
			testData := []TestData{
				{"08813456123", "Smartfren"},
				{"08823456123", "Smartfren"},
				{"08833456123", "Smartfren"},
				{"08843456123", "Smartfren"},
				{"628853456123", "Smartfren"},
				{"628863456123", "Smartfren"},
				{"628873456123", "Smartfren"},
			}

			for _, test := range testData {
				var result string
				main.PhoneNumberChecker(test.Input, &result)
				Expect(result).To(Equal(test.Expected))
			}
		})
	})

	Context("Input not contain code provider", func() {
		It("should insert result with 'invalid'", func() {
			testData := []TestData{
				{"08993456123", "invalid"},
				{"089134561234", "invalid"},
				{"628443456123", "invalid"},
				{"6284534561234", "invalid"},
			}

			for _, test := range testData {
				var result string
				main.PhoneNumberChecker(test.Input, &result)
				Expect(result).To(Equal(test.Expected))
			}
		})
	})
})
