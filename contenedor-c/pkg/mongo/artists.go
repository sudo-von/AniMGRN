package mongo

import (
	mgo "gopkg.in/mgo.v2"
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
func SaveArtists(session *mgo.Session, artists *Artists) error {
	connection := session.DB("animgrn").C("artists")
	for _, song := range artists.Items {
		err := connection.Insert(song)
		if err != nil {
			return err
		}
	}
	return nil
}
