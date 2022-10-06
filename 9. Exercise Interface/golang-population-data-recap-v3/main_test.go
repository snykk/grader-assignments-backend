package main_test

import (
	main "a21hc3NpZ25tZW50"
	"sort"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type TestData struct {
	Input    []string
	Expected []map[string]interface{}
}

var getKeys = func(m map[string]interface{}) []string {
	keys := []string{}
	for k := range m {
		keys = append(keys, k)
	}

	// sorting
	sort.Strings(keys)

	return keys
}

var _ = Describe("PopulationData", func() {
	When("input data is empty", func() {
		It("should return empty map", func() {
			res := main.PopulationData([]string{})

			Expect(res).To(Equal([]map[string]interface{}{}))
			Expect(res).To(BeEmpty())
		})
	})

	When("input data containing full data format", func() {
		It("should return map of string with all key needed", func() {
			testData := []TestData{
				{
					[]string{
						"Budi;23;Jakarta;160.42;true",
						"Joko;30;Bandung;160;true",
						"Susi;25;Bogor;165.42;false",
					},
					[]map[string]interface{}{
						{"name": "Budi", "age": 23, "address": "Jakarta", "height": 160.42, "isMarried": true},
						{"name": "Joko", "age": 30, "address": "Bandung", "height": 160.00, "isMarried": true},
						{"name": "Susi", "age": 25, "address": "Bogor", "height": 165.42, "isMarried": false},
					},
				},
				{
					[]string{
						"Budi;23;Jakarta;160.42;true",
						"Joko;30;Bandung;160;true",
						"Susi;25;Bogor;165.42;false",
						"Santi;27;Jakarta;160;true",
						"Rudi;23;Jakarta;161.1;false",
						"Rudi 2;23;Jakarta;166.5;false",
						"Rudi 3;23;Jakarta;170;false"},
					[]map[string]interface{}{
						{"name": "Budi", "age": 23, "address": "Jakarta", "height": 160.42, "isMarried": true},
						{"name": "Joko", "age": 30, "address": "Bandung", "height": 160.0, "isMarried": true},
						{"name": "Susi", "age": 25, "address": "Bogor", "height": 165.42, "isMarried": false},
						{"name": "Santi", "age": 27, "address": "Jakarta", "height": 160.0, "isMarried": true},
						{"name": "Rudi", "age": 23, "address": "Jakarta", "height": 161.1, "isMarried": false},
						{"name": "Rudi 2", "age": 23, "address": "Jakarta", "height": 166.5, "isMarried": false},
						{"name": "Rudi 3", "age": 23, "address": "Jakarta", "height": 170.0, "isMarried": false},
					},
				},
			}

			for _, test := range testData {
				actual := main.PopulationData(test.Input)

				for i, res := range actual {
					actualKey := getKeys(res)
					expectedKey := getKeys(test.Expected[i])

					Expect(actualKey).To(Equal(expectedKey))
					for k, v := range res {
						Expect(v).To(Equal(test.Expected[i][k]))
					}
				}

				Expect(actual).ToNot(BeEmpty())
			}
		})
	})

	When("input data containing optional data format or not", func() {
		It("should return map of string with the appropriate key", func() {
			testData := []TestData{
				{
					[]string{
						"Budi;23;Jakarta;160.42;true",
						"Joko;30;Bandung;;true",
						"Susi;25;Bogor;;false",
						"Santi;27;Jakarta;160;",
						"Rudi;23;Jakarta;161.1;",
						"Rudi;23;Jakarta;166.5;false",
						"Rudi;23;Jakarta;;",
					}, []map[string]interface{}{
						{"name": "Budi", "age": 23, "address": "Jakarta", "height": 160.42, "isMarried": true},
						{"name": "Joko", "age": 30, "address": "Bandung", "isMarried": true},
						{"name": "Susi", "age": 25, "address": "Bogor", "isMarried": false},
						{"name": "Santi", "age": 27, "address": "Jakarta", "height": 160.0},
						{"name": "Rudi", "age": 23, "address": "Jakarta", "height": 161.1},
						{"name": "Rudi", "age": 23, "address": "Jakarta", "height": 166.5, "isMarried": false},
						{"name": "Rudi", "age": 23, "address": "Jakarta"},
					}},
			}

			for _, test := range testData {
				actual := main.PopulationData(test.Input)

				for i, res := range actual {
					actualKey := getKeys(res)
					expectedKey := getKeys(test.Expected[i])

					Expect(actualKey).To(Equal(expectedKey))
					for k, v := range res {
						Expect(v).To(Equal(test.Expected[i][k]))
					}
				}

				Expect(actual).ToNot(BeEmpty())
			}
		})
	})
})
