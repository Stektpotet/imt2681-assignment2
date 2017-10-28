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

func GetCurrencies(fixerPath string) (payload Currency) {
	response, err := http.Get(fixerBaseURL + fixerPath)
	if err != nil {
		log.Fatalf("Failed obtaining currencies from Fixer: %+v", err.Error())
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&payload)

	if err != nil {
		log.Fatalf("Failed decoding Fixer-payload:%+v\n\n", err)
	}
	return
}
