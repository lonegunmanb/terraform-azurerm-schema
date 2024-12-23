package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermServicebusSubscription = `{
  "block": {
    "attributes": {
      "auto_delete_on_idle": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "batched_operations_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "client_scoped_subscription_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "dead_lettering_on_filter_evaluation_error": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "dead_lettering_on_message_expiration": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "default_message_ttl": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "forward_dead_lettered_messages_to": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "forward_to": {
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
      "lock_duration": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "max_delivery_count": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "requires_session": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "status": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "topic_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "client_scoped_subscription": {
        "block": {
          "attributes": {
            "client_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "is_client_scoped_subscription_durable": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "is_client_scoped_subscription_shareable": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
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
  "version": 1
}`

func AzurermServicebusSubscriptionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermServicebusSubscription), &result)
	return &result
}
