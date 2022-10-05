package main_test

import (
	main "a21hc3NpZ25tZW50"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("SlurredTalk", func() {
	When("parameter pointer of string containing letter 'S'", func() {
		It("should replace 'S' with 'L'", func() {
			var words string = "Semangat pagi semuanya, semoga sehat selalu. Sehingga selalu bisa senyum"
			main.SlurredTalk(&words)
			Expect(words).To(Equal("Lemangat pagi lemuanya, lemoga lehat lelalu. Lehingga lelalu bila lenyum"))
		})
	})

	When("parameter pointer of string containing letter 'R'", func() {
		It("should replace 'R' with 'L'", func() {
			var words string = "Remahan roti menemani kita dengan rindu-rindu roman. Roti, ramen, rasa, rindu, rindang, roman menjadi kenikmatanku"
			main.SlurredTalk(&words)
			Expect(words).To(Equal("Lemahan loti menemani kita dengan lindu-lindu loman. Loti, lamen, lala, lindu, lindang, loman menjadi kenikmatanku"))
		})
	})

	When("parameter pointer of string containing letter 'Z'", func() {
		It("should replace 'Z' with 'L'", func() {
			var words string = "Zakat, zebra, zabib, zabur, zadah, zahid, zai, zaim, zair, zaitun, zakar, zakelek, zakiah, zakum, zal, zalim, zalir, zaman, zamin, zamindar, zamrud, zamzam, zan, zanggi, zantara, zarafah, zarah, zaratit, zat, zatua"
			main.SlurredTalk(&words)
			Expect(words).To(Equal("Lakat, lebla, labib, labul, ladah, lahid, lai, laim, lail, laitun, lakal, lakelek, lakiah, lakum, lal, lalim, lalil, laman, lamin, lamindal, lamlud, lamlam, lan, langgi, lantala, lalafah, lalah, lalatit, lat, latua"))
		})
	})

	When("parameter pointer of string cointain containing 'S', 'R' and 'Z'", func() {
		It("should replace 'S', 'R' and 'Z' with 'L'", func() {
			var words string = "srzfdsfkewfopedjjldfjdsljfdsjflksdjierueiuSKOWNREKRZ:LKJIOJKLSJKFJKLSJFKLiqopeiqoeiqojnnckjzcnuiuqiwoeuqeiLKNQNPKJNXZMRHSIUYRHJKSHHIFIRHNXNJSNkasljdfjskdfjsdfojeiorueiowruewiojdknscnskjcnsdjkfdfjkshfuweryieorqwpoewqopeiqfhsdvnskgghoerutrwiojsljgnsvjnsjvnsdjvdsjn"
			main.SlurredTalk(&words)
			Expect(words).To(Equal("lllfdlfkewfopedjjldfjdlljfdljflkldjielueiuLKOWNLEKLL:LKJIOJKLLJKFJKLLJFKLiqopeiqoeiqojnnckjlcnuiuqiwoeuqeiLKNQNPKJNXLMLHLIUYLHJKLHHIFILHNXNJLNkalljdfjlkdfjldfojeiolueiowluewiojdknlcnlkjcnldjkfdfjklhfuwelyieolqwpoewqopeiqfhldvnlkgghoelutlwiojlljgnlvjnljvnldjvdljn"))
		})
	})
})
