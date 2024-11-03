package main

// CustomRuleCountryEquals implementations
type CustomRuleCountryEquals struct{}

func (o *CustomRuleCountryEquals) Execute(input interface{}, value interface{}) interface{} {
	//resolve the country from location
	var country = "Turkey"

	return country == value.(string)
}