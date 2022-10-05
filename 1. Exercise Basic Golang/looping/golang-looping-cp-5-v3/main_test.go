package main_test

import (
	main "a21hc3NpZ25tZW50"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("ReverseWord", func() {
	When("the string is empty", func() {
		It("should return empty string", func() {
			Expect(main.ReverseWord("")).To(Equal(""))
		})
	})
	When("input is 'Aku Sayang Ibu'", func() {
		It("should return 'Uka Gnayas Ubi'", func() {
			Expect(main.ReverseWord("Aku Sayang Ibu")).To(Equal("Uka Gnayas Ubi"))
		})
	})

	When("input str contains 'ini terlalu mudah'", func() {
		It("should return 'ini ulalret hadum'", func() {
			Expect(main.ReverseWord("ini terlalu mudah")).To(Equal("ini ulalret hadum"))
		})
	})

	When("input str contains 'KITA AKAN SELALU BERSAMA'", func() {
		It("should return 'Atik Naka Ulales Amasreb'", func() {
			Expect(main.ReverseWord("KITA AKAN SELALU BERSAMA")).To(Equal("Atik Naka Ulales Amasreb"))
		})
	})

	When("input str contains 'A bird fly to Germany and got a worm'", func() {
		It("should return ''", func() {
			Expect(main.ReverseWord("A bird fly to Germany and got a worm")).To(Equal("A drib ylf ot Ynamreg dna tog a mrow"))
		})
	})

	When("input str is letter of lorem ipsum", func() {
		It("should return reverse letter of lorem ipsum", func() {
			var (
				lorem    = "Lorem ipsum dolor sit amet consectetur adipiscing elit sed do eiusmod tempor incididunt ut labore et dolore magna aliqua Volutpat Lacus laoreet non curabitur Amet risus nullam eget felis eget nunc Amet risus nullam eget felis eget Consectetur purus ut faucibus pulvinar elementum Suspendisse ultrices gravida dictum fusce ut placerat orci nulla pellentesque Eu nisl nunc mi ipsum Ac turpis egestas maecenas pharetra convallis Egestas tellus rutrum tellus pellentesque eu tincidunt tortor Nunc sed velit dignissim sodales ut Ut venenatis tellus in metus vulputate eu scelerisque felis Feugiat pretium nibh ipsum consequat nisl vel Porta non pulvinar neque laoreet suspendisse interdum consectetur Sed elementum tempus egestas sed sed risus Blandit volutpat maecenas volutpat blandit aliquam etiam erat Porta non pulvinar neque laoreet suspendisse interdum Vel fringilla est ullamcorper eget nulla facilisi etiam dignissim diam Nisl suscipit adipiscing bibendum est ultricies integer quis auctor Vitae tempus quam pellentesque nec nam aliquam sem Malesuada bibendum arcu vitae elementum curabitur vitae nunc sed Quam elementum pulvinar etiam non quam lacus suspendisse Nunc mattis enim ut tellus elementum sagittis vitae et Ultrices eros in cursus turpis massa tincidunt Morbi tincidunt ornare massa eget egestas purus viverra Massa ultricies mi quis hendrerit dolor magna eget Ut eu sem integer vitae justo eget magna fermentum iaculis Diam sit amet nisl suscipit adipiscing Convallis aenean et tortor at risus viverra Nulla posuere sollicitudin aliquam ultrices sagittis orci a Quam viverra orci sagittis eu volutpat odio facilisis mauris Libero volutpat sed cras ornare arcu dui vivamus Hendrerit dolor magna eget est lorem ipsum dolor sit amet Congue eu consequat ac felis donec et odio Etiam tempor orci eu lobortis elementum nibh tellus At auctor urna nunc id cursus Velit egestas dui id ornare arcu odio ut Blandit massa enim nec dui nunc mattis enim ut tellus Risus quis varius quam quisque Interdum velit laoreet id donec ultrices tincidunt arcu non sodales Vel risus commodo viverra maecenas"
				expected = "Merol muspi rolod tis tema rutetcesnoc gnicsipida tile des od domsuie ropmet tnudidicni tu erobal te erolod angam auqila Taptulov Sucal teeroal non rutibaruc Tema susir mallun tege silef tege cnun Tema susir mallun tege silef tege Rutetcesnoc surup tu subicuaf ranivlup mutnemele Essidnepsus secirtlu adivarg mutcid ecsuf tu tarecalp icro allun euqsetnellep Ue lsin cnun im muspi Ca siprut satsege saneceam arterahp sillavnoc Satsege sullet murtur sullet euqsetnellep ue tnudicnit rotrot Cnun des tilev missingid selados tu Tu sitanenev sullet ni sutem etatupluv ue euqsirelecs silef Taiguef muiterp hbin muspi tauqesnoc lsin lev Atrop non ranivlup euqen teeroal essidnepsus mudretni rutetcesnoc Des mutnemele supmet satsege des des susir Tidnalb taptulov saneceam taptulov tidnalb mauqila maite tare Atrop non ranivlup euqen teeroal essidnepsus mudretni Lev allignirf tse reprocmallu tege allun isilicaf maite missingid maid Lsin tipicsus gnicsipida mudnebib tse seicirtlu regetni siuq rotcua Eativ supmet mauq euqsetnellep cen man mauqila mes Adauselam mudnebib ucra eativ mutnemele rutibaruc eativ cnun des Mauq mutnemele ranivlup maite non mauq sucal essidnepsus Cnun sittam mine tu sullet mutnemele sittigas eativ te Secirtlu sore ni susruc siprut assam tnudicnit Ibrom tnudicnit eranro assam tege satsege surup arreviv Assam seicirtlu im siuq tirerdneh rolod angam tege Tu ue mes regetni eativ otsuj tege angam mutnemref silucai Maid tis tema lsin tipicsus gnicsipida Sillavnoc naenea te rotrot ta susir arreviv Allun ereusop niduticillos mauqila secirtlu sittigas icro a Mauq arreviv icro sittigas ue taptulov oido sisilicaf siruam Orebil taptulov des sarc eranro ucra iud sumaviv Tirerdneh rolod angam tege tse merol muspi rolod tis tema Eugnoc ue tauqesnoc ca silef cenod te oido Maite ropmet icro ue sitrobol mutnemele hbin sullet Ta rotcua anru cnun di susruc Tilev satsege iud di eranro ucra oido tu Tidnalb assam mine cen iud cnun sittam mine tu sullet Susir siuq suirav mauq euqsiuq Mudretni tilev teeroal di cenod secirtlu tnudicnit ucra non selados Lev susir odommoc arreviv saneceam"
			)
			Expect(main.ReverseWord(lorem)).To(Equal(expected))

		})
	})

})
