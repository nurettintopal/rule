package main

import (
	"fmt"

	"github.com/nurettintopal/rule"
)

func main() {
	input := `{
		"country": "Turkey",
		"city": "Istanbul",
		"district": "Kadikoy",
		"population": 2000000.00,
		"language": "Turkish"
	}`

	rules := `{
	   "conditions":[
		  {
			 "all":[
				{
				   "field":"external.score",
				   "operator":"greaterThan",
				   "value":4
				},
				{
				   "field":"population",
				   "operator":"custom.eligible",
				   "value": true
				}
			 ],
			 "any":[
				{
				   "field":"country",
				   "operator":"in",
				   "value": ["Turkey", "England"]
				}
			 ]
		  }
	   ]
	}`

	custom := map[string]rule.CustomOperation{
		"score":    &CustomCountryScore{},
		"eligible": &CustomRuleEligible{},
	}

	if rule.Execute(input, rules, custom) {
		fmt.Printf("Rules have been executed. it passed!")
	} else {
		fmt.Printf("Rules have been executed. it failed!")
	}
}

// CustomOperation implementations
type CustomRuleEligible struct{}

func (o *CustomRuleEligible) Execute(input, value interface{}) interface{} {
	fmt.Println("DEBUG: CustomRuleEligible Execute", input, value)
	return true
}

// CustomOperation implementations
type CustomCountryScore struct{}

func (o *CustomCountryScore) Execute(input, value interface{}) interface{} {
	fmt.Println("DEBUG: CustomCountryScore Execute", input, value)
	//TODO: there should be some implementation here.
	return 4.6
}
