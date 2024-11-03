package main

// CustomRuleCityEquals implementations
type CustomRuleCityEquals struct{}

func (o *CustomRuleCityEquals) Execute(input interface{}, value interface{}) interface{} {
	//resolve the city from location
	var city = "Istanbul"

	return city == value.(string)
}