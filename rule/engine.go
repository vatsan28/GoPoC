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
	var order Order
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		log.Println(err)
	}
	log.Println(order.IsProductValid)
	var response = ResponseData{Status: 200, Message: "Ok"}
	json.NewEncoder(w).Encode(response)
	log.Println("In rule engine check.")
}