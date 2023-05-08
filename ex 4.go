package main

import (
	"fmt"
	"time"
)

type Musica struct {
	Titulo  string
	Artista string
	Duracao time.Duration
}

type Playlist struct {
	Nome    string
	Musicas []Musica
}

func (p Playlist) DuracaoTotal() time.Duration {
	total := time.Duration(0)
	for _, m := range p.Musicas {
		total += m.Duracao
	}
	return total
}

func imprimePlaylist(p Playlist) {
	fmt.Println(p.Nome)
	fmt.Println(p.DuracaoTotal())
	for _, m := range p.Musicas {
		fmt.Println(m.Titulo, m.Artista, m.Duracao)
	}
}

func main() {
	playlist := Playlist{
		Nome: "Forró Raiz",
		Musicas: []Musica{
			{Titulo: "Isso é Vaquejada", Artista: "Tarcísio do Acordeon", Duracao: time.Second * 190},
			{Titulo: "Nunca Mais", Artista: "NATTAN e Xand Avião", Duracao: time.Second * 187},
			{Titulo: "Americana na Vaquejada", Artista: "Grandão Vaqueiro e NATTAN", Duracao: time.Second * 156},
		},
	}

	imprimePlaylist(playlist)
}
