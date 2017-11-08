package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/Stektpotet/imt2681-assignment2/database"
	"github.com/Stektpotet/imt2681-assignment2/fixer"
	"github.com/Stektpotet/imt2681-assignment2/util"
	"github.com/Stektpotet/imt2681-assignment2/webhook"
	"gopkg.in/mgo.v2/bson"
)

const ( //DATABASE COLLECTIONS
	dbCurrencyCollection = "currency"
	dbWebhookCollection  = "webhook"
)

const ( //PATHS
	rootPath    = "/api/"
	latestPath  = rootPath + "latest/"
	averagePath = rootPath + "average/"
	triggerPath = rootPath + "evaluationtrigger/"
)

const ( // OTHER CONSTSTANTS
	baseCurrency   = "EUR"
	fixerPath      = "base=" + baseCurrency
	tickerInterval = time.Second /*Minute*/ * 24
)

var globalDB database.DBStorage

func initializeDBConnection() {
	var mongoDBHosts = []string{
		"cluster0-shard-00-00-qvogu.mongodb.net:27017",
		"cluster0-shard-00-01-qvogu.mongodb.net:27017",
		"cluster0-shard-00-02-qvogu.mongodb.net:27017",
	}

	globalDB = &database.MongoDB{
		HostURLs:  mongoDBHosts,
		AdminUser: util.GetEnv("TRACKER_USER"),
		AdminPass: util.GetEnv("TRACKER_PASS"),
		Name:      "currencytrackr",
	}
	globalDB.Init()
	//To make sure there will always be at least 3 entries in the db.
	//(saturday/sunday wont have new entries)
	addEntriesForNPastDays(5)
}

func addEntriesForNPastDays(n int) {
	t := time.Now()
	for i := 0; i < n; i++ {
		globalDB.Add(dbCurrencyCollection, fixer.GetCurrencies(util.DateString(t.Date())))
		t = t.AddDate(0, 0, -1)
	}
}

func main() {
	log.Println("Running CurrencyTrackR")
	// globalDB = &database.CurrencyDB{}
	initializeDBConnection()
	http.HandleFunc(rootPath, SubscriptionHandler)
	http.HandleFunc(latestPath, LatestHandler)
	http.HandleFunc(averagePath, AverageHandler)
	http.HandleFunc(triggerPath, EvaluationTriggerHandler)
	log.Println(util.GetEnv("PORT"))
	http.ListenAndServe(":"+util.GetEnv("PORT"), nil)
}

//SubscriptionHandler - Handles all subscription-related requests
// this includes:
// GET:     root/api/:id
// POST:    root/api/
// DELETE:  root/api/:id
func SubscriptionHandler(w http.ResponseWriter, r *http.Request) {
	status := http.StatusOK
	writeStatus := true
	response := []byte{}
	writeResponse := false

	switch r.Method {
	case http.MethodPost:
		//REGISTER SUBSCRIPTION
		r, ok := subscriptionRegister(r)
		if ok {
			status = http.StatusCreated
			writeResponse = true
			writeStatus = false
			response = []byte(r)
		} else {
			status = http.StatusBadRequest
		}

	case http.MethodGet:
		//GET SUBSCRIPTION
		subscriber, ok := subscriptionGet(r.URL.Path)
		if ok {
			w.Header().Add("content-type", "application/json")
			r, err := json.Marshal(subscriber)
			if err != nil {
				status = http.StatusInternalServerError
			} else {
				response = r
				writeResponse = true
				writeStatus = false
			}
		} else {
			status = http.StatusNotFound
		}
	case http.MethodPut:
		//UPDATE SUBSCRIPTION
		status = http.StatusNotImplemented
		writeStatus = true

	case http.MethodDelete:
		ok := subscriptionDelete(r.URL.Path)
		if ok {
			status = http.StatusAccepted //202
		} else {
			status = http.StatusNotFound //404
		}

	default:
		status = http.StatusMethodNotAllowed
	}
	w.WriteHeader(status)
	if writeResponse {
		w.Write(response)
	}
	if writeStatus {
		fmt.Fprintf(w, "%v\n%s", status, http.StatusText(status))
	}
	return
}

func subscriptionRegister(r *http.Request) (subID string, success bool) {
	parts := strings.Split(r.URL.Path, "/") //root/api/:id, ID = parts[2]
	if len(parts) != 3 {
		success = false
		return
	}
	rBody, err := ioutil.ReadAll(r.Body)
	if err != nil { //failed to read body, i.e. bad request
		log.Printf("Failed reading body of request: %+v", r.Body)
		success = false
		return
	}

	var hook webhook.SubsciptionIn

	err = json.Unmarshal(rBody, &hook)
	if err != nil { //falied to unmarshal, don't add to db
		success = false
		return
	}
	hook.HookID = bson.NewObjectId().Hex()
	err = globalDB.Add(dbWebhookCollection, hook)
	if err != nil { //failed adding to db, don't return success
		success = false
		return
	}
	subID = hook.HookID
	return
}

func subscriptionGet(URLpath string) (sub webhook.SubsciptionOut, success bool) {
	sub = webhook.SubsciptionOut{}
	parts := strings.Split(URLpath, "/") //host/root/:id,  ID = parts[2]
	if len(parts) != 3 {
		success = false
		return
	}
	success = globalDB.Get(dbWebhookCollection, bson.M{"hookid": parts[2]}, &sub)
	return
}

