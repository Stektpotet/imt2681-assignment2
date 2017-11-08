package fixer

import (
	"encoding/json"
	"io/ioutil"
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

//GetLatest - Shorthand function for GetCurrencies("latest")
//allows for extra constraints available in the fixer api
func GetLatest(constraints string) (payload Currency, err error) {

	response, err := http.Get(latestPath + "?" + constraints)
	if err != nil {
		log.Printf("Failed obtaining currencies from Fixer: %+v", err.Error())
		return
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("Failed reading body: %+v", err.Error())
		return
	}
	err = json.Unmarshal(body, &payload)
	if err != nil {
		log.Printf("Failed unmarshaling body: %+v", err.Error())
		return
	}
	return
}

func GetCurrencies(constraints string) (payload Currency) {
	response, err := http.Get(fixerBaseURL + constraints)
	if err != nil {
		log.Printf("Failed obtaining currencies from Fixer: %+v", err.Error())
		log.Println("Using local example file: ./base.json")
		data, err := ioutil.ReadFile("./fixer/base.json")
		if err != nil {
			log.Fatal("Unable to read local example file ./base.json")
		}
		json.Unmarshal(data, &payload)
	} else {
		defer response.Body.Close()
		err = json.NewDecoder(response.Body).Decode(&payload)
	}

	if err != nil {
		body, _ := ioutil.ReadAll(response.Body)
		log.Fatalf("Failed decoding Fixer-payload:%+v\n\n%+v", err, body)
	}
	return
}
