package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermConsumptionBudgetSubscription = `{
  "block": {
    "attributes": {
      "amount": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "filter": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "dimension": [
                "list",
                [
                  "object",
                  {
                    "name": "string",
                    "operator": "string",
                    "values": [
                      "list",
                      "string"
                    ]
                  }
                ]
              ],
              "not": [
                "list",
                [
                  "object",
                  {
                    "dimension": [
                      "list",
                      [
                        "object",
                        {
                          "name": "string",
                          "operator": "string",
                          "values": [
                            "list",
                            "string"
                          ]
                        }
                      ]
                    ],
                    "tag": [
                      "list",
                      [
                        "object",
                        {
                          "name": "string",
                          "operator": "string",
                          "values": [
                            "list",
                            "string"
                          ]
                        }
                      ]
                    ]
                  }
                ]
              ],
              "tag": [
                "list",
                [
                  "object",
                  {
                    "name": "string",
                    "operator": "string",
                    "values": [
                      "list",
                      "string"
                    ]
                  }
                ]
              ]
            }
          ]
        ]
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "notification": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "contact_emails": [
                "list",
                "string"
              ],
              "contact_groups": [
                "list",
                "string"
              ],
              "contact_roles": [
                "list",
                "string"
              ],
              "enabled": "bool",
              "operator": "string",
              "threshold": "number",
              "threshold_type": "string"
            }
          ]
        ]
      },
      "subscription_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "time_grain": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "time_period": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "end_date": "string",
              "start_date": "string"
            }
          ]
        ]
      }
    },
    "block_types": {
      "timeouts": {
        "block": {
          "attributes": {
            "read": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "single"
      }
    },
    "description_kind": "plain"
  },
  "version": 2
}`

func AzurermConsumptionBudgetSubscriptionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermConsumptionBudgetSubscription), &result)
	return &result
}
