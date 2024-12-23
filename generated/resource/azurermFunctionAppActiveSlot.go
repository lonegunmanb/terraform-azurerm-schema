package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermFunctionAppActiveSlot = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "last_successful_swap": {
        "computed": true,
        "description": "The timestamp of the last successful swap with ` + "`" + `Production` + "`" + `",
        "description_kind": "plain",
        "type": "string"
      },
      "overwrite_network_config": {
        "description": "The swap action should overwrite the Production slot's network configuration with the configuration from this slot. Defaults to ` + "`" + `true` + "`" + `.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "slot_id": {
        "description": "The ID of the Slot to swap with ` + "`" + `Production` + "`" + `.",
        "description_kind": "plain",
        "required": true,
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

func AzurermFunctionAppActiveSlotSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermFunctionAppActiveSlot), &result)
	return &result
}
