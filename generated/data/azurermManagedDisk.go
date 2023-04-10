package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermManagedDisk = `{
  "block": {
    "attributes": {
      "create_option": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "disk_access_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "disk_encryption_set_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "disk_iops_read_write": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "disk_mbps_read_write": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "disk_size_gb": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "encryption_settings": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "disk_encryption_key": [
                "list",
                [
                  "object",
                  {
                    "secret_url": "string",
                    "source_vault_id": "string"
                  }
                ]
              ],
              "enabled": "bool",
              "key_encryption_key": [
                "list",
                [
                  "object",
                  {
                    "key_url": "string",
                    "source_vault_id": "string"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "image_reference_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "network_access_policy": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "os_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "resource_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "source_resource_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "source_uri": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "storage_account_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "storage_account_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "zones": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
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
  "version": 1
}`

func AzurermManagedDiskSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermManagedDisk), &result)
	return &result
}
