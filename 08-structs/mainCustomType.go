package main

import "fmt"

type listaTOP []string

func (l listaTOP) contains(value string) bool {
	for _, e := range l {
		if e == value {
			return true
		}
	}
	return false
}

func main() {

	lista := listaTOP{"asd", "dsa"}

	fmt.Println(lista)

	fmt.Println(lista.contains("asd"))

}
