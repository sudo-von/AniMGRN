package mongo

import (
	"time"

	mgo "gopkg.in/mgo.v2"
)

// NewConnection establishes a new connection to the cluster.
func NewConnection(url, username, password string) (*mgo.Session, error) {

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

	return session, nil
}
