package main

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
)

// Rule represents a single condition
type Rule struct {
	Field    string      `json:"field"`
	Operator string      `json:"operator"`
	Value    interface{} `json:"value"`
}

// ConditionSet represents a set of conditions with All/Any logic
type ConditionSet struct {
	All []Rule `json:"all"`
	Any []Rule `json:"any"`
}

// RuleSet represents the overall rule set with multiple condition sets
type RuleSet struct {
	Conditions []ConditionSet `json:"conditions"`
}

// checkRule checks if a single rule is satisfied by a given JSON object
func checkRule(obj map[string]interface{}, rule Rule) bool {
	fieldValue, exists := obj[rule.Field]
	if !exists {
		return false
	}

	switch rule.Operator {
	case "equals":
		return fieldValue == rule.Value
	case "notEquals":
		return fieldValue != rule.Value
	case "greaterThan":
		return compare(fieldValue, rule.Value) > 0
	case "lessThan":
		return compare(fieldValue, rule.Value) < 0
	case "greaterThanInclusive":
		return compare(fieldValue, rule.Value) >= 0
	case "lessThanInclusive":
		return compare(fieldValue, rule.Value) <= 0
	case "in":
		return contains(fieldValue, rule.Value)
	case "notIn":
		return !contains(fieldValue, rule.Value)
	default:
		return false
	}
}

// contains checks if a value is in an array of either strings or integers
func contains(value, array interface{}) bool {
	arr := reflect.ValueOf(array)

	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < arr.Len(); i++ {
			if reflect.DeepEqual(arr.Index(i).Interface(), value) {
				return true
			}
		}
	}
	return false
}

// compare compares two interface{} values
func compare(a, b interface{}) int {
	switch a := a.(type) {
	case float64:
		b := b.(float64)
		return compareValues(a, b)
	case string:
		b := b.(string)
		return compareValues(a, b)
	case int:
		b := b.(int)
		return compareValues(a, b)
	default:
		return 0
	}
}

// compareValues compares two values of the same type
func compareValues[T float64 | string | int](a, b T) int {
	if a > b {
		return 1
	} else if a < b {
		return -1
	}
	return 0
}

// checkConditionSet checks if a given JSON object satisfies a condition set
func checkConditionSet(obj map[string]interface{}, conditionSet ConditionSet) bool {
	// Check "all" conditions
	for _, rule := range conditionSet.All {
		if !checkRule(obj, rule) {
			return false
		}
	}

	// Check "any" conditions
	anyTrue := false
	for _, rule := range conditionSet.Any {
		if checkRule(obj, rule) {
			anyTrue = true
			break
		}
	}
	if !anyTrue && len(conditionSet.Any) > 0 {
		return false
	}

	return true
}

// checkRuleSet checks if a given JSON object satisfies the entire rule set
func checkRuleSet(obj map[string]interface{}, ruleSet RuleSet) bool {
	for _, conditionSet := range ruleSet.Conditions {
		if !checkConditionSet(obj, conditionSet) {
			return false
		}
	}
	return true
}

func run(input string, rules string) bool {
	var objs map[string]interface{}
	err := json.Unmarshal([]byte(input), &objs)
	if err != nil {
		log.Fatal(err)
	}

	var ruleSet RuleSet
	err = json.Unmarshal([]byte(rules), &ruleSet)
	if err != nil {
		log.Fatal(err)
	}

	if checkRuleSet(objs, ruleSet) {
		return true
	}
	return false
}

func main() {
	input := `{
		"country": "Turkey",
		"city": "Istanbul",
		"district": "Kadikoy",
		"population": 20000.00,
		"language": "Turkish"
	}`

	rules := `{
	   "conditions":[
		  {
			 "all":[
				{
				   "field":"country",
				   "operator":"equals",
				   "value":"Turkey"
				},
				{
				   "field":"city",
				   "operator":"equals",
				   "value":"Istanbul"
				},
				{
				   "field":"district",
				   "operator":"equals",
				   "value":"Kadikoy"
				},
				{
				   "field":"population",
				   "operator":"equals",
				   "value":20000.00
				},
				{
				   "field":"population",
				   "operator":"notEquals",
				   "value":50000.00
				},
				{
				   "field":"population",
				   "operator":"lessThan",
				   "value":21000.00
				},
				{
				   "field":"population",
				   "operator":"lessThanInclusive",
				   "value":20000.00
				},
				{
				   "field":"population",
				   "operator":"greaterThan",
				   "value":19000.00
				},
				{
				   "field":"population",
				   "operator":"greaterThanInclusive",
				   "value":20000.00
				},
				{
				   "field":"country",
				   "operator":"in",
				   "value":[
					  "Turkey"
				   ]
				},
				{
				   "field":"country",
				   "operator":"notIn",
				   "value":[
					  "Germany"
				   ]
				}
			 ],
			 "any":[
				{
				   "field":"country",
				   "operator":"equals",
				   "value":"England"
				},
				{
				   "field":"city",
				   "operator":"equals",
				   "value":"London"
				},
				{
				   "field":"population",
				   "operator":"equals",
				   "value":200000.00
				},
				{
				   "field":"country",
				   "operator":"equals",
				   "value":"Turkey"
				},
				{
				   "field":"city",
				   "operator":"equals",
				   "value":"Madrid"
				},
				{
				   "field":"population",
				   "operator":"equals",
				   "value":1000.00
				}
			 ]
		  }
	   ]
	}`

	if run(input, rules) {
		fmt.Printf("Rules have been executed. it passed!")
	} else {
		fmt.Printf("Rules have been executed. it failed!")
	}
}
