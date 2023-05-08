package main

import "fmt"

type Aluno struct {
	Nome  string
	Idade int
	Notas []float64
}

func addNotas(a Aluno) {
	a.Notas = append(a.Notas, 10)
	fmt.Println(a.Notas)
}

func removeNotas(a Aluno) {
	a.Notas = append(a.Notas[:1], a.Notas[2:]...)
	fmt.Println(a.Notas)
}

func media(a Aluno) {
	sum := 0.0
	for i := 0; i < len(a.Notas); i++ {
		sum += a.Notas[i]
	}
	fmt.Println(sum / float64(len(a.Notas)))
}

func imprimeAluno(a Aluno) {
	fmt.Println(a.Nome)
	fmt.Println(a.Idade)
	fmt.Println(a.Notas)
}

func main() {
	a := Aluno{
		Nome:  "Joãozinho",
		Idade: 12,
		Notas: []float64{9, 10, 8, 7},
	}
	imprimeAluno(a)
	addNotas(a)
	removeNotas(a)
	media(a)
}
