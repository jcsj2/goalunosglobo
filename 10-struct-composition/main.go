package main

import (
	"fmt"
	"time"
)

type baseEntity struct {
	id        int
	createdAt time.Time
	version   int
}

type cliente struct {
	baseEntity
	nome  string
	email string
	endereco
}

type clientePessoaFisica struct {
	cliente
	cpf string
}

type clientePessoaJuridica struct {
	cliente
	cnpj string
}

type endereco struct {
	rua    string
	numero uint
}

func main() {

	jose := cliente{
		baseEntity: baseEntity{
			id: 1,
		},
		nome:  "Optimus Prime",
		email: "Optimus@email.com",
		endereco: endereco{
			rua:    "Dark side of the moon",
			numero: 789,
		},
	}

	fmt.Println(jose)
}
