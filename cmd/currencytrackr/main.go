package main

import (
	"log"
	"time"

	"github.com/stektpotet/imt2681-assignment2/database"
	"github.com/stektpotet/imt2681-assignment2/fixer"
	"github.com/stektpotet/imt2681-assignment2/util"
)

const (
	fixerPath = "latest?base=EUR"
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

func main() {

	var mongoDBHosts = []string{
		"cluster0-shard-00-00-qvogu.mongodb.net:27017",
		"cluster0-shard-00-01-qvogu.mongodb.net:27017",
		"cluster0-shard-00-02-qvogu.mongodb.net:27017",
	}

	log.Println("Running CurrencyTrackR")
	// globalDB = &database.CurrencyDB{}

	globalDB = &database.CurrencyMongoDB{
		MongoDB: &database.MongoDB{
			HostURLs:       mongoDBHosts,
			AdminUser:      util.GetEnv("TRACKER_USER"),
			AdminPass:      util.GetEnv("TRACKER_PASS"),
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
