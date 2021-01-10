package http

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/AniMGRN/contenedor-c/pkg/mongo"
)

func FetchSongs() (*mongo.Playlist, error) {
	playlist := &mongo.Playlist{}
	response, err := http.Get("https://vocadb.net/api/songs")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	jsonErr := json.Unmarshal(body, &playlist)
	if jsonErr != nil {
		log.Fatal(jsonErr)
		return nil, jsonErr
	}
	return playlist, nil
}

func FetchArtists() (*mongo.Artists, error) {
	artist := &mongo.Artists{}
	response, err := http.Get("https://vocadb.net/api/artists")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	jsonErr := json.Unmarshal(body, &artist)
	if jsonErr != nil {
		log.Fatal(jsonErr)
		return nil, jsonErr
	}
	return artist, nil
}
