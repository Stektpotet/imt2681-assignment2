package database

import (
	"crypto/tls"
	"net"
	"time"

	"github.com/stektpotet/imt2681-assignment2/fixer"
	mgo "gopkg.in/mgo.v2"
)

type DBStorage interface {
	Init()
	Add(fixer.Currency) error
	Count() int
	Get(key string) (fixer.Currency, bool)
	GetAll() []fixer.Currency
}

// MongoDB - metadata of a db collection
type MongoDB struct {
	HostURLs       []string
	AdminUser      string
	AdminPass      string
	Name           string
	CollectionName string
}

// CreateSession - create DB session
func (db *MongoDB) CreateSession() *mgo.Session {
	// mongodbURL := "mongodb://" + mongodbUser + ":" + mongodbPass + "@,/test?replicaSet=Cluster0-shard-0&authSource=admin"

	dialInfo := &mgo.DialInfo{
		Addrs:    db.HostURLs,
		Username: db.AdminUser,
		Password: db.AdminPass,

		DialServer: func(addr *mgo.ServerAddr) (net.Conn, error) {
			return tls.Dial("tcp", addr.String(), &tls.Config{})
		},
		Timeout: time.Second * 10,
	}

	s, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		panic(err)
	}
	return s
}
