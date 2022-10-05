package main

import (
	"fmt"
	"sort"
	"strings"
)

func FindShortestName(names string) string {
	var separator string
	for index := 0; index < len(names); index++ {
		if fmt.Sprintf("%c", names[index]) == " " {
			separator = " "
			break
		} else if fmt.Sprintf("%c", names[index]) == "," {
			separator = ","
			break
		} else if fmt.Sprintf("%c", names[index]) == ";" {
			separator = ";"
			break
		}
	}

	nameList := strings.Split(names, separator)

	sort.Strings(nameList)
	sort.Slice(nameList, func(i, j int) bool {
		return len(nameList[i]) < len(nameList[j])
	})

	fmt.Println(nameList)
	return nameList[0]
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(FindShortestName("Hanif Joko Tio Andi Budi Caca Hamdan")) // "Tio"
	fmt.Println(FindShortestName("Budi;Tia;Tio"))                         // "Tia"
	fmt.Println(FindShortestName("Ari,Aru,Ara,Andi,Asik"))
}
