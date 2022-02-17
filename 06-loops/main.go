package main

import "fmt"

func main() {

	favoriteFoods := [3]string{"Lasanha", "Pizza", "Carbonara"}

	fmt.Println("Minhas comidas favoritas são")

	for index, food := range favoriteFoods {
		fmt.Println(index, "-", food)
	}

	for i := 0; i < len(favoriteFoods); i++ {
		fmt.Println(i, "-", favoriteFoods[i])
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}

		fmt.Println(n)
	}

	//while
	count := 1
	for count <= 5 {
		fmt.Println("O contador é", count)
		count++
	}

	// FOR "infinito"
	// for {

	// }

}
