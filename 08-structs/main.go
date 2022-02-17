package main

import "fmt"

type videoGame struct {
	name  string
	brand string
	year  uint
}

func (v videoGame) printInfo() {
	fmt.Println("Nome:", v.name, "Marca:", v.brand, "Ano:", v.year)
}

func NewVideoGame(name string, brand string) videoGame {
	return videoGame{
		name:  name,
		brand: brand,
		year:  2983989834,
	}
}

func main() {

	playstation := videoGame{
		name:  "Play One",
		brand: "Sony",
		year:  2000,
	}

	superNintendo := videoGame{
		name:  "Super Nintendo",
		brand: "Nintendo",
	}

	superNintendo.year = 1990

	fmt.Println(playstation)
	fmt.Println(playstation.name)

	fmt.Println(superNintendo)

	superNintendo.printInfo()

}
