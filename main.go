package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

const baseURL = "http://api.fixer.io/latest?base=EUR"

//Something - somethis
type Something struct {
	Key   uint16 `json:"key"`
	Value string `json:"value"`
}

//ServiceHandler - TODO
func ServiceHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	s := Something{}
	s.Key = 0
	s.Value = "HOOOOH BOI"
	json.NewEncoder(w).Encode(s)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set. Defaulting to 8085")
		port = "8085"
	}
	http.HandleFunc("/test/", ServiceHandler)
	http.ListenAndServe(":"+port, nil)
}
