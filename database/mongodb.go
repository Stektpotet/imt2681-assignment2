package database

import (
	"crypto/tls"
	"net"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type DBStorage interface {
	Init()
	Add(collection string, data interface{}) error
	Count(collection string) int
	Get(collection string, query bson.M, data interface{}) (found bool)
	Delete(collection string, query bson.M) (found bool)
	GetAll(collection string, data interface{})
}

// MongoDB - metadata of a db collection
type MongoDB struct {
	HostURLs  []string
	AdminUser string
	AdminPass string
	Name      string
}

// CreateLocalSession - create DB Collection session
func (db *MongoDB) CreateLocalSession() *mgo.Session {
	session, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}
	return session
}
func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// CreateSession - create DB Collection session
func (db *MongoDB) CreateSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost")

	if !contains(db.HostURLs, "localhost") {
		dialInfo := &mgo.DialInfo{
			Addrs:    db.HostURLs,
			Username: db.AdminUser,
			Password: db.AdminPass,

			DialServer: func(addr *mgo.ServerAddr) (net.Conn, error) {
				return tls.Dial("tcp", addr.String(), &tls.Config{})
			},
			Timeout: time.Second * 10,
		}
		s, err = mgo.DialWithInfo(dialInfo)
	}

	if err != nil {
		panic(err)
	}
	return s
}
