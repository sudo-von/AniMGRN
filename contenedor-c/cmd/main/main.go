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

	fmt.Println("Trying to connect to MongoDB at:", url)
	session, err := mongo.NewConnection(url, username, password)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection established to MongoDB at:", url)

	songs := http.WebScrapper{URL: "https://vocadb.net/api/songs"}
	fetchedSongs, err := songs.FetchSongs()
	if err != nil {
		log.Fatal(err)
	}
	err = mongo.SaveSongs(session, fetchedSongs)
	if err != nil {
		log.Fatal(err)
	}

	artists := http.WebScrapper{URL: "https://vocadb.net/api/artists"}
	fetchedArtists, err := artists.FetchArtists()
	if err != nil {
		log.Fatal(err)
	}
	err = mongo.SaveArtists(session, fetchedArtists)
	if err != nil {
		log.Fatal(err)
	}

}
