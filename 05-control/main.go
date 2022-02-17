package main

import "fmt"

func main() {

	points := 70

	if points >= 80 {
		fmt.Println("Prova top")
	} else if points >= 60 { //operators || and && points <= 70
		fmt.Println("Ta safe")
	} else {
		fmt.Println("Ixi")
	}

	var dbzChars []string = []string{"goku"}
	dbzChars = append(dbzChars, "bulma")
	dbzChars = append(dbzChars, "trunks")
	dbzChars = append(dbzChars, "pan")
	dbzChars = append(dbzChars, "vegeta")
	dbzChars = append(dbzChars, "gohan")

	fmt.Println("Temos o goku?", contains(dbzChars, "goku"))

	fruta := "uva"

	switch fruta {
	case "uva":
		fmt.Println("Gosto muito")
	case "kiwi":
		fmt.Println("Gosto nem")
	default:
		fmt.Println("-")
	}

}

func contains(elements []string, value string) bool {
	for _, e := range elements {
		if e == value {
			return true
		}
	}
	return false
}
