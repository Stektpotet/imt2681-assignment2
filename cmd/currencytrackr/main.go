package main

import (
	"log"
	"os"
	"time"

	"github.com/stektpotet/imt2681-assignment2/database"
	"github.com/stektpotet/imt2681-assignment2/fixer"
)

const (
	mongodbURL = "mongodb://localhost"
	fixerPath  = "latest?base=EUR"
)

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

// Tick - Updates the currency database
func Tick(db database.DBStorage) {
	payload := fixer.GetCurrencies(fixerPath)
	db.Add(payload)
	// db.Add(fixer.GetCurrencies(fixerURL))

}

var globalDB database.DBStorage

func GetPort() (port string) {
	port = os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
		// port = "5000"
	}
	return
}

func main() {
	log.Println("Running CurrencyTrackR")
	// globalDB = &database.CurrencyDB{}

	globalDB = &database.CurrencyMongoDB{
		MongoDB: &database.MongoDB{
			HostURL:        mongodbURL,
			Name:           "currencytrackr",
			CollectionName: "currency",
		},
	}

	globalDB.Init()

	Tick(globalDB)
	ticker := time.NewTicker(time.Second * 4)

	for _ = range ticker.C {
		log.Printf("%+v: Tick!", time.Now())
		Tick(globalDB)
	}
}

/*

TODO:
-
- Eventtriggers
- URL to 'hook'
- Payload format

*/
