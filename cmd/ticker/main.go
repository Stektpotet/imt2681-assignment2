package main

import (
	"log"
	"net/http"
	"time"

	"github.com/Stektpotet/imt2681-assignment2/database"
	"github.com/Stektpotet/imt2681-assignment2/fixer"
	"github.com/Stektpotet/imt2681-assignment2/util"
	"github.com/Stektpotet/imt2681-assignment2/webhook"
)

const (
	fixerPath      = "base=EUR"
	tickerInterval = time.Minute * 10
)

func initializeDBConnection(mongoDBHosts []string) {

	globalDB = &database.MongoDB{
		HostURLs:  mongoDBHosts,
		AdminUser: util.GetEnv("TRACKER_USER"),
		AdminPass: util.GetEnv("TRACKER_PASS"),
		Name:      "currencytrackr",
	}
	globalDB.Init()
}

//Tick - Runs at an interval to continuously keep the db up to date
//and to invoke the webhooks
func Tick() {
	newEntry := fixer.GetLatest(fixerPath)
	err := globalDB.Add("currency", newEntry)
	if err == nil {
		log.Printf("invoked %d on new entry:\n%v", InvokeHooks(newEntry), newEntry)
	}
}

//InvokeHooks - Invoke all hooks that are subscribed within the given rate and retun number of envoked elements
func InvokeHooks(current fixer.Currency) int {
	return invokeHooks(current, http.DefaultClient.Post)
}

//InvokeHooks - Invoke all hooks that are subscribed within the given rate
func invokeHooks(current fixer.Currency, poster webhook.PostFunc) (invocations int) {
	invocations = 0
	hooks := []webhook.SubsciptionOut{}
	current.Rates[current.Base] = 1
	globalDB.GetAll("webhook", &hooks)
	for _, hook := range hooks {
		hookRate := current.Rates[hook.Target] / current.Rates[hook.Base]
		if hook.Min <= hookRate && hook.Max >= hookRate {
			hook.Invoke(hookRate, poster)
			invocations++
		}
	}
	return
}

var globalDB database.DBStorage

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
		Tick()
	}
}
