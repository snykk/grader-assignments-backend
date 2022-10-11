package main

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

type RowData struct {
	RankWebsite int
	Domain      string
	TLD         string
	IDN_TLD     string
	Valid       bool
	RefIPs      int
}

func ProcessGetTLD(website RowData, ch chan RowData, chErr chan error) {
	// TODO: answer here
	if website.Domain == "" {
		chErr <- errors.New("domain name is empty")
		return
	}

	if !website.Valid {
		chErr <- errors.New("domain not valid")
		return
	}

	if website.RefIPs == -1 {
		chErr <- errors.New("domain RefIPs not valid")
		return
	}

	substractDom := strings.Split(website.Domain, ".")
	var tld, idn_tld string

	switch string(substractDom[len(substractDom)-1]) {
	case "com":
		tld = ".com"
		idn_tld = ".co.id"
	case "gov":
		tld = ".gov"
		idn_tld = ".go.id"
	case "org":
		tld = ".org"
		idn_tld = ".org.id"
	default:
		tld = "." + string(substractDom[len(substractDom)-1])
		idn_tld = "TLD"
	}

	website.TLD = tld
	website.IDN_TLD = idn_tld

	ch <- website
	time.Sleep(100 * time.Millisecond)
}

// Gunakan variable ini sebagai goroutine di fungsi FilterAndGetDomain
var FuncProcessGetTLD = ProcessGetTLD

func FilterAndFillData(TLD string, data []RowData) ([]RowData, error) {
	ch := make(chan RowData, len(data))
	errCh := make(chan error)
	var result = []RowData{}

	for _, website := range data {

		//perbaiki code di bawah ini menggunakan goroutine
		go FuncProcessGetTLD(website, ch, errCh) // TODO: replace this
	}

	for i := 0; i < len(data); i++ {
		// fungsi ini sebagai receive data dari channel 'messages'
		select {
		case msg := <-errCh:
			return []RowData{}, msg
		case msg := <-ch:
			if TLD == msg.TLD {
				result = append(result, msg)
			}
		}

	}

	return result, nil // TODO: replace this
}

// gunakan untuk melakukan debugging
func main() {
	fmt.Println(FilterAndFillData(".xyz", []RowData{
		{
			RankWebsite: 1,
			Domain:      "google.com",
			Valid:       true,
			RefIPs:      1,
			TLD:         "",
			IDN_TLD:     "",
		}, {
			RankWebsite: 2,
			Domain:      "bukanruang.org",
			Valid:       true,
			RefIPs:      1,
			TLD:         "",
			IDN_TLD:     "",
		}, {
			RankWebsite: 3,
			Domain:      "bukanjudi.xyz",
			Valid:       true,
			RefIPs:      1,
			TLD:         "",
			IDN_TLD:     "",
		},
	}))
}
