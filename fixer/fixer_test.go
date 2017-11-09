package fixer

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"testing"
)

func GetFromHere(url string) (resp *http.Response, err error) {
	resp = &http.Response{}
	data, err := ioutil.ReadFile("../samples/base.json")
	if err != nil {
		log.Fatal("Unable to read local sample file")
	}
	resp.Body = ioutil.NopCloser(bytes.NewBuffer(data))
	return
}

func TestGetLatest(t *testing.T) {
	wantedPayload := Currency{}
	r, _ := GetFromHere("")

	rBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		t.Error(err)
	}
	err = json.Unmarshal(rBody, &wantedPayload)
	if err != nil {
		t.Error(err)
	}

	tests := []struct {
		name        string
		constraints string
		wantPayload Currency
	}{
		{"Test 1", "", wantedPayload},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPayload := getLatest(tt.constraints, GetFromHere)
			if !reflect.DeepEqual(gotPayload, tt.wantPayload) {
				t.Errorf("GetLatest() = %v, want %v", gotPayload, tt.wantPayload)
			}
		})
	}
}

func TestGetCurrencies(t *testing.T) {
	wantedPayload := Currency{}
	r, _ := GetFromHere("")

	rBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		t.Error(err)
	}
	err = json.Unmarshal(rBody, &wantedPayload)
	if err != nil {
		t.Error(err)
	}

	tests := []struct {
		name        string
		constraints string
		wantPayload Currency
	}{
		{"Test 1", "", wantedPayload},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPayload := getCurrencies(tt.constraints, GetFromHere)
			if !reflect.DeepEqual(gotPayload, tt.wantPayload) {
				t.Errorf("GetLatest() = %v, want %v", gotPayload, tt.wantPayload)
			}
		})
	}
}
