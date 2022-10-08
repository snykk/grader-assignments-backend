package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

type LoanData struct {
	StartBalance int
	Data         []Loan
	Employees    []Employee
}

type Loan struct {
	Date        string
	EmployeeIDs []string
}

type Employee struct {
	ID       string
	Name     string
	Position string
}

// json structure
type LoanRecord struct {
	MonthDate    string     `json:"month_date"`
	StartBalance int        `json:"start_balance"`
	EndBalance   int        `json:"end_balance"`
	Borrowers    []Borrower `json:"borrowers"`
}

type Borrower struct {
	Id        string `json:"id"`
	TotalLoan int    `json:"total_loan"`
}

func LoanReport(data LoanData) LoanRecord {
	// init end balance value with start balance
	var finalEndBalance int = data.StartBalance
	var finalBorrower = []Borrower{}
	var isGetOldBorrower bool
	var dateList []string

	// to iterate each "Data" in LoanData
	for _, dataItem := range data.Data {
		dateList = strings.Split(dataItem.Date, "-")

		// to iterate each Employee Id in "Data"
		for _, employeeIdItem := range dataItem.EmployeeIDs {
			isGetOldBorrower = false

			if finalEndBalance <= 0 {
				break
			}

			// to iterate each item in []finalBorrower
			for index, finalBorrowerItem := range finalBorrower {
				if employeeIdItem == finalBorrowerItem.Id {
					finalEndBalance -= 50000
					finalBorrower[index].TotalLoan += 50000
					isGetOldBorrower = true
					break
				}
			}

			if isGetOldBorrower == false {
				newBorrower := Borrower{
					Id:        employeeIdItem,
					TotalLoan: 50000,
				}

				finalEndBalance -= 50000
				finalBorrower = append(finalBorrower, newBorrower)
			}
		}
	}

	newLoanRecord := LoanRecord{
		MonthDate:    fmt.Sprintf("%s %s", dateList[1], dateList[2]),
		StartBalance: data.StartBalance,
		EndBalance:   finalEndBalance,
		Borrowers:    finalBorrower,
	}

	return newLoanRecord // TODO: replace this
}

func RecordJSON(record LoanRecord, path string) error {
	// encode dari struct ke JSON
	jsonData, err := json.Marshal(record)
	if err != nil {
		return err
	}

	// write file JSON ke file 'user.json'
	err = ioutil.WriteFile(path, jsonData, 0644)
	if err != nil {
		return err
	}

	return nil // TODO: replace this
}

// gunakan untuk debug
func main() {
	records := LoanReport(LoanData{
		StartBalance: 500000,
		Data: []Loan{
			{"01-January-2021", []string{"1", "2"}},
			{"02-January-2021", []string{"1", "2", "3"}},
			{"03-January-2021", []string{"2", "3"}},
			{"04-January-2021", []string{"1", "3"}},
		},
		Employees: []Employee{
			{"1", "Employee A", "Manager"},
			{"2", "Employee B", "Staff"},
			{"3", "Employee C", "Staff"},
		},
	})

	err := RecordJSON(records, "loan-records.json")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(records)
}
