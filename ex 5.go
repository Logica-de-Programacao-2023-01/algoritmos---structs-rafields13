package main

import (
	"fmt"
	"time"
)

type Musica struct {
	titulo  string
	artista string
	duracao time.Duration
}

type Playlist struct {
	nome    string
	musicas []Musica
}

func playlistsComMusica(titulo string, playlists []Playlist) []Playlist {
	var playlistsComMusica []Playlist
	for _, playlist := range playlists {
		for _, musica := range playlist.musicas {
			if musica.titulo == titulo {
				playlistsComMusica = append(playlistsComMusica, playlist)
				break
			}
		}
	}
	return playlistsComMusica
}

func main() {
	musica1 := Musica{"Stairway to Heaven", "Led Zeppelin", 8*time.Minute + 2*time.Second}
	musica2 := Musica{"Bohemian Rhapsody", "Queen", 6*time.Minute + 7*time.Second}
	musica3 := Musica{"Imagine", "John Lennon", 3*time.Minute + 3*time.Second}
	playlist1 := Playlist{"Minha Playlist 1", []Musica{musica1, musica2}}
	playlist2 := Playlist{"Minha Playlist 2", []Musica{musica3, musica2}}
	playlist3 := Playlist{"Minha Playlist 3", []Musica{musica1}}
	playlists := []Playlist{playlist1, playlist2, playlist3}
	playlistsEncontradas := playlistsComMusica("Bohemian Rhapsody", playlists)
	fmt.Println(playlistsEncontradas)
}
