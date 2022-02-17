package main

import "fmt"

type game struct {
	name string
	year uint
}

func (g *game) changeName(newName string) {
	g.name = newName
	fmt.Println("Setei novo valor", g)
}

func setSkill(newSkill *string) {
	fmt.Println("Skill recebida", newSkill)
	fmt.Println(newSkill)
	*newSkill = "Super força"
}

func main() {

	mario := game{
		name: "Super Mario",
		year: 1990,
	}

	mario.changeName("Super mario world")

	fmt.Println("Pós", mario)

	skill := "Super pulo"

	fmt.Println(&skill)

	setSkill(&skill)

	fmt.Println("Nova skill", skill)

}
