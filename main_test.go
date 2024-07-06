package main

import "testing"

func TestCompareValues(t *testing.T) {

	if compareValues(5, 5) != 0 {
		t.Errorf("Expected 5 to equal 5")
	}

	if compareValues(4, 5) != -1 {
		t.Errorf("Expected 4 to less than 5")
	}

	if compareValues(6, 5) != 1 {
		t.Errorf("Expected 6 to greaterThan 5")
	}
}

func TestContains(t *testing.T) {
	if contains(1, []int{1, 2, 3}) != true {
		t.Errorf("it includes the input")
	}

	if contains(5, []int{1, 2, 3}) != false {
		t.Errorf("it does not include the input")
	}

	if contains("a", []string{"a", "b", "c"}) != true {
		t.Errorf("it includes the input")
	}

	if contains("d", []string{"a", "b", "c"}) != false {
		t.Errorf("it does not include the input")
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
