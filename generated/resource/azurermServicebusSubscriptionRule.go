package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermServicebusSubscriptionRule = `{
  "block": {
    "attributes": {
      "action": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "filter_type": {
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
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "sql_filter": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "sql_filter_compatibility_level": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "subscription_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "correlation_filter": {
        "block": {
          "attributes": {
            "content_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "correlation_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "label": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "message_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "properties": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "reply_to": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "reply_to_session_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "session_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "to": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
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

func AzurermServicebusSubscriptionRuleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermServicebusSubscriptionRule), &result)
	return &result
}
