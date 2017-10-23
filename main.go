package main

import (
	"log"
	"net/http"
	"os"
)

const baseURL = "http://api.fixer.io/latest?base=EUR"

//ServiceHandler - TODO
func ServiceHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	w.Write([]byte(`{"test":"wat"}`))
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}
	http.HandleFunc("/", ServiceHandler)
	http.ListenAndServe(":"+port, nil)
}
