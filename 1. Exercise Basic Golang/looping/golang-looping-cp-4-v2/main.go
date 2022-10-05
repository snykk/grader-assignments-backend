package main

import "fmt"

func EmailInfo(email string) string {
	var selectingDomain, selectingTLD bool
	var domain, tld string

	for _, char := range email {
		if fmt.Sprintf("%c", char) == "@" && selectingDomain == false {
			selectingDomain = true
			continue
		}

		if fmt.Sprintf("%c", char) == "." && selectingDomain == true && selectingTLD == false {
			selectingTLD = true
			selectingDomain = false
			continue
		}

		if selectingDomain {
			domain += fmt.Sprintf("%c", char)
		}

		if selectingTLD {
			tld += fmt.Sprintf("%c", char)
		}
	}

	return fmt.Sprintf("Domain: %s dan TLD: %s", domain, tld)

}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(EmailInfo("admin@yahoo.co.id"))
}
