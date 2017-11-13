package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/Stektpotet/imt2681-assignment2/database"
	"github.com/Stektpotet/imt2681-assignment2/fixer"
	"github.com/Stektpotet/imt2681-assignment2/webhook"
)

func TestMain(m *testing.M) {
	globalDB = &database.MongoDB{
		HostURLs:  []string{"localhost"},
		AdminUser: "tester",
		AdminPass: "WA9LI7f2DbVQtvbM",
		Name:      "test",
	}
	globalDB.Init()
	globalDB.DropCollection("webhook")

	c := m.Run()

	globalDB.DropCollection("webhook")
	globalDB.Drop()
	os.Exit(c)
}

func PostHere(url, contentType string, r io.Reader) (resp *http.Response, err error) {
	resp = &http.Response{}
	resp.StatusCode = http.StatusOK
	resp.Header = http.Header{}
	resp.Header.Add("contentType", "application/json")
	raw, err := ioutil.ReadFile("../../samples/base.json")
	if err != nil {
		log.Fatal(err) //could not read from file
	}
	resp.Body = ioutil.NopCloser(bytes.NewBuffer(raw))
	return
}

func Test_initializeDBConnection(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		name         string
		mongoDBHosts []string
	}{
		{
			"Valid",
			[]string{"localhost"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			initializeDBConnection(tt.mongoDBHosts)
		})
	}
}

func TestInvokeHooks(t *testing.T) {
	globalDB.DropCollection("webhook")

	testHookCount := 16

	curr := fixer.Currency{
		Base:  "EUR",
		Rates: map[string]float32{"NOK": float32(testHookCount)},
	}

	baseTestHook := webhook.SubsciptionOut{
		Base:   curr.Base,
		Target: "NOK",
		Min:    1,
		Max:    float32(testHookCount),
	}

	hooks := make([]webhook.SubsciptionOut, 0, testHookCount)
	for i := 0; i < testHookCount; i++ {
		hooks = append(hooks, baseTestHook)
		hooks[i].Max -= float32(i)
		globalDB.Add("webhook", hooks[i])
	}

	tests := []struct {
		name         string
		currentRate  float32
		successCount int
	}{
		{"Valid ", 16, 1},
		{"Valid ", 15, 2},
		{"Valid ", 14, 3},
		{"Valid ", 13, 4},
		{"Valid ", 12, 5},
		{"Valid ", 11, 6},
		{"Valid ", 10, 7},
		{"Valid ", 9, 8},
		{"Valid ", 8, 9},
		{"Valid ", 7, 10},
		{"Valid ", 6, 11},
		{"Valid ", 5, 12},
		{"Valid ", 4, 13},
		{"Valid ", 3, 14},
		{"Valid ", 2, 15},
		{"Valid ", 1, 16},
	}
	for i, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			curr.Rates[hooks[i].Target] = tt.currentRate
			successCount := invokeHooks(curr, PostHere)
			if tt.successCount != successCount {
				t.Errorf("main.invokeHooks() = %v, want %v", successCount, tt.successCount)
			}
		})
	}
}
