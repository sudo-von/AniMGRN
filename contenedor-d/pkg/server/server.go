package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/AniMGRN/contenedor-d/pkg/mongo"
	"github.com/gorilla/mux"
)

// Server interface.
type Server interface {
	Router() http.Handler
}

type api struct {
	router     http.Handler
	repository *mongo.Repository
}

func New(repo *mongo.Repository) Server {
	a := &api{repository: repo}
	router := mux.NewRouter()
	router.HandleFunc("/songs", a.fetchSongs).Methods(http.MethodGet)
	router.HandleFunc("/artists", a.fetchArtists).Methods(http.MethodGet)
	a.router = router
	return a
}

func (a *api) fetchSongs(w http.ResponseWriter, r *http.Request) {
	data, err := a.repository.FetchSongs()
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Internal error")
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func (a *api) fetchArtists(w http.ResponseWriter, r *http.Request) {
	data, err := a.repository.FetchArtists()
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Internal error")
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func (a *api) Router() http.Handler {
	return a.router
}
