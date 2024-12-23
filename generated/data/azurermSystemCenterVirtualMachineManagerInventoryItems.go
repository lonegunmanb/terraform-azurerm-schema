package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermSystemCenterVirtualMachineManagerInventoryItems = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "inventory_items": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "id": "string",
              "name": "string",
              "uuid": "string"
            }
          ]
        ]
      },
      "inventory_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "system_center_virtual_machine_manager_server_id": {
        "description_kind": "plain",
        "required": true,
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

func AzurermSystemCenterVirtualMachineManagerInventoryItemsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermSystemCenterVirtualMachineManagerInventoryItems), &result)
	return &result
}
