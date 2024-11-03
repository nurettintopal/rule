package main

import (
	"fmt"
)

// CustomRuleAny implementations
type CustomRuleAny struct{}

func (o *CustomRuleAny) Execute(input interface{}, value interface{}) interface{} {
	fmt.Println("CustomRuleAny: ", input, value)
	return true
}