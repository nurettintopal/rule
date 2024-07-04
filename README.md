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
        {
          "Field": "country",
          "Operator": "==",
          "Value": "Turkey"
        },
        {
          "Field": "city",
          "Operator": "==",
          "Value": "Istanbul"
        },
        {
          "Field": "district",
          "Operator": "==",
          "Value": "Kadikoy"
        },
        {
          "Field": "population",
          "Operator": "==",
          "Value": 20000
        },
        {
          "Field": "population",
          "Operator": "<",
          "Value": 21000
        },
        {
          "Field": "population",
          "Operator": ">",
          "Value": 19000
        }
      ],
      "any": [
        {
          "Field": "country",
          "Operator": "==",
          "Value": "Turkey"
        },
        {
          "Field": "city",
          "Operator": "==",
          "Value": "Madrid"
        },
        {
          "Field": "population",
          "Operator": "==",
          "Value": 1000
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


## features
* Multiple rules by all or any

## dependencies
* Go

## contributing
* if you want to add anything, contributions are welcome.
* Open a pull request that has your explanations

## license
leaderboard is open-sourced software licensed under the [MIT license](LICENSE).
