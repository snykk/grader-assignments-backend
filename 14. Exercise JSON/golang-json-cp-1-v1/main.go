package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Study struct {
	StudyName   string `json:"study_name"`
	StudyCredit int    `json:"study_credit"`
	Grade       string `json:"grade"`
}

type Report struct {
	Id       string  `json:"id"`
	Name     string  `json:"name"`
	Date     string  `json:"date"`
	Semester int     `json:"semester"`
	Studies  []Study `json:"studies"`
}

func scaleHelper(grade string) float64 {
	mapHelper := map[string]float64{
		"A":  4,
		"AB": 3.5,
		"B":  3,
		"BC": 2.5,
		"C":  2,
		"CD": 1.5,
		"D":  1,
		"DE": 0.5,
		"E":  0,
	}

	return mapHelper[grade]
}

// gunakan fungsi ini untuk mengambil data dari file json
// kembalian berupa struct 'Report' dan error
func ReadJSON(filename string) (Report, error) {
	jsonData, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	var report Report
	err = json.Unmarshal(jsonData, &report)
	if err != nil {
		panic(err)
	}

	return report, nil
}

func GradePoint(report Report) float64 {
	if len(report.Studies) == 0 {
		return 0
	}

	var totalGrade float64
	var totalCredit int

	for _, detailStudy := range report.Studies {
		totalGrade += float64(detailStudy.StudyCredit) * scaleHelper(detailStudy.Grade)
		totalCredit += detailStudy.StudyCredit
	}

	return float64(totalGrade) / float64(totalCredit) // TODO: replace this
}

func main() {
	// bisa digunakan untuk menguji test case
	report, err := ReadJSON("report.json")
	if err != nil {
		panic(err)
	}

	gradePoint := GradePoint(report)
	fmt.Println(gradePoint)
}
