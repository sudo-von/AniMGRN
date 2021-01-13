package mongo

import (
	"fmt"
)

// Artists struct.
type Artists struct {
	Items []Artist `json:"items" bson:"items"`
}

// Artist struct.
type Artist struct {
	ArtistType          string `json:"artistType" bson:"artist_type"`
	CreateDate          string `json:"createDate" bson:"create_date"`
	DefaultName         string `json:"defaultName" bson:"default_name"`
	DefaultNameLanguage string `json:"defaultNameLanguage" bson:"default_name_language"`
	Name                string `json:"name" bson:"name"`
	PictureMime         string `json:"pictureMime" bson:"picture_mime"`
	Status              string `json:"status" bson:"status"`
	Version             int    `json:"version" bson:"version"`
}

// SaveArtists stores playlist items in a collection.
func (r *Repository) SaveArtists(artists *Artists) error {
	connection := r.Session.DB("animgrn").C("artists")
	for _, artist := range artists.Items {
		err := connection.Insert(artist)
		if err != nil {
			return err
		}
	}
	fmt.Println("[contenedor-c]: Stored", len(artists.Items), "artists 🎤")
	return nil
}
