package main_test

import (
	main "a21hc3NpZ25tZW50"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("", func() {
	Describe("HTTP Client GET", func() {
		When("hit API https://animechan.vercel.app/api/quotes/anime?title=naruto and decode to struct Animechan", func() {
			It("should return struct Animechan with contents of the hit API", func() {
				res, err := main.ClientGet()
				Expect(err).To(BeNil())
				Expect(res).To(HaveLen(10))
				Expect(res[0].Anime).To(Equal("Boruto: Naruto Next Generations"))
				Expect(res[0].Character).To(Equal("Naruto Uzumaki"))
				Expect(res[0].Quote).To(Equal("The many lives lost during long years of conflict... because of those selfless sacrifices, we are able to bathe in peace and prosperity now. To ingrain this history within the new generation will be a vital cog in helping to maintain the peace."))

				Expect(res[4].Anime).To(Equal("Naruto"))
				Expect(res[4].Character).To(Equal("Itachi Uchiha"))
				Expect(res[4].Quote).To(Equal("Those who forgive themselves, and are able to accept their true nature... They are the strong ones!"))

				Expect(res[9].Anime).To(Equal("Naruto"))
				Expect(res[9].Character).To(Equal("Neji Hyuuga"))
				Expect(res[9].Quote).To(Equal("Fear. That is what we live with. And we live it everyday. Only in death are we free of it."))
			})
		})
	})

	Describe("HTTP Client POST", func() {
		When("hit API https://postman-echo.com/post and decode to struct Postman", func() {
			It("should return struct Postman with contents of the hit API", func() {
				res, err := main.ClientPost()
				Expect(err).To(BeNil())
				Expect(res.Url).To(Equal("https://postman-echo.com/post"))
				Expect(res.Data.Name).To(Equal("Dion"))
				Expect(res.Data.Email).To(Equal("dionbe2022@gmail.com"))
			})
		})
	})
})
