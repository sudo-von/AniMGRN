package mongo

import "gopkg.in/mgo.v2/bson"

// Playlist struct.
type Playlist struct {
	Items []Song `json:"items" bson:"items"`
}

// Song struct.
type Song struct {
	ID                  bson.ObjectId `bson:"_id"`
	ArtistString        string        `json:"artistString" bson:"artist_tring"`
	CreateDate          string        `json:"createDate" bson:"create_date"`
	DefaultName         string        `json:"defaultName" bson:"default_name"`
	DefaultNameLanguage string        `json:"defaultNameLanguage" bson:"default_name_language"`
	FavoritedTimes      int           `json:"favoritedTimes" bson:"favorited_times"`
	LengthSeconds       int           `json:"lengthSeconds" bson:"length_seconds"`
	Name                string        `json:"name" bson:"name"`
	PVServices          string        `json:"pvServices" bson:"pv_services"`
	PublishDate         string        `json:"publishDate" bson:"publish_date"`
	RatingScore         int           `json:"ratingScore" bson:"rating_score"`
	SongType            string        `json:"songType" bson:"song_type"`
	Status              string        `json:"status" bson:"status"`
	Version             int           `json:"version" bson:"version"`
}

// FetchSongs method.
func (r *Repository) FetchSongs() (*[]Song, error) {
	songs := &[]Song{}
	connection := r.Session.DB("animgrn").C("songs")
	err := connection.Find(nil).All(songs)
	if err != nil {
		return nil, err
	}
	return songs, nil
}
