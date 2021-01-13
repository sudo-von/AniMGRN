package server

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// Server interface.
type Server interface {
	Router() http.Handler
}

type api struct {
	router http.Handler
}

func New() Server {
	api := &api{}
	r := mux.NewRouter()
	r.HandleFunc("/songs", api.fetchSongs).Methods(http.MethodGet)
	r.HandleFunc("/artists", api.fetchArtists).Methods(http.MethodGet)
	api.router = r
	return api
}

func (a *api) fetchSongs(w http.ResponseWriter, r *http.Request) {
	gophers, _ := a.repository.FetchGophers()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(gophers)
}

func (a *api) fetchArtists(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	gopher, err := a.repository.FetchGopherByID(vars["ID"])
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusNotFound) // We use not found for simplicity
		json.NewEncoder(w).Encode("Gopher Not found")
		return
	}
	json.NewEncoder(w).Encode(gopher)
}

func (a *api) Router() http.Handler {
	return a.router
}
