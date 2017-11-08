package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"

	"github.com/Stektpotet/imt2681-assignment2/database"
	"github.com/Stektpotet/imt2681-assignment2/fixer"
	"github.com/Stektpotet/imt2681-assignment2/webhook"
	"gopkg.in/mgo.v2/bson"
)

func Fail(err error, t *testing.T) {
	if err != nil {
		t.Error(err)
	}
}
func FailOk(ok bool, t *testing.T, context string) {
	if !ok {
		t.Error(context)
	}
}
func FailOkf(ok bool, t *testing.T, context string, args ...interface{}) {
	if !ok {
		t.Errorf(context, args)
	}
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
	globalDB.Drop()
	c := m.Run()
	globalDB.Drop()
	os.Exit(c)
}

var requestOnHandlerCode = func(method, path string, body []byte, handler http.HandlerFunc) int {
	request, _ := http.NewRequest(method, rootPath+path, bytes.NewBuffer(body))
	rr := httptest.NewRecorder()
	hfunc := http.HandlerFunc(handler)
	hfunc.ServeHTTP(rr, request)
	return rr.Code
}

func Expect(context string, got, expected interface{}, t *testing.T) {
	FailOk(got == expected, t, fmt.Sprintf("%s: got %v want %v", context, got, expected))
}

func Test_addEntriesForNPastDays(t *testing.T) {
	// DB := &globalDB
	//
	// expectedDates := []fixer.Currency{}
	// time := time.Now()
	// for i := 0; i < 7; i++ {
	// 	f := fixer.GetCurrencies(util.DateString(time.Date()))
	// 	expectedDates = append(expectedDates, f)
	// 	time = time.AddDate(0, 0, -1)
	// }
	//
	// c := gomock.NewController(t)
	// defer c.Finish()
	// mock := *database.NewMockDBStorage(c)
	// globalDB = &mock
	//
	// mock.EXPECT().Add(dbCurrencyCollection, gomock.Any()).AnyTimes()
	//
	// addEntriesForNPastDays(7)
	//
	// globalDB = *DB

}

var subscriptionRequest = func(method, ID string, body bool) *http.Request {
	raw := []byte{}
	if body {
		r, err := ioutil.ReadFile("../../samples/hook.json")
		if err != nil {
			log.Fatal(err)
		}
		raw = r
	}
	r, err := http.NewRequest(method, rootPath+ID, bytes.NewBuffer(raw))
	if err != nil {
		log.Fatal(err)
	}
	return r
}

func TestSubscriptionHandlerRRR(t *testing.T) {

	globalDB.DropCollection(dbWebhookCollection)

	var rSub webhook.SubsciptionOut
	r := subscriptionRequest(http.MethodPost, "", true)
	rBody, err := ioutil.ReadAll(r.Body)
	Fail(err, t)
	json.Unmarshal(rBody, &rSub)
	r = subscriptionRequest(http.MethodPost, "", true)
	id, ok := subscriptionRegister(r)

	FailOkf(ok, t, "subscribing with %+v should register, but does not.", *r)

	hookIDPath := r.URL.Path + id
	sub, ok := subscriptionGet(hookIDPath)
	FailOkf(ok, t, "getting with %+v should return the subscribed hook", id)
	Expect("Subscription in/out", sub, rSub, t)

	r.Method = http.MethodDelete
	subscriptionDelete(hookIDPath)

}

func Test_subscriptionRegister(t *testing.T) {
	//
	// var rSub webhook.SubsciptionOut
	// r := subscriptionRequest(http.MethodGet, "")
	// rBody, err := ioutil.ReadAll(r.Body)
	//
	// Fail(err, t)
	//
	// json.Unmarshal(rBody, &rSub)

	// raw, err := ioutil.ReadFile("../../webhook/sampleHook.json")
	// r, err := http.NewRequest(http.MethodPost, "", bytes.NewBuffer(raw))
	//
	// mockCtrl := gomock.NewController(t)
	// defer mockCtrl.Finish()
	// mock := database.NewMockDBStorage(mockCtrl)
	// // mock.EXPECT().Add()
}

