package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/AniMGRN/contenedor-d/pkg/mongo"
	"github.com/AniMGRN/contenedor-d/pkg/server"
)

const (
	url          = "localhost:27017"
	username     = ""
	password     = ""
	databaseName = "anime_db"
)

func main() {

	fmt.Println("[contenedor-d]: Trying to connect to MongoDB at:", url)
	session, err := mongo.NewConnection(url, username, password)
	if err != nil {
		log.Fatal("[contenedor-d]:", err)
	}
	fmt.Println("[contenedor-d]: Connection established to MongoDB at:", url)

	s := server.New(session)
	log.Fatal(http.ListenAndServe(":8080", s.Router()))
}
