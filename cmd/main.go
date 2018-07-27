package main

import (
	"net/http"
	"log"
	"encoding/json"
	"GoPoc/models"
	"GoPoc/rule_engine"
)

const (
	port = "8000"
)


func main() {
	http.HandleFunc("/", validate)
	http.ListenAndServe(":"+port, nil)
}

func validate(w http.ResponseWriter, r *http.Request) {
	cartPayload := new(models.Cart)
	err := json.NewDecoder(r.Body).Decode(&cartPayload)
	var response = models.ResponseData{}
	if err != nil {
		log.Println(err)
		response = prepareResponse(422,"Error with the body in the request.", *cartPayload)
	}

	log.Println(cartPayload)
	var env models.RulesEnv
	env.Cart = *cartPayload
	log.Println(env)
	result := rule_engine.RunEngine(env)

	response = prepareResponse(200,"All ok.", result.Cart)
	json.NewEncoder(w).Encode(response)
}

func prepareResponse(status int, message string, resultEnv models.Cart) models.ResponseData {
	return models.ResponseData{Status: status, Message: message, Cart: resultEnv}
}