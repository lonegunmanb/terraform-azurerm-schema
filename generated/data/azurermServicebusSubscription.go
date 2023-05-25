package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermServicebusSubscription = `{
  "block": {
    "attributes": {
      "auto_delete_on_idle": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "dead_lettering_on_filter_evaluation_error": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "dead_lettering_on_message_expiration": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "default_message_ttl": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "enable_batched_operations": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "forward_dead_lettered_messages_to": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "forward_to": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "lock_duration": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "max_delivery_count": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "namespace_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "requires_session": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "resource_group_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "topic_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "topic_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
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

func AzurermServicebusSubscriptionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermServicebusSubscription), &result)
	return &result
}
