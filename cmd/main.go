package main

import (
	"net/http"
	"log"
	"GoPoc/rule"
)

const (
	port = "8000"
)

func main() {

	http.HandleFunc("/", rule.Isvalid)
	http.ListenAndServe(":"+port, nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	log.Println("Hello world. 1.0")
}