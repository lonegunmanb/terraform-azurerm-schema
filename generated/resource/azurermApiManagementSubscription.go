package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermApiManagementSubscription = `{
  "block": {
    "attributes": {
      "allow_tracing": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "api_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "api_management_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "display_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "primary_key": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "sensitive": true,
        "type": "string"
      },
      "product_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "resource_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "secondary_key": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "sensitive": true,
        "type": "string"
      },
      "state": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "subscription_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "user_id": {
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

func AzurermApiManagementSubscriptionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermApiManagementSubscription), &result)
	return &result
}
