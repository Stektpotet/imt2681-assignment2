package database

import (
	"fmt"

	"github.com/stektpotet/imt2681-assignment2/fixer"
	"gopkg.in/mgo.v2/bson"
)

type CurrencyOut struct {
	ID bson.ObjectId `bson:"_id,omitempty"`
	*fixer.Currency
}

type CurrencyMongoDB struct {
	*MongoDB
}

func (db *CurrencyMongoDB) Init() {
	session := db.CreateSession()
	defer session.Close()
	// err := session.DB(db.Name).C(db.CollectionName)
}

func (db *CurrencyMongoDB) Add(record fixer.Currency) (err error) {
	session := db.CreateSession()
	defer session.Close()

	err = session.DB(db.Name).C(db.CollectionName).Insert(record)
	return
}

func (db *CurrencyMongoDB) Count() int {
	session := db.CreateSession()
	defer session.Close()

	count, err := session.DB(db.Name).C(db.CollectionName).Count()
	if err != nil {
		fmt.Printf("Unable to count: %s", err.Error())
	}
	return count
}
func (db *CurrencyMongoDB) Get(key string) (value fixer.Currency, ok bool) {
	session := db.CreateSession()
	defer session.Close()

	value = fixer.Currency{}
	ok = true
	err := session.DB(db.Name).C(db.CollectionName).Find(bson.M{"date": key}).One(&value)
	if err != nil {
		ok = false
	}
	return
}
func (db *CurrencyMongoDB) GetAll() (currencies []fixer.Currency) {
	session := db.CreateSession()
	defer session.Close()

	currencies = make([]fixer.Currency, 0, db.Count())
	err := session.DB(db.Name).C(db.CollectionName).Find(bson.M{}).All(&currencies)
	if err != nil {
		fmt.Printf("Unable obtain all: %s", err.Error())
	}
	return
}
