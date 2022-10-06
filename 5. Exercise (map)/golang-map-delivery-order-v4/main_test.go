package main_test

import (
	main "a21hc3NpZ25tZW50"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type TestParam struct {
	Data []string
	Day  string
}

type TestData struct {
	Test     TestParam
	Expected map[string]float32
}

var _ = Describe("Main", func() {
	Describe("DeliveryOrderDay", func() {
		When("data is empty", func() {
			It("should return empty map", func() {
				res := main.DeliveryOrder([]string{}, "senin")

				Expect(res).NotTo(BeNil())
				Expect(res).To(BeEmpty())
			})
		})

		When("delivery location not contain delivery day", func() {
			It("should return empty map", func() {
				res := main.DeliveryOrder([]string{"Andi:Sukirman:15000:BKS"}, "senin")

				Expect(res).NotTo(BeNil())
				Expect(res).To(Equal(map[string]float32{}))
			})
		})

		When("delivery day is minggu", func() {
			It("should return empty map", func() {
				res := main.DeliveryOrder([]string{"Andi:Sukirman:15000:DPK", "Anggi:Anggraini:20000:BDG", "Andi:Gunawan:40000:BKS", "Budi:Gunawan:50000:JKT"}, "minggu")

				Expect(res).NotTo(BeNil())
				Expect(res).To(Equal(map[string]float32{}))
			})
		})
		When("data contain delivery day", func() {
			It("should eliminate delivery location data and get map total cost delivery", func() {
				var listTestData = []TestData{

					{
						Test: TestParam{
							Data: []string{"Anggi:Anggraini:20000:DPK", "Andi:Sukirman:15000:JKT"},
							Day:  "selasa",
						},
						Expected: map[string]float32{
							"Anggi-Anggraini": 21000, "Andi-Sukirman": 15750,
						},
					},
					{
						Test: TestParam{
							Data: []string{"Budi:Gunawan:10000:JKT", "Andi:Sukirman:20000:JKT", "Budi:Sukirman:30000:BDG", "Andi:Gunawan:40000:BKS", "Budi:Gunawan:50000:DPK"},
							Day:  "sabtu",
						},
						Expected: map[string]float32{
							"Andi-Sukirman": 21000, "Budi-Gunawan": 10500, "Budi-Sukirman": 31500,
						},
					},
				}

				for i := 0; i < len(listTestData); i++ {
					res := main.DeliveryOrder(listTestData[i].Test.Data, listTestData[i].Test.Day)

					Expect(res).NotTo(BeNil())

					for k := range listTestData[i].Expected {
						Expect(res).To(HaveKey(k))
						Expect(res[k]).To(Equal(listTestData[i].Expected[k]))
					}
				}
			})
		})

		It("should eliminate big delivery location data and get map total cost delivery 1", func() {
			var testData = TestData{
				Test: TestParam{
					Data: []string{
						"Durward:Miller:20000:BKS", "Nora:Wuckert:37000:BDG", "Candido:Mitchell:30000:JKT", "Ryley:Kozey:20000:BDG", "Kiley:Keebler:50000:JKT", "Johnpaul:Bashirian:45000:JKT", "Megane:Ortiz:35000:JKT", "Vida:Spinka:15000:DPK", "Brielle:Gaylord:63000:JKT", "Marcelle:Heaney:37000:DPK", "Randal:Beatty:37000:DPK", "Donnell:Donnelly:50000:BKS", "Clifford:Goyette:45000:BKS", "Coby:Berge:75000:DPK", "Lina:Toy:63000:DPK", "Nia:Labadie:10000:BKS", "Mina:Funk:35000:BKS", "Dante:Parisian:50000:BDG", "Daniela:Hansen:35000:BDG", "Beaulah:Stehr:20000:DPK", "Sheldon:Luettgen:20000:BDG", "Catherine:Hoeger:50000:BKS", "Kennedi:Considine:35000:JKT", "Cathrine:Carter:20000:BKS", "Julius:Paucek:37000:BDG", "Lloyd:Baumbach:20000:BKS", "Tremaine:Simonis:35000:JKT", "Arthur:Hettinger:55500:BDG", "Kirsten:Bosco:37000:DPK", "Joana:Cremin:35000:DPK", "Julia:Stanton:30000:JKT", "Morton:Reichert:10000:BKS", "Mario:Ruecker:10000:JKT", "Jerad:Shanahan:15000:BKS", "Jillian:Littel:37000:JKT", "Angel:Dibbert:45000:BKS", "Dedric:Grant:50000:JKT", "Eliezer:Kunze:55500:BKS", "Joesph:Lehner:50000:DPK", "Alia:Aufderhar:50000:DPK", "Florencio:Nader:20000:BDG", "Lucas:Koepp:37000:BKS", "Alyson:Dickens:50000:BKS", "Maxie:Breitenberg:15000:BKS", "Giovanny:Davis:55500:BDG", "Darrick:Sanford:45000:BKS", "Efren:Schimmel:15000:BDG", "Faustino:Stamm:63000:BDG", "Verner:Heidenreich:35000:BDG", "Alisha:Stroman:10000:BDG", "Quinton:Johnston:10000:JKT", "Israel:Hane:15000:DPK", "Diana:Miller:35000:JKT", "Alexandro:Leannon:55500:JKT", "Keyon:Pagac:35000:JKT", "Jesse:Weissnat:10000:JKT", "Taya:Kiehn:30000:JKT", "Armando:Hegmann:20000:DPK", "Janessa:Mueller:45000:JKT", "Forest:Carroll:30000:BKS", "Walter:Bechtelar:63000:DPK", "Zachariah:Krajcik:55500:DPK", "Pattie:Nader:20000:DPK", "Reta:Keeling:50000:BDG", "Mavis:Hagenes:75000:BKS", "Kelvin:Leffler:15000:BKS", "Abigale:Rowe:20000:BDG", "Emmitt:Kshlerin:20000:BDG", "Rosie:Stracke:75000:BDG", "Karlie:Labadie:15000:JKT", "Carmel:Beatty:45000:JKT", "Reagan:Cronin:75000:JKT", "Savion:Kunde:55500:BDG", "Zackary:Russel:20000:JKT", "Bettie:Bergnaum:20000:JKT",
					},
					Day: "selasa",
				},
				Expected: map[string]float32{
					"Catherine-Hoeger": 52500, "Lucas-Koepp": 38850, "Israel-Hane": 15750, "Mina-Funk": 36750, "Diana-Miller": 36750, "Coby-Berge": 78750, "Jerad-Shanahan": 15750, "Eliezer-Kunze": 58275, "Maxie-Breitenberg": 15750, "Carmel-Beatty": 47250, "Randal-Beatty": 38850, "Quinton-Johnston": 10500, "Forest-Carroll": 31500, "Megane-Ortiz": 36750, "Beaulah-Stehr": 21000, "Mavis-Hagenes": 78750, "Armando-Hegmann": 21000, "Kirsten-Bosco": 38850, "Jillian-Littel": 38850, "Alia-Aufderhar": 52500, "Alyson-Dickens": 52500, "Kelvin-Leffler": 15750, "Candido-Mitchell": 31500, "Vida-Spinka": 15750, "Brielle-Gaylord": 66150, "Marcelle-Heaney": 38850, "Dedric-Grant": 52500, "Bettie-Bergnaum": 21000, "Lloyd-Baumbach": 21000, "Tremaine-Simonis": 36750, "Keyon-Pagac": 36750, "Jesse-Weissnat": 10500, "Taya-Kiehn": 31500, "Durward-Miller": 21000, "Nia-Labadie": 10500, "Kiley-Keebler": 52500, "Donnell-Donnelly": 52500, "Clifford-Goyette": 47250, "Mario-Ruecker": 10500, "Joana-Cremin": 36750, "Pattie-Nader": 21000, "Morton-Reichert": 10500, "Janessa-Mueller": 47250, "Karlie-Labadie": 15750, "Darrick-Sanford": 47250, "Alexandro-Leannon": 58275, "Zachariah-Krajcik": 58275, "Kennedi-Considine": 36750, "Cathrine-Carter": 21000, "Julia-Stanton": 31500, "Angel-Dibbert": 47250, "Joesph-Lehner": 52500, "Johnpaul-Bashirian": 47250, "Lina-Toy": 66150, "Walter-Bechtelar": 66150, "Reagan-Cronin": 78750, "Zackary-Russel": 21000,
				},
			}

			res := main.DeliveryOrder(testData.Test.Data, testData.Test.Day)

			Expect(res).NotTo(BeNil())
			Expect(len(res)).To(Equal(len(testData.Expected)))

			for k := range testData.Expected {
				Expect(res).To(HaveKey(k))
				Expect(res[k]).To(Equal(testData.Expected[k]))
			}
		})

		It("should eliminate big delivery location data and get map total cost delivery 2", func() {
			var testData = TestData{
				Test: TestParam{
					Data: []string{
						"Hettie:Ryan:45000:DPK", "Jamarcus:Rohan:55500:BDG", "Maegan:Crona:30000:BKS", "Mohammad:Rodriguez:37000:BDG", "Diana:Schaden:35000:BDG", "Roscoe:Jacobi:37000:BDG", "Royce:Sporer:15000:BKS", "Elian:Bashirian:45000:BDG", "Guiseppe:Kub:15000:BDG", "Nettie:Pollich:55500:BDG", "Johathan:Kozey:20000:JKT", "Everardo:DuBuque:50000:JKT", "Kelsi:Donnelly:55500:BDG", "Winston:Hyatt:63000:JKT", "Thad:Mitchell:63000:DPK", "Elda:Swift:75000:BKS", "Braeden:Kshlerin:50000:DPK", "Imani:Hoppe:63000:BKS", "Freddie:Considine:45000:BKS", "Margaretta:Feil:30000:DPK", "Zack:King:37000:BKS", "Jaden:Labadie:45000:BDG", "Preston:Schamberger:37000:BDG", "Cecilia:Barrows:15000:BDG", "Hadley:Casper:30000:BDG", "Gonzalo:Cormier:50000:BDG", "Mariano:Mueller:63000:DPK", "Loraine:Daniel:55500:JKT", "Gudrun:Quitzon:45000:DPK", "Carmine:Goyette:55500:DPK", "Jovani:Howe:50000:JKT", "Veronica:Bogisich:20000:BKS", "Sierra:Torp:30000:JKT", "Alan:Spencer:15000:BDG", "Hellen:Kuvalis:55500:DPK", "Elisabeth:Miller:20000:BDG", "Brigitte:Mraz:35000:BDG", "Dagmar:Gutkowski:30000:BDG", "Pearline:White:75000:DPK", "Bethel:Wilderman:55500:JKT", "Hilma:Batz:50000:DPK", "Lorenza:Kuvalis:20000:BKS", "Arturo:Jenkins:75000:BKS", "Tyrique:Klocko:15000:JKT", "Frankie:Morissette:30000:JKT", "Chanel:Nienow:50000:BDG", "Malvina:Lind:55500:BDG", "Rhea:Krajcik:63000:JKT", "Dorian:Dickinson:35000:DPK", "Noemi:Schmidt:45000:DPK", "Luella:Hand:35000:DPK", "Major:Bailey:20000:DPK", "Franco:Kuhn:15000:DPK", "Adella:Schiller:35000:BKS", "Cory:Rice:37000:BDG", "Bernadine:Dibbert:30000:BDG", "Mavis:Tillman:63000:BDG", "Antonia:Blick:20000:BDG", "Selina:Cole:55500:BKS", "Jacky:Hackett:10000:DPK", "Delphine:Hilll:15000:DPK", "Kris:Goyette:10000:JKT", "Stuart:Yost:45000:JKT", "Haleigh:Larson:35000:DPK", "Brice:Lueilwitz:75000:DPK", "Maia:Kreiger:50000:DPK", "Dejon:Deckow:45000:BDG", "Letha:Gislason:45000:BKS", "Brigitte:Walker:45000:DPK", "Ellsworth:Turcotte:15000:DPK", "Hillary:Bechtelar:55500:BKS", "Stan:Cremin:55500:BDG", "Layla:Fritsch:37000:BDG", "Rachael:Farrell:63000:JKT", "Reba:Abshire:63000:DPK", "Thomas:Reynolds:10000:BDG", "Jameson:Grant:10000:BDG", "Jimmy:Lockman:75000:BDG", "Kraig:Zemlak:45000:DPK", "Roslyn:Stroman:45000:BKS", "Glen:Murazik:50000:DPK", "Marisa:Lockman:37000:DPK", "Rene:Mueller:10000:DPK", "Corrine:Douglas:15000:DPK", "Darian:Schumm:55500:BKS", "Trever:Wehner:50000:BKS", "Maegan:Hartmann:45000:BKS", "Gaylord:Crooks:75000:BDG", "Forrest:Wiegand:35000:BKS", "Jeramy:Mosciski:35000:JKT", "Kay:Strosin:45000:JKT", "Anne:Emard:37000:JKT", "Chad:Cassin:35000:JKT", "Finn:Casper:10000:JKT", "Lysanne:Eichmann:63000:JKT", "Unique:Denesik:63000:DPK", "Pearlie:Parisian:35000:JKT", "Patrick:Zulauf:35000:BKS", "Amya:Wiegand:37000:BKS", "Romaine:Berge:30000:BKS", "Russ:Anderson:63000:BKS", "Edgar:Considine:37000:BKS", "Justina:Schuster:75000:BKS", "Loyal:Doyle:35000:BDG", "Hans:Turner:37000:BKS", "Cheyanne:Halvorson:35000:JKT", "Billie:DuBuque:30000:JKT", "Amari:McClure:63000:JKT", "Queen:Stanton:63000:JKT", "Dee:Thiel:15000:JKT", "Jailyn:Kilback:10000:JKT", "Hilbert:Kuvalis:10000:JKT", "Tre:Hudson:45000:DPK", "Schuyler:McCullough:30000:BKS", "Lonnie:Schowalter:10000:JKT", "Roderick:Greenholt:50000:DPK", "Alan:Parker:20000:BDG", "Raina:Kiehn:55500:BDG", "Kristofer:Ziemann:10000:BKS", "Madelynn:Kozey:15000:BDG", "Nicolas:Fadel:55500:JKT", "Lamar:Spencer:20000:JKT", "Vivian:Kuphal:75000:DPK", "Luigi:Predovic:35000:JKT", "Arturo:Buckridge:10000:DPK", "Hal:Stehr:37000:BKS", "Zaria:Walsh:30000:BDG", "Trevor:Dicki:30000:BKS", "Beth:Wilderman:45000:BDG", "Donna:Blick:10000:BKS", "Maritza:Marquardt:30000:BDG", "Lesly:Durgan:45000:JKT", "Maritza:Fisher:30000:DPK", "Ana:Boyer:75000:BDG", "Immanuel:Howell:15000:DPK", "Josiah:Bailey:20000:BKS", "Leonie:Weissnat:30000:BKS", "Talon:Muller:75000:DPK", "Emmett:Rice:55500:BKS", "Danielle:Hessel:45000:BDG", "Cathryn:Reynolds:63000:DPK", "Dejuan:Raynor:75000:BDG", "Austen:Walter:30000:BDG", "Rod:OReilly:35000:JKT", "Brisa:Osinski:75000:BKS", "Verda:Sipes:35000:BKS", "Dandre:Legros:75000:DPK", "Nola:Douglas:50000:DPK", "Mariam:Upton:30000:BKS", "Eduardo:Jacobi:30000:BDG", "Paula:Satterfield:45000:JKT", "Alexie:Bernier:15000:JKT", "Ellie:Bergstrom:55500:JKT", "Idell:Weber:50000:JKT", "Lonnie:Becker:55500:DPK", "Leif:Veum:55500:JKT", "Boris:Zieme:45000:DPK", "Valerie:Haley:55500:BDG", "Noemi:Jenkins:30000:BKS", "Nelle:Schmeler:55500:BDG", "Ted:Quitzon:55500:BDG", "Lelia:Kerluke:30000:BKS", "Talon:Dietrich:63000:BDG", "Michaela:Kilback:75000:BDG", "Emmalee:Yundt:75000:BKS", "Ivah:Gorczany:63000:JKT", "Madge:Sipes:30000:JKT", "August:Feeney:45000:DPK", "Forest:Olson:50000:DPK", "Glen:Kertzmann:55500:JKT", "Stan:Borer:15000:DPK", "Mable:Marks:50000:BDG", "Deron:Hartmann:37000:JKT", "Jayce:Nader:10000:BKS", "Stuart:Hammes:20000:BKS", "Novella:Stark:30000:BKS", "Luigi:Langworth:15000:BKS", "Yaamir:Collins:37000:BKS", "Jayden:Cummerata:35000:BKS", "Kellie:Rice:63000:DPK", "Ozella:Bashirian:55500:BKS", "Ted:Bins:30000:JKT", "Kellen:Langosh:50000:JKT", "Milan:Feest:75000:JKT", "Marisol:Murphy:50000:JKT", "Arlie:Treutel:50000:JKT", "Jared:Parker:55500:JKT", "Leda:Rodriguez:45000:JKT", "Salma:Graham:20000:BDG", "Jackie:OKon:50000:BKS", "Elbert:Hilll:45000:BKS", "Germaine:OReilly:55500:BKS", "Genoveva:Waelchi:55500:BKS", "Melisa:Altenwerth:50000:DPK", "River:Morissette:35000:DPK", "Darwin:McClure:50000:DPK", "Landen:Stark:30000:DPK", "Ida:Upton:55500:JKT", "Ambrose:Bergnaum:15000:JKT", "Vinnie:Rempel:30000:BDG", "Santos:Brekke:35000:BDG", "Stone:Hane:45000:BDG", "Clarabelle:Christiansen:30000:BDG", "Maryam:Shields:55500:DPK", "Jacky:Kozey:35000:JKT", "Julian:Brekke:30000:DPK", "Vergie:Koss:45000:JKT", "Merritt:Stracke:45000:JKT", "Joseph:Windler:37000:BDG", "Emmanuelle:Sipes:55500:BDG", "Addie:Wilderman:10000:DPK", "Douglas:Steuber:55500:DPK", "Garrison:Grimes:50000:BDG", "Shanel:Price:63000:BDG", "Torrance:Runolfsson:37000:BDG", "Marina:Johnston:20000:BDG", "Maggie:Kiehn:63000:BDG", "Russell:Littel:37000:BDG", "Otto:Lehner:15000:BDG", "Meta:Runolfsdottir:15000:DPK", "Audreanne:McLaughlin:75000:JKT", "Gudrun:Hirthe:35000:JKT", "Zelma:Gutmann:20000:JKT", "Samanta:Kemmer:55500:JKT", "Chanelle:Cruickshank:15000:JKT", "Joan:Sauer:20000:BDG", "Dorian:Conn:75000:BDG", "Trace:Wyman:75000:DPK", "Aditya:Goyette:50000:BKS", "Arden:Bayer:45000:BDG", "Javonte:Hermiston:37000:BDG", "Emelia:Beatty:50000:BKS", "Narciso:Heller:55500:BDG", "Jacklyn:Grimes:10000:JKT", "Braxton:Metz:15000:BDG", "Susanna:Keebler:55500:BDG", "Kaia:Swaniawski:55500:BDG", "Karina:Mante:30000:DPK", "Linda:Daniel:55500:DPK", "Ford:Kertzmann:75000:DPK", "Cyril:Breitenberg:30000:DPK", "Theresa:Hilpert:15000:DPK", "Milford:Gleason:30000:JKT", "Loyce:Feeney:35000:DPK", "Kiana:Marquardt:50000:JKT", "Gail:Tillman:35000:BDG", "Loraine:Abshire:35000:BDG", "Kennedy:Grady:15000:BDG", "Emil:Moore:30000:BKS", "Rita:Lind:30000:BKS", "Samir:Jenkins:10000:BKS", "Marcus:OKeefe:30000:BKS", "Roy:Kreiger:37000:BKS", "Kristofer:Kuvalis:50000:JKT", "Queen:Veum:55500:JKT", "Malvina:Stroman:50000:JKT", "Glenna:Dach:75000:DPK", "Kyleigh:Dicki:20000:DPK", "Margarete:Bernier:45000:DPK", "Stanford:Hettinger:15000:DPK", "Garland:VonRueden:50000:DPK", "Ova:Franecki:10000:DPK", "Laisha:Shanahan:45000:BKS", "Orlo:Green:35000:DPK", "Laurie:Skiles:35000:BDG", "Michelle:OKeefe:15000:BDG", "Germaine:Thompson:50000:BDG", "Werner:Schneider:55500:BDG", "Mariana:Bode:37000:JKT", "Gabriella:Ebert:75000:JKT", "Brionna:Hand:55500:DPK", "Donnell:Vandervort:63000:DPK", "Odell:Balistreri:55500:DPK", "Berta:Reichel:37000:JKT", "Bud:Ward:63000:DPK", "Sandy:Conroy:45000:DPK", "Esperanza:Rath:63000:JKT", "Ambrose:Emard:63000:DPK", "Adonis:McCullough:20000:JKT", "Erin:Cremin:45000:DPK", "Malachi:Kunze:50000:BDG", "Suzanne:Williamson:10000:BDG", "Samantha:Corwin:37000:BKS", "Eli:Wisozk:20000:BDG", "Brent:Johnston:55500:BKS", "Mohamed:Cummings:20000:DPK", "Floy:Pfannerstill:55500:JKT", "Ana:Gaylord:37000:DPK", "Tito:Herzog:50000:DPK", "Alfred:Emard:10000:JKT", "Madisyn:Green:30000:DPK", "Everett:Abbott:50000:DPK", "Sasha:Krajcik:10000:DPK", "Dagmar:Bernier:20000:BKS", "Freddy:Marks:30000:BDG", "Nelda:Dicki:20000:BKS", "Jimmy:Huels:35000:BKS", "Ardella:Klocko:37000:BKS", "Anthony:Stoltenberg:10000:BKS", "William:Schinner:50000:BKS",
					},
					Day: "sabtu",
				},
				Expected: map[string]float32{
					"Loyal-Doyle": 36750, "Dee-Thiel": 15750, "Paula-Satterfield": 47250, "Torrance-Runolfsson": 38850, "Tyrique-Klocko": 15750, "Alfred-Emard": 10500, "Glen-Kertzmann": 58275, "Emmanuelle-Sipes": 58275, "Cory-Rice": 38850, "Lamar-Spencer": 21000, "Rod-OReilly": 36750, "Idell-Weber": 52500, "Maggie-Kiehn": 66150, "Audreanne-McLaughlin": 78750, "Sierra-Torp": 31500, "Brigitte-Mraz": 36750, "Layla-Fritsch": 38850, "Amari-McClure": 66150, "Diana-Schaden": 36750, "Alan-Parker": 21000, "Stone-Hane": 47250, "Ambrose-Bergnaum": 15750, "Shanel-Price": 66150, "Joan-Sauer": 21000, "Mohammad-Rodriguez": 38850, "Thomas-Reynolds": 10500, "Leif-Veum": 58275, "Ted-Quitzon": 58275, "Raina-Kiehn": 58275, "Jacky-Kozey": 36750, "Kristofer-Kuvalis": 52500, "Laurie-Skiles": 36750, "Adonis-McCullough": 21000, "Frankie-Morissette": 31500, "Pearlie-Parisian": 36750, "Nicolas-Fadel": 58275, "Lesly-Durgan": 47250, "Valerie-Haley": 58275, "Santos-Brekke": 36750, "Vergie-Koss": 47250, "Roscoe-Jacobi": 38850, "Mavis-Tillman": 66150, "Austen-Walter": 31500, "Ellie-Bergstrom": 58275, "Dagmar-Gutkowski": 31500, "Gonzalo-Cormier": 52500, "Antonia-Blick": 21000, "Dejuan-Raynor": 78750, "Javonte-Hermiston": 38850, "Jacklyn-Grimes": 10500, "Mariana-Bode": 38850, "Billie-DuBuque": 31500, "Chanelle-Cruickshank": 15750, "Malvina-Stroman": 52500, "Germaine-Thompson": 52500, "Cecilia-Barrows": 15750, "Russell-Littel": 38850, "Floy-Pfannerstill": 58275, "Zelma-Gutmann": 21000, "Jaden-Labadie": 47250, "Zaria-Walsh": 31500, "Michaela-Kilback": 78750, "Marisol-Murphy": 52500, "Milan-Feest": 78750, "Kay-Strosin": 47250, "Hilbert-Kuvalis": 10500, "Madelynn-Kozey": 15750, "Luigi-Predovic": 36750, "Loraine-Abshire": 36750, "Nettie-Pollich": 58275, "Ivah-Gorczany": 66150, "Deron-Hartmann": 38850, "Leda-Rodriguez": 47250, "Merritt-Stracke": 47250, "Joseph-Windler": 38850, "Jovani-Howe": 52500, "Kris-Goyette": 10500, "Cheyanne-Halvorson": 36750, "Eduardo-Jacobi": 31500, "Kennedy-Grady": 15750, "Johathan-Kozey": 21000, "Braxton-Metz": 15750, "Susanna-Keebler": 58275, "Gail-Tillman": 36750, "Dorian-Conn": 78750, "Madge-Sipes": 31500, "Otto-Lehner": 15750, "Elian-Bashirian": 47250, "Bernadine-Dibbert": 31500, "Chad-Cassin": 36750, "Alexie-Bernier": 15750, "Ida-Upton": 58275, "Gudrun-Hirthe": 36750, "Werner-Schneider": 58275, "Winston-Hyatt": 66150, "Rhea-Krajcik": 66150, "Jailyn-Kilback": 10500, "Lonnie-Schowalter": 10500, "Kiana-Marquardt": 52500, "Jamarcus-Rohan": 58275, "Elisabeth-Miller": 21000, "Gaylord-Crooks": 78750, "Beth-Wilderman": 47250, "Vinnie-Rempel": 31500, "Samanta-Kemmer": 58275, "Milford-Gleason": 31500, "Gabriella-Ebert": 78750, "Jameson-Grant": 10500, "Lysanne-Eichmann": 66150, "Ted-Bins": 31500, "Salma-Graham": 21000, "Guiseppe-Kub": 15750, "Preston-Schamberger": 38850, "Garrison-Grimes": 52500, "Michelle-OKeefe": 15750, "Queen-Veum": 58275, "Alan-Spencer": 15750, "Ana-Boyer": 78750, "Mable-Marks": 52500, "Jared-Parker": 58275, "Danielle-Hessel": 47250, "Talon-Dietrich": 66150, "Arden-Bayer": 47250, "Narciso-Heller": 58275, "Dejon-Deckow": 47250, "Jeramy-Mosciski": 36750, "Anne-Emard": 38850, "Queen-Stanton": 66150, "Berta-Reichel": 38850, "Esperanza-Rath": 66150, "Eli-Wisozk": 21000, "Freddy-Marks": 31500, "Jimmy-Lockman": 78750, "Arlie-Treutel": 52500, "Malachi-Kunze": 52500, "Chanel-Nienow": 52500, "Malvina-Lind": 58275, "Stuart-Yost": 47250, "Rachael-Farrell": 66150, "Everardo-DuBuque": 52500, "Bethel-Wilderman": 58275, "Nelle-Schmeler": 58275, "Clarabelle-Christiansen": 31500, "Kellen-Langosh": 52500, "Marina-Johnston": 21000, "Kelsi-Donnelly": 58275, "Hadley-Casper": 31500, "Stan-Cremin": 58275, "Maritza-Marquardt": 31500, "Loraine-Daniel": 58275, "Finn-Casper": 10500, "Kaia-Swaniawski": 58275, "Suzanne-Williamson": 10500,
				},
			}

			res := main.DeliveryOrder(testData.Test.Data, testData.Test.Day)

			Expect(res).NotTo(BeNil())
			Expect(len(res)).To(Equal(len(testData.Expected)))

			for k := range testData.Expected {
				Expect(res).To(HaveKey(k))
				Expect(res[k]).To(Equal(testData.Expected[k]))
			}
		})
	})
})
