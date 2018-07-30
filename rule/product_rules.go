package rule

import (
	"GoPoc/models"
	"log"
)


func IsProductDummy(env models.RulesEnv) models.RulesEnv {
	log.Println("In validity check.")
	env.Cart.Result["IsProductDummy"] = env.Cart.IsProductDummy
	if env.Cart.IsProductDummy {
		env.RuleStack.Push(UserIsEligibleForEmployeeOptions)
	}
	return env
}

func UserIsEligibleForEmployeeOptions(env models.RulesEnv)  models.RulesEnv {
	log.Println("In availability check.")
	env.Cart.Result["UserIsEligibleForEmployeeOptions"] = env.Cart.UserIsEligibleForEmployeeOptions
	if env.Cart.UserIsEligibleForEmployeeOptions {
		env.RuleStack.Push(IsProductSigninForPrice)
	}
	return env
}

func IsProductSigninForPrice(env models.RulesEnv) models.RulesEnv {
	log.Println("In product sign for price.")
	env.Cart.Result["IsProductSigninForPrice"] = env.Cart.IsProductSigninForPrice
	if env.Cart.IsProductSigninForPrice {
		env.RuleStack.Push(IsProductCustomerSpecific)
	}
	return env
}

func IsProductCustomerSpecific(env models.RulesEnv) models.RulesEnv {
	log.Println("In customer specific.")
	env.Cart.Result["IsProductCustomerSpecific"] = env.Cart.IsProductCustomerSpecific
	if env.Cart.IsProductCustomerSpecific {
		env.RuleStack.Push(IsAccountWithCustomerSpecificProduct)
	}
	return env
}

func IsAccountWithCustomerSpecificProduct(env models.RulesEnv) models.RulesEnv {
	log.Println("In account with customer specific product.")
	env.Cart.Result["IsAccountWithCustomerSpecificProduct"] = env.Cart.IsAccountWithCustomerSpecificProduct
	return env
}