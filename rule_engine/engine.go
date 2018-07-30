package rule_engine

import (
	"GoPoc/models"
	"log"
	"GoPoc/rule"
)

func RunEngine(env models.RulesEnv) models.RulesEnv {
	log.Println("In rule engine.", env)
	env.Cart.Result = make(map[string]bool)
	env.RuleStack.Push(rule.IsProductDummy)

	for env.RuleStack.IsEmpty() == false {
		rule, err := env.RuleStack.Pop()
		log.Println(rule)
		if err != nil {
			log.Println(err)
		}
		env = rule(env)
		log.Println(env)
	}
	log.Println("Done checking rule engine.", env)
	return env
}