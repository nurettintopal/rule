package main

import (
	"github.com/nurettintopal/rule"
)

// CustomRuleCityIn implementations
type CustomRuleCityIn struct{}

func (o *CustomRuleCityIn) Execute(input interface{}, value interface{}) interface{} {
	//resolve the city from location
	var city = "Istanbul"

	return rule.Contains(city, value)
}