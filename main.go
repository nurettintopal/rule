package main

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
)

// Rule represents a single condition
type Rule struct {
	Field    string      `json:"Field"`
	Operator string      `json:"Operator"`
	Value    interface{} `json:"Value"`
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

func main() {
	jsonData := `{
		"country": "Turkey",
		"city": "Istanbul",
		"district": "Kadikoy",
		"population": 20000.00,
		"language": "Turkish"
	}`
	var objs map[string]interface{}
	err := json.Unmarshal([]byte(jsonData), &objs)
	if err != nil {
		log.Fatal(err)
	}

	ruleData := `{
		"conditions": [
			{
				"all": [
					{"Field": "country", "Operator": "equals", "Value": "Turkey"},
					{"Field": "city", "Operator": "equals", "Value": "Istanbul"},
					{"Field": "district", "Operator": "equals", "Value": "Kadikoy"},
					{"Field": "population", "Operator": "equals", "Value": 20000.00},
					{"Field": "population", "Operator": "notEquals", "Value": 50000.00},
					{"Field": "population", "Operator": "lessThan", "Value": 21000.00},
					{"Field": "population", "Operator": "lessThanInclusive", "Value": 20000.00},
					{"Field": "population", "Operator": "greaterThan", "Value": 19000.00},
					{"Field": "population", "Operator": "greaterThanInclusive", "Value": 20000.00},
					{"Field": "country", "Operator": "in", "Value": ["Turkey"]},
					{"Field": "country", "Operator": "notIn", "Value": ["Germany"]}
				],
				"any": [
					{"Field": "country", "Operator": "equals", "Value": "England"},
					{"Field": "city", "Operator": "equals", "Value": "London"},
					{"Field": "population", "Operator": "equals", "Value": 200000.00},
					{"Field": "country", "Operator": "equals", "Value": "Turkey"},
					{"Field": "city", "Operator": "equals", "Value": "Madrid"},
					{"Field": "population", "Operator": "equals", "Value": 1000.00}
				]
			}
		]
	}`

	var ruleSet RuleSet
	err = json.Unmarshal([]byte(ruleData), &ruleSet)
	if err != nil {
		log.Fatal(err)
	}

	if checkRuleSet(objs, ruleSet) {
		fmt.Printf("Rules have been executed. it passed!")
		fmt.Printf("")
	} else {
		fmt.Printf("Rules have been executed. it failed!")
		fmt.Printf("")
	}
}
