package rule

import (
	"net/http"
	"log"
	"encoding/json"
)

//This is the expected payload mapping.
type Order struct {
	IsProductValid bool `json:"isProductValid"`
}

// This struct defines the engine's response type.
type ResponseData struct {
	Status int
	Message string
}

func Isvalid(w http.ResponseWriter, r *http.Request) {
	var payload Order
	err := json.NewDecoder(r.Body).Decode(&payload)
	var response = ResponseData{}
	if err != nil {
		log.Println(err)
		response = prepareResponse(422,"Error with the body in the request.")
	}

	if payload.IsProductValid == true {
		response = prepareResponse(200, "All good.")
	} else {
		response = prepareResponse(422,"Product validity failed.")
	}
	json.NewEncoder(w).Encode(response)
	log.Println("In rule engine check.")
}

func prepareResponse(status int, message string) ResponseData {
	return ResponseData{Status: status, Message: message}
}