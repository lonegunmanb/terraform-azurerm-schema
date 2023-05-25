package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermServicebusQueue = `{
  "block": {
    "attributes": {
      "auto_delete_on_idle": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
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
      "duplicate_detection_history_time_window": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "enable_batched_operations": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "enable_express": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "enable_partitioning": {
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
      "max_size_in_megabytes": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "namespace_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "namespace_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "requires_duplicate_detection": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
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
      "status": {
        "computed": true,
        "description_kind": "plain",
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

func AzurermServicebusQueueSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermServicebusQueue), &result)
	return &result
}
