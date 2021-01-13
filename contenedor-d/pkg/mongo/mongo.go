package mongo

import (
	"time"

	mgo "gopkg.in/mgo.v2"
)

// Interface struct.
type IRepository interface {
	FetchArtists() error
	FetchSongs() error
}

// Repository struct.
type Repository struct {
	Session *mgo.Session
}

// NewConnection establishes a new connection to the cluster.
func NewConnection(url, username, password string) (*Repository, error) {

	info := &mgo.DialInfo{
		Addrs:    []string{url},
		Timeout:  60 * time.Second,
		Username: username,
		Password: password,
	}

	session, err := mgo.DialWithInfo(info)
	if err != nil {
		return nil, err
	}

	repository := &Repository{Session: session}
	return repository, nil
}
