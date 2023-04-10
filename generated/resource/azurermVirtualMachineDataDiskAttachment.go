package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermVirtualMachineDataDiskAttachment = `{
  "block": {
    "attributes": {
      "caching": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "create_option": {
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
      "lun": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "managed_disk_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "virtual_machine_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "write_accelerator_enabled": {
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

func AzurermVirtualMachineDataDiskAttachmentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermVirtualMachineDataDiskAttachment), &result)
	return &result
}
