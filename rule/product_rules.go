package rule

import (
	"GoPoc/models"
	"log"
)


func IsValidProduct(env models.RulesEnv) models.RulesEnv {
	log.Println("In validity check.")
	env.Cart.Result["IsValidProduct"] = env.Cart.IsProductValid
	return env
}

func IsProductAvailable(env models.RulesEnv)  models.RulesEnv {
	log.Println("In availability check.")
	env.Cart.Result["IsProductAvailable"] = env.Cart.IsProductAvailable
	return env
}