package main

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/Stektpotet/imt2681-assignment2/database"
	"github.com/Stektpotet/imt2681-assignment2/fixer"
	"github.com/Stektpotet/imt2681-assignment2/webhook"
	"github.com/subosito/gotenv"
)

func init() {
	gotenv.MustLoad("../../.env")
}

func TestMain(m *testing.M) {
	var mongoDBHosts = []string{
		"cluster0-shard-00-00-qvogu.mongodb.net:27017",
		"cluster0-shard-00-01-qvogu.mongodb.net:27017",
		"cluster0-shard-00-02-qvogu.mongodb.net:27017",
	}
	globalDB = &database.MongoDB{
		HostURLs:  mongoDBHosts,
		AdminUser: "tester",
		AdminPass: "WA9LI7f2DbVQtvbM",
		Name:      "test",
	}
	globalDB.Init()
	globalDB.DropCollection("webhook")
	defer globalDB.DropCollection("webhook")
	defer globalDB.Drop()

	c := m.Run()
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
			[]string{
				"cluster0-shard-00-00-qvogu.mongodb.net:27017",
				"cluster0-shard-00-01-qvogu.mongodb.net:27017",
				"cluster0-shard-00-02-qvogu.mongodb.net:27017",
			},
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
	bigRangeHook := webhook.SubsciptionOut{
		Base:   "NOK",
		Target: "EUR",
		URL:    "localhost",
		Min:    1,
		Max:    10,
	}
	smallRangeHook := webhook.SubsciptionOut{
		Base:   "NOK",
		Target: "EUR",
		URL:    "localhost",
		Min:    1,
		Max:    1.1,
	}

	curr := fixer.Currency{}

	r, _ := PostHere("localhost", "application/json", nil)

	json.NewDecoder(r.Body).Decode(&curr)

	tests := []struct {
		name         string
		hooks        []webhook.SubsciptionOut
		successCount int
	}{
		{
			"Valid 1",
			[]webhook.SubsciptionOut{
				bigRangeHook,
				bigRangeHook,
				smallRangeHook,
			},
			2,
		},
		{
			"Valid 2",
			[]webhook.SubsciptionOut{
				bigRangeHook,
				bigRangeHook,
				bigRangeHook,
				bigRangeHook,
				smallRangeHook,
				smallRangeHook,
				smallRangeHook,
			},
			4,
		},
		{
			"Valid 3",
			[]webhook.SubsciptionOut{
				bigRangeHook,
				smallRangeHook,
				bigRangeHook,
				smallRangeHook,
			},
			2,
		},
		{
			"Not Valid 1",
			[]webhook.SubsciptionOut{
				smallRangeHook,
				smallRangeHook,
				smallRangeHook,
			},
			0,
		},
		{
			"Not Valid ",
			[]webhook.SubsciptionOut{
				smallRangeHook,
				smallRangeHook,
				smallRangeHook,
				smallRangeHook,
				smallRangeHook,
				smallRangeHook,
				smallRangeHook,
				smallRangeHook,
				smallRangeHook,
				smallRangeHook,
				smallRangeHook,
				smallRangeHook,
				smallRangeHook,
				smallRangeHook,
				smallRangeHook,
			},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			invokeHooks(curr, PostHere)
		})
	}
}
