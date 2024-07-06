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

	if !op.Apply(5.5, 5.5) {
		t.Errorf("Expected 5.5 equals 5.5")
	}
	if op.Apply(5.5, 10.5) {
		t.Errorf("Expected 5.5 not equals 10.5")
	}

	if !op.Apply("a", "a") {
		t.Errorf("Expected a equals a")
	}
	if op.Apply("a", "b") {
		t.Errorf("Expected a not equals b")
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

	if !op.Apply(5.5, 10.5) {
		t.Errorf("Expected 5.5 not equals 10.5")
	}
	if op.Apply(5.5, 5.5) {
		t.Errorf("Expected 5.5 equals 5.5")
	}

	if !op.Apply("a", "b") {
		t.Errorf("Expected a not equals b")
	}
	if op.Apply("a", "a") {
		t.Errorf("Expected a equals a")
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
	if !op.Apply(100, []int{100, 200}) {
		t.Errorf("Expected '100' in [100, 200]")
	}
	if op.Apply(300, []int{100, 200}) {
		t.Errorf("Expected '300' not in [100, 200]")
	}
	if !op.Apply(100.1, []float64{100.1, 200.1}) {
		t.Errorf("Expected '100.1' in [100.1, 200.1]")
	}
	if op.Apply(300.1, []float64{100.1, 200.1}) {
		t.Errorf("Expected '300.1' not in [100.1, 200.1]")
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

func TestStartsWithOperator(t *testing.T) {
	op := StartsWithOperator{}
	if !op.Apply("Turkey", "Tur") {
		t.Errorf("Expected Turkey not starts with Tur")
	}
}

func TestEndsWithOperator(t *testing.T) {
	op := EndsWithOperator{}
	if !op.Apply("Turkey", "key") {
		t.Errorf("Expected Turkey not ends with key")
	}
}

func TestContainsOperator(t *testing.T) {
	op := ContainsOperator{}
	if !op.Apply("Turkey", "urke") {
		t.Errorf("Expected Turkey not contains urke")
	}
}

func TestRegexOperator(t *testing.T) {
	op := RegexOperator{}
	if !op.Apply("New York City", "[A-z]ork") {
		t.Errorf("Expected New York City not contains regex")
	}

	if op.Apply("New York City", "[A-z]tanbul") {
		t.Errorf("Expected New York City contains regex")
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

	input2 := `{
	   "country":"Korea",
	   "city":"Seul",
	   "district":"Samsung"
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

	rules2 := `{
	   "conditions":[
		  {
			 "all":[
				{
				   "field":"country",
				   "operator":"notAnExistingRule",
				   "value":"Türkiye"
				}
			 ]
		  }
	   ]
	}`

	if run(input, rules) != true {
		t.Errorf("it is not passed")
	}

	if run(input2, rules) == true {
		t.Errorf("it is passed")
	}

	if run(input2, rules2) == true {
		t.Errorf("it is passed")
	}
}
