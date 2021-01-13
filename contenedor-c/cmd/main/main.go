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

	fmt.Println("[contenedor-c]: Trying to connect to MongoDB at:", url)
	repository, err := mongo.NewConnection(url, username, password)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("[contenedor-c]: Connection established to MongoDB at:", url)

	songs := http.WebScrapper{URL: "https://vocadb.net/api/songs"}
	fetchedSongs, err := songs.FetchSongs()
	if err != nil {
		log.Fatal(err)
	}
	err = repository.SaveSongs(fetchedSongs)
	if err != nil {
		log.Fatal(err)
	}

	artists := http.WebScrapper{URL: "https://vocadb.net/api/artists"}
	fetchedArtists, err := artists.FetchArtists()
	if err != nil {
		log.Fatal(err)
	}
	err = repository.SaveArtists(fetchedArtists)
	if err != nil {
		log.Fatal(err)
	}

}
