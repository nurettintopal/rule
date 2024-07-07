package main

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"regexp"
	"strings"
)

// Operator defines an interface for all operators
type Operator interface {
	Apply(fieldValue, ruleValue interface{}) bool
}

// EqualsOperator checks if fieldValue equals ruleValue
type EqualsOperator struct{}

func (o EqualsOperator) Apply(fieldValue, ruleValue interface{}) bool {
	return reflect.DeepEqual(fieldValue, ruleValue)
}

// NotEqualsOperator checks if fieldValue not equals ruleValue
type NotEqualsOperator struct{}

func (o NotEqualsOperator) Apply(fieldValue, ruleValue interface{}) bool {
	return !reflect.DeepEqual(fieldValue, ruleValue)
}

// GreaterThanOperator checks if fieldValue is greater than ruleValue
type GreaterThanOperator struct{}

func (o GreaterThanOperator) Apply(fieldValue, ruleValue interface{}) bool {
	return compare(fieldValue, ruleValue) > 0
}

// LessThanOperator checks if fieldValue is less than ruleValue
type LessThanOperator struct{}

func (o LessThanOperator) Apply(fieldValue, ruleValue interface{}) bool {
	return compare(fieldValue, ruleValue) < 0
}

// GreaterThanInclusiveOperator checks if fieldValue is greater than or equals to ruleValue
type GreaterThanInclusiveOperator struct{}

func (o GreaterThanInclusiveOperator) Apply(fieldValue, ruleValue interface{}) bool {
	return compare(fieldValue, ruleValue) >= 0
}

// LessThanInclusiveOperator checks if fieldValue is less than or equals to ruleValue
type LessThanInclusiveOperator struct{}

func (o LessThanInclusiveOperator) Apply(fieldValue, ruleValue interface{}) bool {
	return compare(fieldValue, ruleValue) <= 0
}

// InOperator checks if fieldValue is in ruleValue array
type InOperator struct{}

func (o InOperator) Apply(fieldValue, ruleValue interface{}) bool {
	return contains(fieldValue, ruleValue)
}

// NotInOperator checks if fieldValue is not in ruleValue array
type NotInOperator struct{}

func (o NotInOperator) Apply(fieldValue, ruleValue interface{}) bool {
	return !contains(fieldValue, ruleValue)
}

// StartsWithOperator checks if fieldValue starts with ruleValue
type StartsWithOperator struct{}

func (o StartsWithOperator) Apply(fieldValue, ruleValue interface{}) bool {
	return strings.HasPrefix(fieldValue.(string), ruleValue.(string))
}

// EndsWithOperator checks if fieldValue ends with ruleValue
type EndsWithOperator struct{}

func (o EndsWithOperator) Apply(fieldValue, ruleValue interface{}) bool {
	return strings.HasSuffix(fieldValue.(string), ruleValue.(string))
}

// ContainsOperator checks if fieldValue contains ruleValue
type ContainsOperator struct{}

func (o ContainsOperator) Apply(fieldValue, ruleValue interface{}) bool {
	return strings.Contains(fieldValue.(string), ruleValue.(string))
}

// NotContainsOperator checks if fieldValue does not contains ruleValue
type NotContainsOperator struct{}

func (o NotContainsOperator) Apply(fieldValue, ruleValue interface{}) bool {
	return !strings.Contains(fieldValue.(string), ruleValue.(string))
}

// RegexOperator checks if fieldValue contains any match of the regular expression pattern
type RegexOperator struct{}

func (o RegexOperator) Apply(fieldValue, ruleValue interface{}) bool {
	result, _ := regexp.MatchString(ruleValue.(string), fieldValue.(string))
	return result
}

// OperatorFactory to create operators based on string representation
type OperatorFactory struct{}

func (f OperatorFactory) Create(operator string) Operator {
	switch operator {
	case "equals":
		return EqualsOperator{}
	case "notEquals":
		return NotEqualsOperator{}
	case "greaterThan":
		return GreaterThanOperator{}
	case "lessThan":
		return LessThanOperator{}
	case "greaterThanInclusive":
		return GreaterThanInclusiveOperator{}
	case "lessThanInclusive":
		return LessThanInclusiveOperator{}
	case "in":
		return InOperator{}
	case "notIn":
		return NotInOperator{}
	case "startsWith":
		return StartsWithOperator{}
	case "endsWith":
		return EndsWithOperator{}
	case "contains":
		return ContainsOperator{}
	case "notContains":
		return NotContainsOperator{}
	case "regex":
		return RegexOperator{}
	default:
		return nil
	}
}

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

// RuleChecker checks rules against an object
type RuleChecker struct {
	OperatorFactory OperatorFactory
}

func (rc RuleChecker) CheckRule(obj map[string]interface{}, rule Rule) bool {
	fieldValue, exists := obj[rule.Field]
	if !exists {
		return false
	}

	operator := rc.OperatorFactory.Create(rule.Operator)
	if operator == nil {
		return false
	}
	return operator.Apply(fieldValue, rule.Value)
}

// ConditionSetChecker checks condition sets against an object
type ConditionSetChecker struct {
	RuleChecker RuleChecker
}

func (cc ConditionSetChecker) CheckConditionSet(obj map[string]interface{}, conditionSet ConditionSet) bool {
	// Check "all" conditions
	for _, rule := range conditionSet.All {
		if !cc.RuleChecker.CheckRule(obj, rule) {
			return false
		}
	}

	// Check "any" conditions
	if len(conditionSet.Any) > 0 {
		for _, rule := range conditionSet.Any {
			if cc.RuleChecker.CheckRule(obj, rule) {
				return true
			}
		}
		return false
	}

	return true
}

// RuleSetChecker checks rule sets against an object
type RuleSetChecker struct {
	ConditionSetChecker ConditionSetChecker
}

func (rsc RuleSetChecker) CheckRuleSet(obj map[string]interface{}, ruleSet RuleSet) bool {
	for _, conditionSet := range ruleSet.Conditions {
		if !rsc.ConditionSetChecker.CheckConditionSet(obj, conditionSet) {
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

	operatorFactory := OperatorFactory{}
	ruleChecker := RuleChecker{OperatorFactory: operatorFactory}
	conditionSetChecker := ConditionSetChecker{RuleChecker: ruleChecker}
	ruleSetChecker := RuleSetChecker{ConditionSetChecker: conditionSetChecker}

	return ruleSetChecker.CheckRuleSet(objs, ruleSet)
}

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
				   "field":"country",
				   "operator":"equals",
				   "value":"Turkey"
				},
				{
				   "field":"population",
				   "operator":"lessThan",
				   "value": 2000000.01
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

	if run(input, rules) {
		fmt.Printf("Rules have been executed. it passed!")
	} else {
		fmt.Printf("Rules have been executed. it failed!")
	}
}
