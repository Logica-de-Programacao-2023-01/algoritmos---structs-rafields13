package main

import "fmt"

type Filme struct {
	Titulo     string
	Diretor    string
	Ano        int
	Avaliacoes []int
}

func addAvaliacoes(f Filme) {
	f.Avaliacoes = append(f.Avaliacoes, 10)
	fmt.Println(f.Avaliacoes)
}

func removeAvaliacoes(f Filme) {
	f.Avaliacoes = append(f.Avaliacoes[:1], f.Avaliacoes[2:]...)
	fmt.Println(f.Avaliacoes)
}

func mediaAvaliacoes(f Filme) {
	sum := 0
	for i := 0; i < len(f.Avaliacoes); i++ {
		sum += f.Avaliacoes[i]
	}
	fmt.Println(sum / (len(f.Avaliacoes)))
}

func imprimeFilme(f Filme) {
	fmt.Println(f.Titulo)
	fmt.Println(f.Diretor)
	fmt.Println(f.Ano)
	fmt.Println(f.Avaliacoes)
}

func main() {
	f := Filme{
		Titulo:     "Titanic",
		Diretor:    "James Cameron",
		Ano:        1914,
		Avaliacoes: []int{9, 10, 8, 7},
	}
	imprimeFilme(f)
	addAvaliacoes(f)
	removeAvaliacoes(f)
	mediaAvaliacoes(f)
}
