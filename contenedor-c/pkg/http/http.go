package http

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/AniMGRN/contenedor-c/pkg/mongo"
)

// ScrapperRepository interface.
type ScrapperRepository interface {
	FetchSongs() (*mongo.Playlist, error)
	FetchArtists() (*mongo.Artists, error)
}

// WebScrapper struct.
type WebScrapper struct {
	URL string
}

// FetchSongs method.
func (webScrapper *WebScrapper) FetchSongs() (*mongo.Playlist, error) {
	playlist := &mongo.Playlist{}
	response, err := http.Get(webScrapper.URL)
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
	fmt.Println("[contenedor-c]: Fetched", len(playlist.Items), "songs 🎵")
	return playlist, nil
}

// FetchArtists method.
func (webScrapper *WebScrapper) FetchArtists() (*mongo.Artists, error) {
	artist := &mongo.Artists{}
	response, err := http.Get(webScrapper.URL)
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
	fmt.Println("[contenedor-c]: Fetched", len(artist.Items), "artists 🎤")
	return artist, nil
}
