package fixer

import (
	"encoding/json"
	"log"
	"net/http"
)

const fixerBaseURL = "http://api.fixer.io/"

type Currency struct {
	Base  string             `json:"base"`
	Date  string             `json:"date"`
	Rates map[string]float32 `json:"rates"`
}

func GetCurrencies(fixerURL string) Currency {
	response, err := http.Get(fixerBaseURL + fixerURL)
	if err != nil {
		log.Fatalf("Failed obtaining currencies from Fixer: %+v", err.Error())
	}
	defer response.Body.Close()

	var p Currency
	err = json.NewDecoder(response.Body).Decode(&p)

	if err != nil {
		log.Fatalf("Failed decoding Fixer-payload:%+v\n\n", err)
	}

	return p
}
