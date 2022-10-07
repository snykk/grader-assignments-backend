package main_test

import (
	main "a21hc3NpZ25tZW50"
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

const PATH_TEST_CASE = "report_test.json"

var BeforeEachHandler = func() {
	if _, err := os.Stat(PATH_TEST_CASE); os.IsNotExist(err) {
		var f *os.File
		if f, err = os.Create(PATH_TEST_CASE); err != nil {
			panic(err)
		}
		defer f.Close()
	}
}

var AfterEachHandler = func() {
	os.Remove(PATH_TEST_CASE)
}

var InsertTestCase = func(data []byte) {
	if err := os.WriteFile(PATH_TEST_CASE, data, 0644); err != nil {
		panic(err)
	}
}

var _ = Describe("ReadJSON", func() {
	BeforeEach(BeforeEachHandler)

	When("studies data in JSON is empty", func() {
		It("should return struct Report with empty Studies data", func() {
			InsertTestCase([]byte(`{"id":"M-20220606","name":"Imam Jaya Permana","date":"12/08/2022","semester":4,"studies":[]}`))

			report, err := main.ReadJSON(PATH_TEST_CASE)
			Expect(err).To(BeNil())
			Expect(report).To(Equal(main.Report{
				Id:       "M-20220606",
				Name:     "Imam Jaya Permana",
				Date:     "12/08/2022",
				Semester: 4,
				Studies:  []main.Study{},
			}))
		})
	})
	When("all grade studies data JSON is not empty", func() {
		It("should return struct Report with not empty grader Studies data", func() {
			InsertTestCase([]byte(`{"id":"M-20220606","name":"Imam Jaya Permana","date":"12/08/2022","semester":4,"studies":[{"study_name":"A","study_credit":3,"grade":"A"},{"study_name":"B","study_credit":3,"grade":"A"},{"study_name":"C","study_credit":3,"grade":"A"}]}`))

			report, err := main.ReadJSON(PATH_TEST_CASE)
			Expect(err).To(BeNil())
			Expect(report).To(Equal(main.Report{
				Id:       "M-20220606",
				Name:     "Imam Jaya Permana",
				Date:     "12/08/2022",
				Semester: 4,
				Studies: []main.Study{
					{"A", 3, "A"},
					{"B", 3, "A"},
					{"C", 3, "A"},
				},
			}))
		})
	})

	AfterEach(AfterEachHandler)

})

var _ = Describe("GradePoint", func() {
	When("the grade studies data is empty", func() {
		It("should return Grader point average is 0", func() {
			reportData := main.Report{
				Id:       "M-20220606",
				Name:     "Imam Jaya Permana",
				Date:     "12/08/2022",
				Semester: 4,
				Studies:  []main.Study{},
			}

			Expect(main.GradePoint(reportData)).To(Equal(0.0))
		})
	})

	When("all the grade studies is A", func() {
		It("should return Grader point average is 4", func() {
			reportData := main.Report{
				Id:       "M-20220606",
				Name:     "Imam Jaya Permana",
				Date:     "12/08/2022",
				Semester: 4,
				Studies: []main.Study{
					{"Pemrograman Berorientasi Objek", 3, "A"},
					{"Pemrograman Web", 3, "A"},
					{"Pemrograman Mobile", 3, "A"},
					{"Pemrograman Desktop", 3, "A"},
				}}

			Expect(main.GradePoint(reportData)).To(Equal(4.0))
		})
	})

	When("all the grade studies is AB", func() {
		It("should return Grader point average is 3.5", func() {
			reportData := main.Report{
				Id: "M-20220606", Name: "Imam Jaya Permana", Date: "12/08/2022", Semester: 4,
				Studies: []main.Study{
					{"Pemrograman Berorientasi Objek", 3, "AB"}, {"Pemrograman Web", 3, "AB"}, {"Pemrograman Mobile", 2, "AB"}, {"Pemrograman Desktop", 2, "AB"},
				}}

			Expect(main.GradePoint(reportData)).To(Equal(3.5))
		})
	})
	When("all the grade studies is B", func() {
		It("should return Grader point average is 3", func() {
			reportData := main.Report{
				Id: "M-20220606", Name: "Imam Jaya Permana", Date: "12/08/2022", Semester: 4,
				Studies: []main.Study{
					{"Pemrograman Berorientasi Objek", 3, "B"}, {"Pemrograman Web", 3, "B"}, {"Pemrograman Mobile", 2, "B"}, {"Pemrograman Desktop", 2, "B"},
				}}

			Expect(main.GradePoint(reportData)).To(Equal(3.0))
		})
	})
	When("all the grade studies is BC", func() {
		It("should return Grader point average is 2.5", func() {
			reportData := main.Report{
				Id: "M-20220606", Name: "Imam Jaya Permana", Date: "12/08/2022", Semester: 4,
				Studies: []main.Study{
					{"Pemrograman Berorientasi Objek", 3, "BC"}, {"Pemrograman Web", 3, "BC"}, {"Pemrograman Mobile", 2, "BC"}, {"Pemrograman Desktop", 2, "BC"},
				}}

			Expect(main.GradePoint(reportData)).To(Equal(2.5))
		})
	})
	When("all the grade studies is C", func() {
		It("should return Grader point average is 2", func() {
			reportData := main.Report{
				Id: "M-20220606", Name: "Imam Jaya Permana", Date: "12/08/2022", Semester: 4,
				Studies: []main.Study{
					{"Pemrograman Berorientasi Objek", 3, "C"}, {"Pemrograman Web", 3, "C"}, {"Pemrograman Mobile", 2, "C"}, {"Pemrograman Desktop", 2, "C"},
				}}

			Expect(main.GradePoint(reportData)).To(Equal(2.0))
		})
	})
	When("all the grade studies is CD", func() {
		It("should return Grader point average is 1.5", func() {
			reportData := main.Report{
				Id: "M-20220606", Name: "Imam Jaya Permana", Date: "12/08/2022", Semester: 4,
				Studies: []main.Study{
					{"Pemrograman Berorientasi Objek", 3, "CD"}, {"Pemrograman Web", 3, "CD"}, {"Pemrograman Mobile", 2, "CD"}, {"Pemrograman Desktop", 2, "CD"},
				}}

			Expect(main.GradePoint(reportData)).To(Equal(1.5))
		})
	})
	When("all the grade studies is D", func() {
		It("should return Grader point average is 1", func() {
			reportData := main.Report{
				Id: "M-20220606", Name: "Imam Jaya Permana", Date: "12/08/2022", Semester: 4,
				Studies: []main.Study{
					{"Pemrograman Berorientasi Objek", 3, "D"}, {"Pemrograman Web", 3, "D"}, {"Pemrograman Mobile", 2, "D"}, {"Pemrograman Desktop", 2, "D"},
				}}

			Expect(main.GradePoint(reportData)).To(Equal(1.0))
		})
	})
	When("all the grade studies is DE", func() {
		It("should return Grader point average is 0.5", func() {
			reportData := main.Report{
				Id: "M-20220606", Name: "Imam Jaya Permana", Date: "12/08/2022", Semester: 4,
				Studies: []main.Study{
					{"Pemrograman Berorientasi Objek", 3, "DE"}, {"Pemrograman Web", 3, "DE"}, {"Pemrograman Mobile", 2, "DE"}, {"Pemrograman Desktop", 2, "DE"},
				}}

			Expect(main.GradePoint(reportData)).To(Equal(0.5))
		})
	})
	When("all the grade studies is E", func() {
		It("should return Grader point average is 0", func() {
			reportData := main.Report{
				Id: "M-20220606", Name: "Imam Jaya Permana", Date: "12/08/2022", Semester: 4,
				Studies: []main.Study{
					{"Pemrograman Berorientasi Objek", 3, "E"}, {"Pemrograman Web", 3, "E"}, {"Pemrograman Mobile", 2, "E"}, {"Pemrograman Desktop", 2, "E"},
				}}

			Expect(main.GradePoint(reportData)).To(Equal(0.0))
		})
	})
})
