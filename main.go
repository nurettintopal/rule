package main

import (
	"encoding/json"
	"fmt"
	"log"
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
	case "==":
		return fieldValue == rule.Value
	case "!=":
		return fieldValue != rule.Value
	case ">":
		return compare(fieldValue, rule.Value) > 0
	case "<":
		return compare(fieldValue, rule.Value) < 0
	case ">=":
		return compare(fieldValue, rule.Value) >= 0
	case "<=":
		return compare(fieldValue, rule.Value) <= 0
	default:
		return false
	}
}

// compare compares two interface{} values
func compare(a, b interface{}) int {
	switch a := a.(type) {
	case float64:
		b := b.(float64)
		if a > b {
			return 1
		} else if a < b {
			return -1
		} else {
			return 0
		}
	case string:
		b := b.(string)
		if a > b {
			return 1
		} else if a < b {
			return -1
		} else {
			return 0
		}
	case int:
		b := b.(int)
		if a > b {
			return 1
		} else if a < b {
			return -1
		} else {
			return 0
		}
	default:
		return 0
	}
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
					{"Field": "country", "Operator": "==", "Value": "Turkey"},
					{"Field": "city", "Operator": "==", "Value": "Istanbul"},
					{"Field": "district", "Operator": "==", "Value": "Kadikoy"},
					{"Field": "population", "Operator": "==", "Value": 20000.00},
					{"Field": "population", "Operator": "<", "Value": 21000.00},
					{"Field": "population", "Operator": ">", "Value": 19000.00}
				],
				"any": [
					{"Field": "country", "Operator": "==", "Value": "England"},
					{"Field": "city", "Operator": "==", "Value": "London"},
					{"Field": "population", "Operator": "==", "Value": 20000.00}
				],
				"any": [
					{"Field": "country", "Operator": "==", "Value": "Turkey"},
					{"Field": "city", "Operator": "==", "Value": "Madrid"},
					{"Field": "population", "Operator": "==", "Value": 1000.00}
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
