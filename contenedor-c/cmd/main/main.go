package main

import (
	"fmt"
	"log"

	"github.com/AniMGRN/contenedor-c/pkg/http"
	"github.com/AniMGRN/contenedor-c/pkg/mongo"
)

const (
	url          = "localhost:27017"
	username     = ""
	password     = ""
	databaseName = "anime_db"
)

func main() {

	session, err := mongo.NewConnection(url, username, password)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected established to MongoDB at:", url)

	songs, err := http.FetchSongs()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Fetched", len(songs.Items), "songs 🎵")
	err = mongo.SaveSongs(session, songs)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Stored", len(songs.Items), "songs 🎵")

	artists, err := http.FetchArtists()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Fetched", len(artists.Items), "artists 🎤")
	err = mongo.SaveArtists(session, artists)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Stored", len(artists.Items), "artists 🎤")

}
