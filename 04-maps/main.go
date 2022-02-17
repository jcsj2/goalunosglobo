package main

import "fmt"

func main() {
	languageLevel := make(map[string]int)

	languageLevel["Java"] = 10
	languageLevel["PHP"] = 5
	languageLevel["Javascript"] = 9
	languageLevel["GoLang"] = 7
	languageLevel["Dart"] = 3

	//the keys are in asc order
	fmt.Println(languageLevel)

	languageLevel["Dart"] = 8

	fmt.Println(languageLevel)

	delete(languageLevel, "Dart")

	fmt.Println(languageLevel)

	fmt.Println("Quantidade", len(languageLevel))

	// -- //

	consoleAndGames := make(map[string][]string)

	consoleAndGames["Super Nintendo"] = append(consoleAndGames["Super Nintendo"], "Donkey Kong")

	var superNintendo = consoleAndGames["Super Nintendo"]

	consoleAndGames["Super Nintendo"] = append(superNintendo, "Super mario world")
	consoleAndGames["Super Nintendo"] = append(superNintendo, "Star Fox")

	consoleAndGames["Mega drive"] = append(consoleAndGames["Mega drive"], "Sonic")
	consoleAndGames["Mega drive"] = append(consoleAndGames["Mega drive"], "Streets of rage")
	consoleAndGames["Mega drive"] = append(consoleAndGames["Mega drive"], "Shinobi")

	fmt.Println(consoleAndGames)
	fmt.Println(consoleAndGames["Super Nintendo"])
	fmt.Println(consoleAndGames["Super Nintendo"][1])

	// delete
	consoleAndGames["Super Nintendo"] = append(consoleAndGames["Super Nintendo"][:1], consoleAndGames["Super Nintendo"][2:3]...)

	fmt.Println(consoleAndGames["Super Nintendo"])
}
