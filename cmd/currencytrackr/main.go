package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/stektpotet/imt2681-assignment2/database"
	"github.com/stektpotet/imt2681-assignment2/fixer"
)

const rootURL = "/api/"
const webhookBaseURL = "/hooks/"
const mongodbURL = "mongodb://localhost"
const fixerURL = "latest?base=EUR"

//ServiceHandler - TODO
func ServiceHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

}

// func ServiceHandler2(w http.ResponseWriter, r *http.Request) {
// 	payload := MessagePayload{}
// 	payload.Base = "base"
// 	payload.MaxTriggerValue = 10
// 	payload.MinTriggerValue = 1
// 	payload.Target = "target"
// 	payload.URL = "/test/"
//
// 	json.NewEncoder(w).Encode(payload)
// }

// UpdateCurrencies - Updates the currency database
func UpdateCurrencies(db database.DBStorage) {
	//TODO
	db.Add(fixer.GetCurrencies(fixerURL))
	// log.Printf("%+v", payload)
}

var GlobalDB database.DBStorage

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set.")
	}
	GlobalDB = &database.CurrencyDB{}
	// GlobalDB = &database.CurrencyDB{
	// 	"mongodb://localhost:" + port,
	// 	"currencyTrackr",
	// 	"currency",
	// }

	GlobalDB.Init()

	UpdateCurrencies(GlobalDB)
	// GetFixerCurrencies()
	ticker := time.NewTicker(time.Hour * 24)

	for _ = range ticker.C {
		UpdateCurrencies(GlobalDB)
		// GetFixerCurrencies()
	}
}

/*

TODO:
-
- Eventtriggers
- URL to 'hook'
- Payload format

*/
