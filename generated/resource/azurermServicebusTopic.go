package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermServicebusTopic = `{
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
      "default_message_ttl": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "duplicate_detection_history_time_window": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "express_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "max_message_size_in_kilobytes": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "max_size_in_megabytes": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "namespace_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "partitioning_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "requires_duplicate_detection": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "status": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "support_ordering": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
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

func AzurermServicebusTopicSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermServicebusTopic), &result)
	return &result
}
