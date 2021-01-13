package mongo

import "gopkg.in/mgo.v2/bson"

// Artists struct.
type Artists struct {
	Items []Artist `json:"items" bson:"items"`
}

// Artist struct.
type Artist struct {
	ID                  bson.ObjectId `json:"_id" bson:"_id"`
	ArtistType          string        `json:"artistType" bson:"artist_type"`
	CreateDate          string        `json:"createDate" bson:"create_date"`
	DefaultName         string        `json:"defaultName" bson:"default_name"`
	DefaultNameLanguage string        `json:"defaultNameLanguage" bson:"default_name_language"`
	Name                string        `json:"name" bson:"name"`
	PictureMime         string        `json:"pictureMime" bson:"picture_mime"`
	Status              string        `json:"status" bson:"status"`
	Version             int           `json:"version" bson:"version"`
}

// FetchArtists method.
func (r *Repository) FetchArtists() (*[]Artist, error) {
	artists := &[]Artist{}
	connection := r.Session.DB("animgrn").C("songs")
	err := connection.Find(nil).All(artists)
	if err != nil {
		return nil, err
	}
	return artists, nil
}
