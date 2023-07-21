package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermServicebusTopic = `{
  "block": {
    "attributes": {
      "auto_delete_on_idle": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
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
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
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
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "requires_duplicate_detection": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "resource_group_name": {
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "support_ordering": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
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

func AzurermServicebusTopicSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermServicebusTopic), &result)
	return &result
}
