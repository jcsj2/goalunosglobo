package main

import (
	"fmt"
)

func main() {

	a := sendMessages("Teste 1", "Teste 2")
	fmt.Println(a)

	msg, ok := checkAge("Junior", 9)

	fmt.Println(msg, ok)

}

func sendMessages(msgs ...string) bool {
	for _, message := range msgs {
		fmt.Println("Enviando mensagem", message)
	}

	return true
}

func checkAge(name string, age uint) (string, bool) {
	if age >= 18 {
		return fmt.Sprint(name, " permitido"), true
	}

	return fmt.Sprint(name, " não permitido"), false
}
