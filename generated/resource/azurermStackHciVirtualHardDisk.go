package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermStackHciVirtualHardDisk = `{
  "block": {
    "attributes": {
      "block_size_in_bytes": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "custom_location_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "disk_file_format": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "disk_size_in_gb": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "dynamic_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "hyperv_generation": {
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
      "location": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "logical_sector_in_bytes": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "physical_sector_in_bytes": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "resource_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "storage_path_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tags": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
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

func AzurermStackHciVirtualHardDiskSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermStackHciVirtualHardDisk), &result)
	return &result
}
