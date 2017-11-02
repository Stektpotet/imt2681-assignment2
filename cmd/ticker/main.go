package main

import (
	"log"
	"time"

	"github.com/stektpotet/imt2681-assignment2/database"
	"github.com/stektpotet/imt2681-assignment2/fixer"
	"github.com/stektpotet/imt2681-assignment2/util"
	"github.com/stektpotet/imt2681-assignment2/webhook"
)

var globalDB database.DBStorage

func initializeDBConnection(mongoDBHosts []string) {

	globalDB = &database.CurrencyMongoDB{
		MongoDB: &database.MongoDB{
			HostURLs:  mongoDBHosts,
			AdminUser: util.GetEnv("TRACKER_USER"),
			AdminPass: util.GetEnv("TRACKER_PASS"),
			Name:      "currencytrackr",
		},
	}
	globalDB.Init()
}

func Tick() {
	newEntry, err := fixer.GetLatest(fixerPath)
	if err != nil {
		log.Println(err)
	}
	err = globalDB.Add("currency", newEntry)
	if err == nil {
		InvokeHooks(newEntry)
	}
}

func InvokeHooks(current fixer.Currency) {
	hooks := []webhook.SubsciptionOut{}
	current.Rates[current.Base] = 1
	globalDB.GetAll("webhook", &hooks)
	for _, hook := range hooks {
		hookRate := current.Rates[hook.Target] / current.Rates[hook.Base]
		if hook.Min <= hookRate && hook.Max >= hookRate {
			hook.Invoke(hookRate)
		}
	}
}

const (
	fixerPath      = "base=EUR"
	tickerInterval = time.Second /*Minute*/ * 24
)

func main() {
	hosts := []string{
		"cluster0-shard-00-00-qvogu.mongodb.net:27017",
		"cluster0-shard-00-01-qvogu.mongodb.net:27017",
		"cluster0-shard-00-02-qvogu.mongodb.net:27017",
	}

	initializeDBConnection(hosts)
	Tick()
	ticker := time.NewTicker(tickerInterval) //util.UntilTomorrow())
	for _ = range ticker.C {
		log.Printf("%+v: Tick!", time.Now())
		Tick()
	}
}
