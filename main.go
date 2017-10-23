package main

import (
	"net/http"
)

const baseURL = "http://api.fixer.io/latest?base=EUR"

//ServiceHandler - TODO
func ServiceHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	w.Write([]byte(`{"test":"wat"}`))
}

func main() {
	http.HandleFunc("/", ServiceHandler)
}
