package main

import "fmt"

type Funcionario struct {
	Nome    string
	Salario float64
	Idade   int
}

func aumentarSalario(f Funcionario) {
	fmt.Println(f.Salario + (f.Salario * 0.2))
}

func diminuirSalario(f Funcionario) {
	fmt.Println((f.Salario) - f.Salario*0.2)
}

func tempoServico(f Funcionario) {
	fmt.Println(f.Idade - 18)
}

func imprimeFuncionario(f Funcionario) {
	fmt.Println(f)
}

func main() {
	f := Funcionario{
		Nome:    "Fulano",
		Salario: 1212,
		Idade:   49,
	}
	imprimeFuncionario(f)
	aumentarSalario(f)
	diminuirSalario(f)
	tempoServico(f)
}
