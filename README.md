rule-engine
==============================================
a basic rule engine implementation in Golang

## what is a rule engine?

> Rules engines or rule expressions serve as pluggable software components which execute business rules that a business rules approach has externalized or separated from application code. This externalization or separation allows business users to modify the rules without the need for a big effort.

if you want to look into the details, follow [this link](https://en.wikipedia.org/wiki/Business_rules_engine), please.

## usage
* TBD
* TBD

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
}
```

#### result:
```json
Rules have been executed. it passed!
```


## features
* Multiple rules by all or any

## dependencies
* Go

## contributing
* if you want to add anything, contributions are welcome.
* Open a pull request that has your explanations

## license
leaderboard is open-sourced software licensed under the [MIT license](LICENSE).
