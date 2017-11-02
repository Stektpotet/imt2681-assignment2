package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stektpotet/imt2681-assignment2/database"
	"github.com/stektpotet/imt2681-assignment2/util"
)

//
// func TestIInit(t *testing.T) {
// 	mockctrl := gomock.NewController(t)
// 	defer mockctrl.Finish()
// 	mock := NewMockDBService(mockctrl)
// 	initCall := mock.EXPECT().init().Times(1)
// 	mock.EXPECT().Insert("Something").Times(2).After(initCall)
// 	mock.EXPECT().Count().AnyTimes().After(initCall)
//
// 	//Invoke test
// 	Invoke(mock)
// }

var testDB database.DBStorage

func TestMain(m *testing.M) {
	var mongoDBHosts = []string{
		"cluster0-shard-00-00-qvogu.mongodb.net:27017",
		"cluster0-shard-00-01-qvogu.mongodb.net:27017",
		"cluster0-shard-00-02-qvogu.mongodb.net:27017",
	}

	globalDB = &database.CurrencyMongoDB{
		MongoDB: &database.MongoDB{
			HostURLs:  mongoDBHosts,
			AdminUser: util.GetEnv("TRACKER_USER"),
			AdminPass: util.GetEnv("TRACKER_PASS"),
			Name:      "testing",
		},
	}
}

var requestOnHandlerCode = func(method, path string, body []byte, handler http.HandlerFunc) int {
	request, _ := http.NewRequest(method, rootPath+path, bytes.NewBuffer(body))
	rr := httptest.NewRecorder()
	hfunc := http.HandlerFunc(handler)
	hfunc.ServeHTTP(rr, request)
	return rr.Code
}

func Expect(expectString string, got, expected interface{}, t *testing.T) {
	if got != expected {
		t.Errorf("%s: got %v want %v", expectString, got, expected)
	}
}

func TestSubscriptionPOST(t *testing.T) {

}
func TestSubscriptionGET(t *testing.T) {
}
func TestSubscriptionPUT(t *testing.T) {
}
func TestSubscriptionDELETE(t *testing.T) {
}
func TestSubscriptionFULL(t *testing.T) {
}
func TestUpdateCurrencies(t *testing.T) {
	// var testDB database.DBStorage
	// testDB = &database.CurrencyDB{}
	// testDB.Init()
	//
	// p := fixer.GetCurrencies("2017-10-25")
	// testDB.Add(p)
	// v, _ := testDB.Get(p.Date)
	// log.Printf("Something: %+v\n\n\n", v)
	//
	// Tick(testDB)
	// v, _ = testDB.Get("2017-10-27")
	// log.Printf("Something else: %+v\n\n\n", v)
}

func TestUpdateCurrenciesMongo(t *testing.T) {

}
