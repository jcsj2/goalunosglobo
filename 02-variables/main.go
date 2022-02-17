package main

import "fmt"

func main() {

	var name = "Ze Claudio"
	const idade = 30
	var firstName, lastName = "Jose Claudio", "da Silva Junior"
	hobby := "Jogar video game"
	var favoriteGame string = "The king of fighters"

	//https://pkg.go.dev/fmt#hdr-Printing
	fmt.Println("Olá, eu sou", name, "e tenho", idade, "anos")
	fmt.Printf("Meu nome completo é: %s %s \n", firstName, lastName)
	fmt.Printf("Meu hobby é: %v \n", hobby)
	fmt.Println("Meu game favorito é:", favoriteGame)

	//variable types
	/*
		bool, string,
		int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64,

		uintptr, --> unsafe.Pointer
		rune, --> Rune literals are just 32-bit integer values

		byte,
		float32, float64,
		complex64, complex128 --> Complex Numbers contain two parts - real numbers and imaginary numbers, for example, a complex number is 12.8i.
	*/

}
