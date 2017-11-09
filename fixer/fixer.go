package fixer

import (
	"encoding/json"
	"log"
	"net/http"
)

const (
	fixerBaseURL = "http://api.fixer.io/"
	latestPath   = fixerBaseURL + "latest"
)

//Currency - the default payload of the fixer.io api
type Currency struct {
	Base  string             `json:"base"`
	Date  string             `json:"date"`
	Rates map[string]float32 `json:"rates"`
}

//Conversion - Encapsulated conversion request
//Use as keys in Currency.Rates to find relation
type Conversion struct {
	Base   string `json:"baseCurrency"`
	Target string `json:"targetCurrency"`
}

//getFunc - type to allow for injecting the getter functionality
type getFunc func(string) (*http.Response, error)

//GetLatest - Shorthand function for GetCurrencies("latest")
//allows for extra constraints available in the fixer api
func GetLatest(constraints string) (payload Currency) {
	return getLatest(constraints, http.DefaultClient.Get)
}

func getLatest(constraints string, getter getFunc) (payload Currency) {
	return getCurrencies("latest?"+constraints, getter)
}

// GetCurrencies - Obtain a Currency object based on the constraints given to the
//fixer.io api
func GetCurrencies(constraints string) (payload Currency) {
	return getCurrencies(constraints, http.DefaultClient.Get)
}

func getCurrencies(constraints string, getter getFunc) (payload Currency) {
	response, err := getter(fixerBaseURL + constraints)
	defer response.Body.Close()
	if err != nil {
		log.Printf("Failed obtaining currencies: %+v", err.Error())
		return
	}
	err = json.NewDecoder(response.Body).Decode(&payload)
	if err != nil {
		log.Fatalf("Failed decoding Fixer-payload:%+v", err)
	}
	return
}
