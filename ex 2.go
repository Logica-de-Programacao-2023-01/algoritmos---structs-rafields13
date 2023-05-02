package main

import "fmt"

type Endereco struct {
	Rua    string
	Numero int
	Cidade string
	Estado string
}

type Pessoa struct {
	Nome     string
	Idade    int
	Endereco Endereco
}

func info(p Pessoa) {
	fmt.Println(p.Nome)
	fmt.Println(p.Idade)
	fmt.Println(p.Endereco)
}

func main() {
	p := Pessoa{Nome: "Joãozinho", Idade: 12, Endereco: Endereco{Rua: "Joãozinho Guimarães", Numero: 12, Cidade: "Santos", Estado: "São Paulo"}}
	info(p)
}
