package main

import (
	"fmt"
	"net/http"

	"github.com/stektpotet/imt2681-assignment2/database"
	"github.com/stektpotet/imt2681-assignment2/util"
)

//WebhookPayload - The payload of webhooks in the system
type WebhookPayload struct {
	URL             string  `json:"webhookURL"`
	Base            string  `json:"baseCurrency"`
	Target          string  `json:"targetCurrency"`
	MinTriggerValue float32 `json:"minTriggerValue"`
	MaxTriggerValue float32 `json:"maxTriggerValue"`
}

const (
	rootPath    = "/api/"
	latestPath  = rootPath + "latest/"
	averagePath = rootPath + "average/"
	triggerPath = rootPath + "evaluationtrigger"
	mongodbURL  = "mongodb://localhost"
)

func sanetizeHook(hookURL string) string {
	return fmt.Sprintf("Sanetization not implemented. Hook:\n%+v", hookURL)
}

func serviceHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	switch r.Method {
	case http.MethodPost:

	case http.MethodGet:

	case http.MethodDelete:

	}
}

func latestHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

}

func averageHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
}

func evaluationTriggerHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	//obtain database's webhook collection
}

var globalDB database.DBStorage

func main() {

	var mongoDBHosts = []string{
		"cluster0-shard-00-00-qvogu.mongodb.net:27017",
		"cluster0-shard-00-01-qvogu.mongodb.net:27017",
		"cluster0-shard-00-02-qvogu.mongodb.net:27017",
	}

	globalDB = &database.CurrencyMongoDB{
		MongoDB: &database.MongoDB{
			HostURLs:       mongoDBHosts,
			AdminUser:      util.GetEnv("WEBHOOK_USER"),
			AdminPass:      util.GetEnv("WEBHOOK_PASS"),
			Name:           "currencytrackr",
			CollectionName: "currency",
		},
	}
	globalDB.Init()

	http.HandleFunc(rootPath, serviceHandler)
	http.HandleFunc(latestPath, latestHandler)
	http.HandleFunc(averagePath, averageHandler)
	http.HandleFunc(triggerPath, evaluationTriggerHandler)
	http.ListenAndServe(":"+util.GetEnv("PORT"), nil)
}
