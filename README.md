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


## features
* Multiple rules by all or any

## dependencies
* Go

## contributing
* if you want to add anything, contributions are welcome.
* Open a pull request that has your explanations

## license
leaderboard is open-sourced software licensed under the [MIT license](LICENSE).