func Test_subscriptionGet(t *testing.T) {

	type args struct {
		URLpath string
	}
	tests := []struct {
		name        string
		args        args
		wantSub     webhook.SubsciptionOut
		wantSuccess bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSub, gotSuccess := subscriptionGet(tt.args.URLpath)
			if !reflect.DeepEqual(gotSub, tt.wantSub) {
				t.Errorf("subscriptionGet() gotSub = %v, want %v", gotSub, tt.wantSub)
			}
			if gotSuccess != tt.wantSuccess {
				t.Errorf("subscriptionGet() gotSuccess = %v, want %v", gotSuccess, tt.wantSuccess)
			}
		})
	}
}

func Test_subscriptionDelete(t *testing.T) {
	type args struct {
		URLpath string
	}
	tests := []struct {
		name        string
		args        args
		wantSuccess bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSuccess := subscriptionDelete(tt.args.URLpath); gotSuccess != tt.wantSuccess {
				t.Errorf("subscriptionDelete() = %v, want %v", gotSuccess, tt.wantSuccess)
			}
		})
	}
}

var conversionRequest = func(method, path string, body bool) *http.Request {
	if body {
		raw, err := ioutil.ReadFile("../../samples/conversion.json")
		if err != nil {
			log.Fatal("Could not read conversion file " + err.Error())
		}
		rWithBody, err := http.NewRequest(method, path, bytes.NewBuffer(raw))
		if err != nil {
			log.Fatal(err)
		}
		return rWithBody
	}
	rNoBody, err := http.NewRequest(method, path, bytes.NewBufferString(""))
	if err != nil {
		log.Fatal(err)
	}
	return rNoBody
}

func Test_findLastEntry(t *testing.T) {
	latest, err := fixer.GetLatest("")
	if err != nil {
		t.Fatal(err)
	}
	globalDB.DropCollection(dbCurrencyCollection)
	err = globalDB.Add(dbCurrencyCollection, latest) //Make sure there is a "latest entry"

	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name  string
		entry *fixer.Currency
		want  bool
	}{
		{"OK", &latest, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLastEntry(tt.entry); got != tt.want {
				t.Errorf("findLastEntry() = %v, want %v", got, tt.want)
			}
		})
	}

	globalDB.DropCollection(dbCurrencyCollection)
}

func TestLatestHandler(t *testing.T) {

	globalDB.DropCollection(dbWebhookCollection)

	latest, err := fixer.GetLatest("")
	if err != nil {
		t.Fatal(err)
	}
	err = globalDB.Add(dbCurrencyCollection, latest) //Make sure there is a "latest entry"

	if err != nil {
		t.Fatal(err)
	}
	tests := []struct {
		name     string
		r        *http.Request
		wantCode int
	}{
		//Valid
		{"Valid 1", conversionRequest(http.MethodPost, "latest", true), http.StatusOK},
		{"Valid 2", conversionRequest(http.MethodPost, "latest/", true), http.StatusOK},
		//Missing Body
		{"Bad Request, no body", conversionRequest(http.MethodPost, "latest/", false), http.StatusBadRequest},
		//Nonallowed Method
		{"Method not allowed", conversionRequest(http.MethodGet, "latest/", true), http.StatusMethodNotAllowed},
		{"Method not allowed", conversionRequest(http.MethodGet, "latest", true), http.StatusMethodNotAllowed},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rr := httptest.NewRecorder()
			hfunc := http.HandlerFunc(LatestHandler)
			hfunc.ServeHTTP(rr, tt.r)

			Expect("Wrong Status Code", rr.Code, tt.wantCode, t)
		})
	}
	globalDB.DropCollection(dbCurrencyCollection)
}

