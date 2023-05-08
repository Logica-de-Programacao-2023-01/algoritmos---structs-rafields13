package main

import "fmt"

type Viagem struct {
	Origem  string
	Destino string
	Data    string
	Preco   float64
}

func maisCara(viagens []Viagem) Viagem {
	var viagemMaisCara Viagem
	for i, viagem := range viagens {
		if i == 0 || viagem.Preco > viagemMaisCara.Preco {
			viagemMaisCara = viagem
		}
	}
	return viagemMaisCara
}

func main() {
	viagem1 := Viagem{
		Origem:  "Brasília",
		Destino: "São Paulo",
		Data:    "10/06/2023",
		Preco:   445,
	}
	viagem2 := Viagem{
		Origem:  "São Paulo",
		Destino: "Miami",
		Data:    "11/06/2023",
		Preco:   2404,
	}
	viagens := []Viagem{viagem1, viagem2}
	viagemMaisCara := maisCara(viagens)
	fmt.Println(viagemMaisCara)
}
