package main_test

import (
	main "a21hc3NpZ25tZW50"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("CountVowelConsonant", func() {
	When("input str is 'Hidup Itu Indah'", func() {
		It("should return 6, 7, false", func() {
			vowel, cons, err := main.CountVowelConsonant("Hidup Itu Indah")
			Expect(vowel).To(Equal(6))
			Expect(cons).To(Equal(7))
			Expect(err).To(Equal(false))
		})
	})

	When("input str is 'Selamat pagi dunia yang indah dan penuh dengan semangat'", func() {
		It("should return 19, 28, false", func() {
			vowel, cons, err := main.CountVowelConsonant("Selamat pagi dunia yang indah dan penuh dengan semangat")
			Expect(vowel).To(Equal(19))
			Expect(cons).To(Equal(28))
			Expect(err).To(Equal(false))
		})
	})

	When("input str is 'HALO DUNIA INDAH YANG BAIK'", func() {
		It("should return 10, 12, false", func() {
			vowel, cons, err := main.CountVowelConsonant("HALO DUNIA INDAH YANG BAIK")
			Expect(vowel).To(Equal(10))
			Expect(cons).To(Equal(12))
			Expect(err).To(Equal(false))
		})
	})

	When("input str is 'SEMANGAT PAGI, itu kata orang yang baru datang ke rumahku'", func() {
		It("should return 20, 28, false", func() {
			vowel, cons, err := main.CountVowelConsonant("SEMANGAT PAGI, itu kata orang yang baru datang ke rumahku")
			Expect(vowel).To(Equal(20))
			Expect(cons).To(Equal(27))
			Expect(err).To(Equal(false))
		})
	})

	When("input str is 'bbbbbbbb vvvvvvvv  ddddd sssss  wwww'", func() {
		It("should return 0, 30, true", func() {
			vowel, cons, err := main.CountVowelConsonant("bbbbbbbb vvvvvvvv  ddddd sssss  wwww")
			Expect(vowel).To(Equal(0))
			Expect(cons).To(Equal(30))
			Expect(err).To(Equal(true))
		})
	})

	When("input str is text of lorem ipsum", func() {
		It("Should counting vowel and consonant letter in lorem ipsum text", func() {
			var lorem = "Lorem ipsum dolor sit amet consectetur adipiscing elit sed do eiusmod tempor incididunt ut labore et dolore magna aliqua Volutpat lacus laoreet non curabitur Amet risus nullam eget felis eget nunc Amet risus nullam eget felis eget Consectetur purus ut faucibus pulvinar elementum Suspendisse ultrices gravida dictum fusce ut placerat orci nulla pellentesque Eu nisl nunc mi ipsum Ac turpis egestas maecenas pharetra convallis Egestas tellus rutrum tellus pellentesque eu tincidunt tortor Nunc sed velit dignissim sodales ut Ut venenatis tellus in metus vulputate eu scelerisque felis Feugiat pretium nibh ipsum consequat nisl vel Porta non pulvinar neque laoreet suspendisse interdum consectetur Sed elementum tempus egestas sed sed risus Blandit volutpat maecenas volutpat blandit aliquam etiam erat Porta non pulvinar neque laoreet suspendisse interdum Vel fringilla est ullamcorper eget nulla facilisi etiam dignissim diam Nisl suscipit adipiscing bibendum est ultricies integer quis auctor Vitae tempus quam pellentesque nec nam aliquam sem Malesuada bibendum arcu vitae elementum curabitur vitae nunc sed Quam elementum pulvinar etiam non quam lacus suspendisse Nunc mattis enim ut tellus elementum sagittis vitae et Ultrices eros in cursus turpis massa tincidunt Morbi tincidunt ornare massa eget egestas purus viverra Massa ultricies mi quis hendrerit dolor magna eget Ut eu sem integer vitae justo eget magna fermentum iaculis Diam sit amet nisl suscipit adipiscing Convallis aenean et tortor at risus viverra Nulla posuere sollicitudin aliquam ultrices sagittis orci a Quam viverra orci sagittis eu volutpat odio facilisis mauris Libero volutpat sed cras ornare arcu dui vivamus Hendrerit dolor magna eget est lorem ipsum dolor sit amet Congue eu consequat ac felis donec et odio Etiam tempor orci eu lobortis elementum nibh tellus At auctor urna nunc id cursus Velit egestas dui id ornare arcu odio ut"

			vowel, cons, err := main.CountVowelConsonant(lorem)

			Expect(vowel).To(Equal(693))
			Expect(cons).To(Equal(931))
			Expect(err).To(Equal(false))
		})
	})
})