func Test_findNLatestEntries(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name        string
		args        args
		wantEntries []fixer.Currency
		wantOk      bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotEntries, gotOk := findNLatestEntries(tt.args.n)
			if !reflect.DeepEqual(gotEntries, tt.wantEntries) {
				t.Errorf("findNLatestEntries() gotEntries = %v, want %v", gotEntries, tt.wantEntries)
			}
			if gotOk != tt.wantOk {
				t.Errorf("findNLatestEntries() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func TestAverageHandler(t *testing.T) {

	globalDB.DropCollection(dbWebhookCollection)
	addEntriesForNPastDays(7)

	tests := []struct {
		name     string
		r        *http.Request
		wantCode int
	}{
		//Valid
		{"Valid 1", conversionRequest(http.MethodPost, "average", true), http.StatusOK},
		{"Valid 2", conversionRequest(http.MethodPost, "average/", true), http.StatusOK},
		//Missing Body
		{"Bad Request, no body", conversionRequest(http.MethodPost, "average/", false), http.StatusBadRequest}, //Cant make "NO BODY"
		//Nonallowed Method
		{"Method not allowed", conversionRequest(http.MethodGet, "average/", true), http.StatusMethodNotAllowed},
		{"Method not allowed", conversionRequest(http.MethodGet, "average", true), http.StatusMethodNotAllowed},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rr := httptest.NewRecorder()
			hfunc := http.HandlerFunc(AverageHandler)
			hfunc.ServeHTTP(rr, tt.r)
			// SubscriptionHandler(rr, tt.r)
			Expect("Wrong Status Code", rr.Code, tt.wantCode, t)
		})
	}

	globalDB.DropCollection(dbCurrencyCollection)
}

func TestEvaluationTriggerHandler(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			EvaluationTriggerHandler(tt.args.w, tt.args.r)
		})
	}
}

func Test_initializeDBConnection(t *testing.T) {
	// c := gomock.NewController(t)
	// defer c.Finish()
	// mock := database.NewMockDBStorage(c)
	// mock.EXPECT().CreateSession().Times(1)
}

var subGetRequest = func(path string) *http.Request {
	return subscriptionRequest(http.MethodGet, path, false)
}
var subPostRequest = func(path string, body bool) *http.Request {
	return subscriptionRequest(http.MethodPost, path, body)
}
var subDeleteRequest = func(path string) *http.Request {
	return subscriptionRequest(http.MethodDelete, path, false)
}
var subMethodResuest = func(method string) *http.Request {
	return subscriptionRequest(method, "", false)
}

func TestSubscriptionHandlerMethods(t *testing.T) {

	globalDB.DropCollection(dbWebhookCollection)

	raw, err := ioutil.ReadFile("../../samples/hook.json")
	if err != nil {
		t.Fatal(err)
	}
	sub := &webhook.SubsciptionIn{}
	json.Unmarshal(raw, sub)
	sub.HookID = bson.NewObjectId().Hex()
	globalDB.Add(dbWebhookCollection, sub)

	tests := []struct {
		name     string
		r        *http.Request
		wantCode int
	}{
		//GET
		{"ValidGet", subGetRequest(sub.HookID), http.StatusOK},
		{"NonExisting hook id", subGetRequest("oooooo"), http.StatusNotFound},
		{"Invalid path Get", subGetRequest("/a/sd/"), http.StatusNotFound},
		//POST
		{"ValidPost", subPostRequest("", true), http.StatusCreated},
		{"Missing body", subPostRequest("", false), http.StatusBadRequest}, //Unable to make missing body
		{"Invalid path Post", subPostRequest("/sa/asd", false), http.StatusBadRequest},
		//DELETE
		{"ValidDelete", subDeleteRequest(sub.HookID), http.StatusAccepted},
		{"NonExisting hook id Delete", subDeleteRequest("oooooo"), http.StatusNotFound},
		{"Invalid path Delete", subDeleteRequest("/a/sd/"), http.StatusNotFound},
		//PUT
		{"Invalid method: Put", subMethodResuest(http.MethodPut), http.StatusNotImplemented},
		//HEAD
		{"Invalid method: Head", subMethodResuest(http.MethodHead), http.StatusMethodNotAllowed},
		//PATCH
		{"Invalid method: Patch", subMethodResuest(http.MethodPatch), http.StatusMethodNotAllowed},
		//TRACE
		{"Invalid method: Trace", subMethodResuest(http.MethodTrace), http.StatusMethodNotAllowed},
		//CONNECT
		{"Invalid method: Connect", subMethodResuest(http.MethodConnect), http.StatusMethodNotAllowed},
		//OPTIONS
		{"Invalid method: Options", subMethodResuest(http.MethodOptions), http.StatusMethodNotAllowed},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rr := httptest.NewRecorder()
			hfunc := http.HandlerFunc(SubscriptionHandler)
			hfunc.ServeHTTP(rr, tt.r)
			Expect("Wrong Status Code", rr.Code, tt.wantCode, t)
		})
	}
}
