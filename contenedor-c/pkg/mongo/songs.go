package mongo

import (
	mgo "gopkg.in/mgo.v2"
)

// Playlist struct.
type Playlist struct {
	Items []Song `json:"items" bson:"items"`
}

// Song struct.
type Song struct {
	ArtistString        string `json:"artistString" bson:"artist_tring"`
	CreateDate          string `json:"createDate" bson:"create_date"`
	DefaultName         string `json:"defaultName" bson:"default_name"`
	DefaultNameLanguage string `json:"defaultNameLanguage" bson:"default_name_language"`
	FavoritedTimes      int    `json:"favoritedTimes" bson:"favorited_times"`
	LengthSeconds       int    `json:"lengthSeconds" bson:"length_seconds"`
	Name                string `json:"name" bson:"name"`
	PVServices          string `json:"pvServices" bson:"pv_services"`
	PublishDate         string `json:"publishDate" bson:"publish_date"`
	RatingScore         int    `json:"ratingScore" bson:"rating_score"`
	SongType            string `json:"songType" bson:"song_type"`
	Status              string `json:"status" bson:"status"`
	Version             int    `json:"version" bson:"version"`
}

// SaveSongs stores playlist items in a collection.
func SaveSongs(session *mgo.Session, playlist *Playlist) error {
	connection := session.DB("animgrn").C("songs")
	for _, song := range playlist.Items {
		err := connection.Insert(song)
		if err != nil {
			return err
		}
	}
	return nil
}
