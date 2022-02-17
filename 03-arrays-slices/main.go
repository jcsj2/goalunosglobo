package main

import "fmt"

func main() {

	//arrays começam por posição 0, 1, ?
	var colors = [3]string{"Azul", "Vinho"}

	fmt.Println("O array tem", len(colors), "posições")
	fmt.Println("Os valores são", colors)
	fmt.Println("O primeiro valor é:", colors[0])

	colors[2] = "Vermelho"

	fmt.Println("O primeiro valor é:", colors[2])

	//slices
	var favoriteGames []string

	favoriteGames = append(favoriteGames, "The King of fighters") // 0 -> 1
	favoriteGames = append(favoriteGames, "Bloodborne")           // 1 -> 2

	favoriteGames = append(favoriteGames, "Street figher") // 2 -> 3

	favoriteGames = append(favoriteGames, "Super mario world") // 3 -> 4
	favoriteGames = append(favoriteGames, "Sonic")             // 4 -> 5

	fmt.Println(favoriteGames)

	//to "delete", append two slices with ranges
	favoriteGames = append(favoriteGames[0:2], favoriteGames[3:]...) //função variadic ... para 1 ou mais valores

	fmt.Println("Meus jogos", favoriteGames)

}
