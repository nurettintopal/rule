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
				   "field":"population",
				   "operator":"greaterThan",
				   "value":4
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

	if rule.Execute(input, rules, nil) {
		fmt.Printf("Rules have been executed. it passed!")
	} else {
		fmt.Printf("Rules have been executed. it failed!")
	}
}
