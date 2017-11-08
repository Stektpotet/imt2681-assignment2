package database

import (
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/Stektpotet/imt2681-assignment2/util"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//DBStorage - Interface of a MongoDB database
type DBStorage interface {
	CreateSession() *mgo.Session
	Init()
	Add(collection string, data interface{}) error
	Count(collection string) int
	Get(collection string, query bson.M, data interface{}) (found bool)
	Delete(collection string, query bson.M) (found bool)
	GetAll(collection string, data interface{})
	DropCollection(collection string)
	Drop()
}

//MongoDB - MongoDB database
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

// CreateSession - create DB Collection session
func (db *MongoDB) CreateSession() *mgo.Session {

	if util.Contains(db.HostURLs, "localhost") {
		return db.CreateLocalSession() //No need for dialing with info
	}
	//ELSE
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

//Init - initialize database object db with essential properties for use in this system
func (db *MongoDB) Init() {
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

func (db *MongoDB) ensureIndex(s *mgo.Session, c string, i mgo.Index) {
	err := s.DB(db.Name).C(c).EnsureIndex(i)
	if err != nil {
		log.Fatalf("ERROR: %s", err.Error())
	}
}

//Add - Adds element to the collection
func (db *MongoDB) Add(collection string, element interface{}) (err error) {
	session := db.CreateSession()
	defer session.Close()
	err = session.DB(db.Name).C(collection).Insert(element)
	return
}

//Count - Count elements in collection
func (db *MongoDB) Count(collection string) int {
	session := db.CreateSession()
	defer session.Close()

	count, err := session.DB(db.Name).C(collection).Count()
	if err != nil {
		fmt.Printf("Unable to count: %s", err.Error())
	}
	return count
}

//Get - retrieve an element from a collection where the element matches the query
func (db *MongoDB) Get(collection string, query bson.M, data interface{}) (ok bool) {
	session := db.CreateSession()
	defer session.Close()

	ok = true
	err := session.DB(db.Name).C(collection).Find(query).One(data)
	log.Print()
	if err != nil {
		ok = false
	}
	return
}

//Delete - remove all elements in collection matching the query
func (db *MongoDB) Delete(collection string, query bson.M) (ok bool) {
	session := db.CreateSession()
	defer session.Close()

	ok = true
	info, err := session.DB(db.Name).C(collection).RemoveAll(query)
	if err != nil || info.Removed == 0 {
		ok = false
	}

	return
}

//GetAll - retrieve all elements in a collection
func (db *MongoDB) GetAll(collection string, data interface{}) {
	session := db.CreateSession()
	defer session.Close()

	// elements := make([]interface{}, 0, db.Count(collection))
	err := session.DB(db.Name).C(collection).Find(bson.M{}).All(data)
	if err != nil {
		fmt.Printf("Unable to obtain all: %s", err.Error())
	}
	return
}

//DropCollection - Drop the specified collection
func (db *MongoDB) DropCollection(collection string) {
	session := db.CreateSession()
	defer session.Close()

	err := session.DB(db.Name).C(collection).DropCollection()
	if err != nil {
		fmt.Printf("Unable to drop collection: %s", err.Error())
	}
	return
}

//Drop - Drop the database, (Init will restore it if called afterwards)
func (db *MongoDB) Drop() {
	session := db.CreateSession()
	defer session.Close()

	err := session.DB(db.Name).DropDatabase()
	if err != nil {
		fmt.Printf("Unable to drop database: %s", err.Error())
	}
	return
}
