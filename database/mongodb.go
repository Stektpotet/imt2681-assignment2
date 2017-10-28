package database

import (
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
	HostURL        string
	Name           string
	CollectionName string
}

// CreateSession - create DB session
func (db *MongoDB) CreateSession() *mgo.Session {
	s, err := mgo.Dial(db.HostURL)
	if err != nil {
		panic(err)
	}
	return s
}
