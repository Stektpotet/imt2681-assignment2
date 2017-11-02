package database

import (
	"fmt"
	"log"

	"github.com/stektpotet/imt2681-assignment2/fixer"
	mgo "gopkg.in/mgo.v2"
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
	currencyCollectionIndex := mgo.Index{
		Key:        []string{"date"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
	webhookCollectionIndex := currencyCollectionIndex
	webhookCollectionIndex.Key = []string{"hookid"}
	db.ensureIndex(session, "currency", currencyCollectionIndex)
	db.ensureIndex(session, "webhook", webhookCollectionIndex)
}

func (db *CurrencyMongoDB) ensureIndex(s *mgo.Session, c string, i mgo.Index) {
	err := s.DB(db.Name).C(c).EnsureIndex(i)
	if err != nil {
		log.Fatalf("ERROR: %s", err.Error())
	}
}

func (db *CurrencyMongoDB) Add(collection string, record interface{}) (err error) {
	session := db.CreateSession()
	defer session.Close()
	err = session.DB(db.Name).C(collection).Insert(record)
	return
}

func (db *CurrencyMongoDB) Count(collection string) int {
	session := db.CreateSession()
	defer session.Close()

	count, err := session.DB(db.Name).C(collection).Count()
	if err != nil {
		fmt.Printf("Unable to count: %s", err.Error())
	}
	return count
}
func (db *CurrencyMongoDB) Get(collection string, query bson.M, data interface{}) (ok bool) {
	session := db.CreateSession()
	defer session.Close()

	ok = true
	err := session.DB(db.Name).C(collection).Find(query).One(data)
	if err != nil {
		ok = false
	}
	return
}

func (db *CurrencyMongoDB) Delete(collection string, query bson.M) (ok bool) {
	session := db.CreateSession()
	defer session.Close()

	ok = true
	info, err := session.DB(db.Name).C(collection).RemoveAll(query)
	if err != nil || info.Removed == 0 {
		ok = false
	}

	return
}

func (db *CurrencyMongoDB) GetAll(collection string, data interface{}) {
	session := db.CreateSession()
	defer session.Close()

	// elements := make([]interface{}, 0, db.Count(collection))
	err := session.DB(db.Name).C(collection).Find(bson.M{}).All(data)
	if err != nil {
		fmt.Printf("Unable to obtain all: %s", err.Error())
	}
	return
}
