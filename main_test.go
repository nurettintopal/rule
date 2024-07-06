package main

import (
	"testing"
)

func TestEqualsOperator(t *testing.T) {
	op := EqualsOperator{}
	if !op.Apply(5, 5) {
		t.Errorf("Expected 5 equals 5")
	}
	if op.Apply(5, 10) {
		t.Errorf("Expected 5 not equals 10")
	}
}

func TestNotEqualsOperator(t *testing.T) {
	op := NotEqualsOperator{}
	if !op.Apply(5, 10) {
		t.Errorf("Expected 5 not equals 10")
	}
	if op.Apply(5, 5) {
		t.Errorf("Expected 5 equals 5")
	}
}

func TestGreaterThanOperator(t *testing.T) {
	op := GreaterThanOperator{}
	if !op.Apply(10, 5) {
		t.Errorf("Expected 10 greater than 5")
	}
	if op.Apply(5, 10) {
		t.Errorf("Expected 5 not greater than 10")
	}
}

func TestLessThanOperator(t *testing.T) {
	op := LessThanOperator{}
	if !op.Apply(5, 10) {
		t.Errorf("Expected 5 less than 10")
	}
	if op.Apply(10, 5) {
		t.Errorf("Expected 10 not less than 5")
	}
}

func TestGreaterThanInclusiveOperator(t *testing.T) {
	op := GreaterThanInclusiveOperator{}
	if !op.Apply(10, 5) {
		t.Errorf("Expected 10 greater than or equals 5")
	}
	if !op.Apply(5, 5) {
		t.Errorf("Expected 5 greater than or equals 5")
	}
	if op.Apply(5, 10) {
		t.Errorf("Expected 5 not greater than or equals 10")
	}
}

func TestLessThanInclusiveOperator(t *testing.T) {
	op := LessThanInclusiveOperator{}
	if !op.Apply(5, 10) {
		t.Errorf("Expected 5 less than or equals 10")
	}
	if !op.Apply(5, 5) {
		t.Errorf("Expected 5 less than or equals 5")
	}
	if op.Apply(10, 5) {
		t.Errorf("Expected 10 not less than or equals 5")
	}
}

func TestInOperator(t *testing.T) {
	op := InOperator{}
	if !op.Apply("Turkey", []string{"Turkey", "Germany"}) {
		t.Errorf("Expected 'Turkey' in ['Turkey', 'Germany']")
	}
	if op.Apply("France", []string{"Turkey", "Germany"}) {
		t.Errorf("Expected 'France' not in ['Turkey', 'Germany']")
	}
}

func TestNotInOperator(t *testing.T) {
	op := NotInOperator{}
	if !op.Apply("France", []string{"Turkey", "Germany"}) {
		t.Errorf("Expected 'France' not in ['Turkey', 'Germany']")
	}
	if op.Apply("Turkey", []string{"Turkey", "Germany"}) {
		t.Errorf("Expected 'Turkey' in ['Turkey', 'Germany']")
	}
}

func TestCheckRule(t *testing.T) {
	obj := map[string]interface{}{
		"country": "Turkey",
		"age":     30,
	}

	ruleChecker := RuleChecker{OperatorFactory: OperatorFactory{}}

	tests := []struct {
		rule     Rule
		expected bool
	}{
		{Rule{"country", "equals", "Turkey"}, true},
		{Rule{"country", "notEquals", "Germany"}, true},
		{Rule{"age", "greaterThan", 25}, true},
		{Rule{"age", "lessThan", 35}, true},
		{Rule{"country", "in", []string{"Turkey", "Germany"}}, true},
		{Rule{"country", "notIn", []string{"France", "Italy"}}, true},
	}

	for _, test := range tests {
		result := ruleChecker.CheckRule(obj, test.rule)
		if result != test.expected {
			t.Errorf("CheckRule(%v, %v) = %v; expected %v", obj, test.rule, result, test.expected)
		}
	}
}

func TestCheckConditionSet(t *testing.T) {
	obj := map[string]interface{}{
		"country": "Turkey",
		"age":     30,
	}

	ruleChecker := RuleChecker{OperatorFactory: OperatorFactory{}}
	conditionSetChecker := ConditionSetChecker{RuleChecker: ruleChecker}

	tests := []struct {
		conditionSet ConditionSet
		expected     bool
	}{
		{ConditionSet{All: []Rule{{"country", "equals", "Turkey"}}, Any: []Rule{}}, true},
		{ConditionSet{All: []Rule{{"country", "equals", "Germany"}}, Any: []Rule{}}, false},
		{ConditionSet{All: []Rule{}, Any: []Rule{{"country", "equals", "Germany"}, {"country", "equals", "Turkey"}}}, true},
		{ConditionSet{All: []Rule{}, Any: []Rule{{"country", "equals", "France"}}}, false},
	}

	for _, test := range tests {
		result := conditionSetChecker.CheckConditionSet(obj, test.conditionSet)
		if result != test.expected {
			t.Errorf("CheckConditionSet(%v, %v) = %v; expected %v", obj, test.conditionSet, result, test.expected)
		}
	}
}

func TestCheckRuleSet(t *testing.T) {
	obj := map[string]interface{}{
		"country": "Turkey",
		"age":     30,
	}

	ruleChecker := RuleChecker{OperatorFactory: OperatorFactory{}}
	conditionSetChecker := ConditionSetChecker{RuleChecker: ruleChecker}
	ruleSetChecker := RuleSetChecker{ConditionSetChecker: conditionSetChecker}

	tests := []struct {
		ruleSet  RuleSet
		expected bool
	}{
		{RuleSet{Conditions: []ConditionSet{{All: []Rule{{"country", "equals", "Turkey"}}, Any: []Rule{}}}}, true},
		{RuleSet{Conditions: []ConditionSet{{All: []Rule{{"country", "equals", "Germany"}}, Any: []Rule{}}}}, false},
		{RuleSet{Conditions: []ConditionSet{{All: []Rule{}, Any: []Rule{{"country", "equals", "Germany"}, {"country", "equals", "Turkey"}}}}}, true},
		{RuleSet{Conditions: []ConditionSet{{All: []Rule{}, Any: []Rule{{"country", "equals", "France"}}}}}, false},
	}

	for _, test := range tests {
		result := ruleSetChecker.CheckRuleSet(obj, test.ruleSet)
		if result != test.expected {
			t.Errorf("CheckRuleSet(%v, %v) = %v; expected %v", obj, test.ruleSet, result, test.expected)
		}
	}
}

func TestRun(t *testing.T) {
	input := `{
	   "country":"Türkiye",
	   "city":"İstanbul",
	   "district":"Kadıköy"
	}`

	rules := `{
	   "conditions":[
		  {
			 "all":[
				{
				   "field":"country",
				   "operator":"equals",
				   "value":"Türkiye"
				},
				{
				   "field":"city",
				   "operator":"equals",
				   "value":"İstanbul"
				}
			 ]
		  }
	   ]
	}`

	if run(input, rules) != true {
		t.Errorf("it is not passed")
	}
}
