package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/stektpotet/imt2681-assignment2/database"
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

func GetPort() (port string) {
	port = os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}
	return
}

var globalDB database.DBStorage

func main() {
	port := GetPort()

	globalDB = &database.CurrencyDB{}
	globalDB.Init()

	http.HandleFunc(rootPath, serviceHandler)
	http.HandleFunc(latestPath, latestHandler)
	http.HandleFunc(averagePath, averageHandler)
	http.HandleFunc(triggerPath, evaluationTriggerHandler)
	http.ListenAndServe(":"+port, nil)
}
