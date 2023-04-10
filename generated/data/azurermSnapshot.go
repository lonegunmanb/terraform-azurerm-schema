package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermSnapshot = `{
  "block": {
    "attributes": {
      "creation_option": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
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
      "name": {
        "description_kind": "plain",
        "required": true,
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
      "time_created": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "trusted_launch_enabled": {
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
  "version": 1
}`

func AzurermSnapshotSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermSnapshot), &result)
	return &result
}
