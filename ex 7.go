package main

import (
	"fmt"
	"strings"
)

type Animal struct {
	Nome    string
	Especie string
	Idade   int
	Som     string
}

func modificarSom(a Animal) {
	a.Som = strings.ReplaceAll(a.Som, a.Som, "Miau")
	fmt.Println(a.Som)
}

func imprimeAnimal(a Animal) {
	fmt.Println(a)
}

func main() {
	a := Animal{
		Nome:    "Cachorro",
		Especie: "Canis familiaris",
		Idade:   3,
		Som:     "AuAu",
	}
	imprimeAnimal(a)
	modificarSom(a)
}
