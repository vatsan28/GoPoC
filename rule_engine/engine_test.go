package rule_test

import (
	"testing"
	"GoPoc/models"
	"GoPoc/rule"
)

func TestIsProductValid(t *testing.T) {
	validInput := models.Order{IsProductValid: true}
	invalidInput := models.Order{IsProductValid: false}

	validOutput := rule.IsProductValid(validInput)
	invalidOutput := rule.IsProductValid(invalidInput)

	if validOutput == false {
		t.Errorf("handler returned wrong status code.")
	}
	if invalidOutput == true {
		t.Errorf("handler returned wrong status code.")
	}
}