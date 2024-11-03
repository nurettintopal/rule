package main

import (
	"fmt"
	"github.com/nurettintopal/rule"
)

func main() {
	input := `{
				"country": "Turkey",
				"city": "Istanbul",
				"location": [40.96959346873711, 29.03704527987249]
			}`

	rules := `{
				"conditions": [{
						"all": [{
								"field": "country",
								"operator": "equals",
								"value": "Turkey"
							},
							{
								"field": "country",
								"operator": "in",
								"value": ["Turkey"]
							},
							{
								"field": "location",
								"operator": "custom.countryEquals",
								"value": "Turkey"
							},
							{
								"field": "location",
								"operator": "custom.countryIn",
								"value": ["Turkey"]
							}
						]
					},
					{
						"all": [{
								"field": "city",
								"operator": "equals",
								"value": "Istanbul"
							},
							{
								"field": "city",
								"operator": "in",
								"value": ["Istanbul"]
							},
							{
								"field": "location",
								"operator": "custom.cityEquals",
								"value": "Istanbul"
							},
							{
								"field": "location",
								"operator": "custom.cityIn",
								"value": ["Istanbul"]
							}
						]
					},
					{
						"all": [{
								"field": "city",
								"operator": "custom.all",
								"value": ["Istanbul"]
							},
							{
								"field": "city",
								"operator": "custom.any",
								"value": ["Istanbul"]
							}
						]
					}
				]
			}`

	custom := map[string]rule.CustomOperation{
		"countryEquals": &CustomRuleCountryEquals{},
		"countryIn":     &CustomRuleCountryIn{},
		"cityEquals":    &CustomRuleCityEquals{},
		"cityIn":        &CustomRuleCityIn{},
		"all":    &CustomRuleAll{},
		"any":    &CustomRuleAny{},
	}

	if rule.Execute(input, rules, custom) {
		fmt.Printf("Rules have been executed. it passed!")
	} else {
		fmt.Printf("Rules have been executed. it failed!")
	}
}