func subscriptionDelete(URLpath string) (success bool) {
	parts := strings.Split(URLpath, "/") //root/api/:id, ID = parts[2]
	if len(parts) != 3 {
		success = false
		return
	} //root/api/:id
	success = globalDB.Delete(dbWebhookCollection, bson.M{"hookid": parts[2]})
	return
}

//LatestHandler - Handles requests to root/latest
//finds the currency rating for the given base/target-conversion
func LatestHandler(w http.ResponseWriter, r *http.Request) {
	status := http.StatusOK
	if r.Method != http.MethodPost {
		status = http.StatusMethodNotAllowed
		w.WriteHeader(status)
		fmt.Fprint(w, status, http.StatusText(status), "\nAccepted Methods: ", http.MethodPost)
		return
	}

	writeResponse := true
	entry := new(fixer.Currency)
	conversion := fixer.Conversion{}
	//Obtain requested conversion as object
	// body, err := ioutil.ReadAll(r.Body)
	err := json.NewDecoder(r.Body).Decode(&conversion)
	if err != nil {
		log.Printf("Failed reading body of request: %+v", r.Body)
		status = http.StatusBadRequest //400
		writeResponse = false
	} else {
		// err := json.Unmarshal(body, &conversion)
		// if err != nil {
		// 	log.Printf("Failed reading body of request: %+v", r.Body)
		// 	status = http.StatusBadRequest //400
		// 	writeResponse = false
		// }
		if !findLastEntry(entry) { //No date exists
			status = http.StatusNoContent
			writeResponse = false
		}
	}
	w.WriteHeader(status)
	if writeResponse {
		entry.Rates[entry.Base] = 1 //ensure a given value for this systems base Currency
		fmt.Fprint(w, entry.Rates[conversion.Target]/entry.Rates[conversion.Base])
	} else {
		fmt.Fprint(w, status, http.StatusText(status))
	}
}

func findLastEntry(entry *fixer.Currency) bool {
	//FIND LATEST ENTRY
	n := time.Now()
	entriesRemaining := globalDB.Count(dbCurrencyCollection)
	found := false
	for !found && entriesRemaining+5 >= 0 {
		latestDate := util.DateString(n.Date())
		found = globalDB.Get(dbCurrencyCollection, bson.M{"date": latestDate}, entry)
		n = n.AddDate(0, 0, -1)
	}
	return found
}

func findNLatestEntries(n int) (entries []fixer.Currency, ok bool) {
	ok = false
	t := time.Now()
	entries = make([]fixer.Currency, 0, n)
	remaining := globalDB.Count(dbCurrencyCollection)
	if remaining < n {
		return //too few entries
	}
	f := 0
	for ; f < n && remaining+5 >= 0; remaining-- {
		latestDate := util.DateString(t.Date())
		entry := fixer.Currency{}
		if globalDB.Get(dbCurrencyCollection, bson.M{"date": latestDate}, &entry) {
			entries = append(entries, entry)
			f++
		}
		t = t.AddDate(0, 0, -1)
	}
	ok = f == n
	return
}

//AverageHandler - find the average currency rating for the
//given base/target-conversion for the past three days
func AverageHandler(w http.ResponseWriter, r *http.Request) {
	status := http.StatusOK
	if r.Method != http.MethodPost {
		status = http.StatusMethodNotAllowed
		w.WriteHeader(status)
		fmt.Fprint(w, status, http.StatusText(status), "\nAccepted Methods: ", http.MethodPost)
		return
	}
	writeResponse := true
	var entries []fixer.Currency
	var conversion fixer.Conversion
	//Obtain requested conversion as object
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Failed reading body of request: %+v", r.Body)
		status = http.StatusBadRequest //400
		writeResponse = false
	} else {
		json.Unmarshal(body, &conversion)
		if err != nil {
			log.Printf("Failed reading body of request: %+v", r.Body)
			status = http.StatusBadRequest //400
			writeResponse = false
		}
		e, ok := findNLatestEntries(3)
		if !ok {
			status = http.StatusNoContent
			writeResponse = false
		}
		entries = e
	}

	w.WriteHeader(status)
	if writeResponse {
		var sum float32
		for _, entry := range entries {
			entry.Rates[entry.Base] = 1 //ensure a given value for this systems base Currency
			sum += entry.Rates[conversion.Target] / entry.Rates[conversion.Base]
		}
		fmt.Fprint(w, sum)
	} else {
		fmt.Fprint(w, status, http.StatusText(status))
	}
}

//EvaluationTriggerHandler - Invokes all triggers, ignoring min max ranges
func EvaluationTriggerHandler(w http.ResponseWriter, r *http.Request) {
	hooks := []webhook.SubsciptionOut{}
	var current fixer.Currency
	findLastEntry(&current)
	current.Rates[current.Base] = 1
	globalDB.GetAll(dbWebhookCollection, &hooks)
	fmt.Fprintf(w, "%+v", hooks)
	for _, hook := range hooks {
		hookRate := current.Rates[hook.Target] / current.Rates[hook.Base]
		hook.Invoke(hookRate, *http.DefaultClient)
	}
}
