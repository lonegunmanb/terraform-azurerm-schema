package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermSubscriptions = `{
  "block": {
    "attributes": {
      "display_name_contains": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "display_name_prefix": {
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
      "subscriptions": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "display_name": "string",
              "id": "string",
              "location_placement_id": "string",
              "quota_id": "string",
              "spending_limit": "string",
              "state": "string",
              "subscription_id": "string",
              "tags": [
                "map",
                "string"
              ],
              "tenant_id": "string"
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
  "version": 0
}`

func AzurermSubscriptionsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermSubscriptions), &result)
	return &result
}
