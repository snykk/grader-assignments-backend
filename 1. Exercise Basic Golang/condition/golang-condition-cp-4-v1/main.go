package main

import "fmt"

func GetTicketPrice(VIP, regular, student, day int) float32 {
	total := VIP*30 + regular*20 + student*10
	var discount float32

	if total < 100 {
		return float32(total)
	}

	if day%2 == 0 {
		// genap
		if VIP+regular+student < 5 {
			discount = 0.1
		} else {
			discount = 0.2
		}
	} else {
		// ganjil
		if VIP+regular+student < 5 {
			discount = 0.15
		} else {
			discount = 0.25
		}
	}

	return (1 - discount) * float32(total)
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(GetTicketPrice(1, 1, 1, 20))
}
