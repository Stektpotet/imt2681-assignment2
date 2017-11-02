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

// type CurrencyOut struct {
// 	Base  string             `json:"base"`
// 	Date  string             `json:"date"`
// 	Rates map[string]float32 `json:"rates"`
// }

type Currency struct {
	Base  string             `json:"base"`
	Date  string             `json:"date"`
	Rates map[string]float32 `json:"rates"`
}

type Conversion struct {
	Base   string `json:"baseCurrency"`
	Target string `json:"targetCurrency"`
}

func GetLatest(constraints string) (payload Currency, err error) {
	response, err := http.Get(latestPath + "?" + constraints)
	body, err := ioutil.ReadAll(response.Body)

	json.Unmarshal(body, &payload)
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
