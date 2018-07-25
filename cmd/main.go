package main

import (
	"net/http"
	"log"
	"encoding/json"
	"GoPoc/models"
	"GoPoc/rule"
)

const (
	port = "8000"
)

func main() {
	http.HandleFunc("/", checkValidity)
	http.ListenAndServe(":"+port, nil)
}

func checkValidity(w http.ResponseWriter, r *http.Request) {
	var cart models.Order
	err := json.NewDecoder(r.Body).Decode(&cart)
	var response = models.ResponseData{}
	if err != nil {
		log.Println(err)
		response = prepareResponse(422,"Error with the body in the request.")
	}

	if rule.IsProductValid(cart) == true {
		response = prepareResponse(200, "All good.")
	} else {
		response = prepareResponse(422,"Product validity failed.")
	}
	json.NewEncoder(w).Encode(response)
}

func prepareResponse(status int, message string) models.ResponseData {
	return models.ResponseData{Status: status, Message: message}
}