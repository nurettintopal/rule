[![nurettintopal - rule](https://img.shields.io/static/v1?label=nurettintopal&message=rule&color=blue&logo=github)](https://github.com/nurettintopal/rule "Go to GitHub repo")
[![stars - rule](https://img.shields.io/github/stars/nurettintopal/rule?style=social)](https://github.com/nurettintopal/rule)
[![forks - rule](https://img.shields.io/github/forks/nurettintopal/rule?style=social)](https://github.com/nurettintopal/rule)
[![GitHub release](https://img.shields.io/github/release/nurettintopal/rule?include_prereleases=&sort=semver&color=blue)](https://github.com/nurettintopal/rule/releases/)
[![License](https://img.shields.io/badge/License-MIT-blue)](#license)
[![issues - rule](https://img.shields.io/github/issues/nurettintopal/rule)](https://github.com/nurettintopal/rule/issues)
<img src="./badge.svg"/>

rule
==============================================
a basic rule engine package in Golang

## what is a rule engine?

> Rules engines or rule expressions serve as pluggable software components which execute business rules that a business rules approach has externalized or separated from application code. This externalization or separation allows business users to modify the rules without the need for a big effort.

if you want to look into the details, follow [this link](https://en.wikipedia.org/wiki/Business_rules_engine), please.

## usage

### basic example
```go
package main

import (
	"fmt"

	"github.com/nurettintopal/rule"
)

func main() {
	input := `{
		"country": "Turkey"
	}`

	rules := `{
	   "conditions":[
		  {
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

	if rule.Execute(input, rules, nil) {
		fmt.Printf("Rules have been executed. it passed!")
	} else {
		fmt.Printf("Rules have been executed. it failed!")
	}
}
```

#### input:
```json
{
  "country": "Turkey",
  "city": "Istanbul",
  "district": "Kadikoy",
  "population": 20000,
  "language": "Turkish"
}
```

#### rules:
```json
{
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
}
```

#### result:
```json
Rules have been executed. it passed!
```


## operators
| operator             | meaning                                              | 
----------------------|------------------------------------------------------
| equals               | equals to                                            |
| notEquals            | not equal to                                         |
| lessThan             | less than                                            |
| greaterThan          | greater than                                         |
| lessThanInclusive    | less than or equal to                                |
| greaterThanInclusive | greater than or equal to                             |
| in                   | in a list                                            |
| notIn                | not in a list                                        |
| startsWith           | starts with                                          |
| endsWith             | ends with                                            |
| contains             | contains                                             |
| notContains          | not contains                                         |
| regex                | contains any match of the regular expression pattern |


## how to add custom operator
it has been already supporting a few rules that can be used in your projects, but sometimes, you may need to use custom controls based on your own business rules. do not worry, if you need to add some additional control, you can do it easly. you need to create your own function, then, inject the function to the package, that is all.

```go
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
				   "field":"population",
				   "operator":"custom.eligible",
				   "value": true
				}
			 ]
		  }
	   ]
	}`

custom := map[string]rule.CustomOperation{
    "eligible": &CustomRuleEligible{},
}

if rule.Execute(input, rules, custom) {
    fmt.Printf("Rules have been executed. it passed!")
} else {
    fmt.Printf("Rules have been executed. it failed!")
}


// CustomOperation implementations
type CustomRuleEligible struct{}

func (o *CustomRuleEligible) Execute(input, value interface{}) interface{} {
	// your own logics & controls
	return true
}
```

## how to add custom input
you need to send inputs that should be used for rules, but sometimes, you may need to use custom sources based on your own business rules. do not worry, if you need to add some additional sources for inputs, you can do it easly. you need to create your own function, then, inject the function to the package, that is all.

```go
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
				   "field":"external.score",
				   "operator":"greaterThan",
				   "value":4
				}
			 ]
		  }
	   ]
	}`

custom := map[string]rule.CustomOperation{
    "score":    &CustomCountryScore{},
}

if rule.Execute(input, rules, custom) {
    fmt.Printf("Rules have been executed. it passed!")
} else {
    fmt.Printf("Rules have been executed. it failed!")
}


// CustomOperation implementations
type CustomCountryScore struct{}

func (o *CustomCountryScore) Execute(input, value interface{}) interface{} {
    // your own logics & controls
    return "your value that is retrieved from a different source"
}
```



## dependencies
* Go

## contributing
* if you want to add anything, contributions are welcome.
* Open a pull request that has your explanations

## license
leaderboard is open-sourced software licensed under the [MIT license](LICENSE).
