package main

import (
	"github.com/nurettintopal/rule"
)

// CustomRuleCountryIn implementations
type CustomRuleCountryIn struct{}

func (o *CustomRuleCountryIn) Execute(input interface{}, value interface{}) interface{} {
	//resolve the country from location
	var country = "Turkey"

	return rule.Contains(country, value)
}