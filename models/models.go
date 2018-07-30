package models

import "errors"

type Cart struct {
	IsProductValid bool `json:"isProductValid"`
	IsProductAvailable bool `json:"isProductAvailable"`
	Result map[string]bool
}

type RuleStack []func(RulesEnv) RulesEnv

type RulesEnv struct {
	Cart Cart
	RuleStack RuleStack
}

type ResponseData struct {
	Message string
	Status int
	Cart Cart
}

func (s *RuleStack) Push(newRule func(RulesEnv) RulesEnv) {
	*s = append(*s, newRule)
}

func (s *RuleStack) Pop() (func(RulesEnv) RulesEnv, error) {
	length := len(*s)

	if length == 0 {
		return nil, errors.New("Empty Rule Stack")
	}

	res := (*s)[0]
	*s = (*s)[1:]
	return res, nil
}

func (s *RuleStack) IsEmpty() bool {
	return len(*s) == 0
}



