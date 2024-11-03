package main

import (
	"fmt"
)

// CustomRuleAll implementations
type CustomRuleAll struct{}

func (o *CustomRuleAll) Execute(input interface{}, value interface{}) interface{} {
	fmt.Println("CustomRuleAll: ", input, value)
	return true
}