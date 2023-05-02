package main

import "fmt"

type Musicas []struct {
	Titulo  string
	Artista string
	Duracao float64
}

type Playlist struct {
	Nome    string
	Musicas Musicas
}

func info(p Playlist) {
	fmt.Println(p.Nome)
	fmt.Println(p.Musicas)
}

func main() {
	p := Playlist{Nome: "Samba", Musicas: Musicas []struct{{Titulo: "Tá Escrito", "Nosso Amor Quer Paz", Artista: "Grupo Revelação", "Péricles e Marvvila", Duracao: 233, 174}}}
	info(p)
}
