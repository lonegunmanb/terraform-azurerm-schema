package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermSubscription = `{
  "block": {
    "attributes": {
      "alias": {
        "computed": true,
        "description": "The Alias Name of the subscription. If omitted a new UUID will be generated for this property.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "billing_scope_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "subscription_id": {
        "computed": true,
        "description": "The GUID of the Subscription.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "subscription_name": {
        "description": "The Display Name for the Subscription.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "tenant_id": {
        "computed": true,
        "description": "The Tenant ID to which the subscription belongs",
        "description_kind": "plain",
        "type": "string"
      },
      "workload": {
        "description": "The workload type for the Subscription. Possible values are ` + "`" + `Production` + "`" + ` (default) and ` + "`" + `DevTest` + "`" + `.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "timeouts": {
        "block": {
          "attributes": {
            "create": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "delete": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "read": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "update": {
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
  "version": 0
}`

func AzurermSubscriptionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermSubscription), &result)
	return &result
}
