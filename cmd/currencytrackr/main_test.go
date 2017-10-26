package main

import (
	"log"
	"testing"

	"github.com/stektpotet/imt2681-assignment2/database"
	"github.com/stektpotet/imt2681-assignment2/payload"
)

func TestServiceHandler(t *testing.T) {
	// tempDB := mongodb.MongoDB{}
	// tempDB.Init()
	//Store some copy of the data
	//Modify db
	// UpdateCurrencies(&tempDB)
	//extract data
	//Compare to old data
}
func TestUpdateCurrencies(t *testing.T) {
	var testDB database.DBStorage
	testDB = &database.CurrencyDB{}
	p1 := payload.GetFixerCurrencies("http://api.fixer.io/latest")
	p := payload.GetFixerCurrencies("http://api.fixer.io/2017-10-25")
	testDB.Init()
	testDB.Add(p)
	v, _ := testDB.Get(p.Date)
	log.Printf("Something: %+v\n\n\n", v)

	UpdateCurrencies(testDB)
	v, _ = testDB.Get(p1.Date)
	log.Printf("Something: %+v\n\n\n", v)
}

func TestUpdateCurrenciesMongo(t *testing.T) {
	// var TestDB database.DBStorage
}
