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
		response.Message = "Error with the body in the request."
		response.Status = 422
	}

	if payload.IsProductValid == true {
		response.Message = "All good."
		response.Status = 200
	} else {
		response.Message = "Product validity failed."
		response.Status = 422
	}

	json.NewEncoder(w).Encode(response)
	log.Println("In rule engine check.")
}