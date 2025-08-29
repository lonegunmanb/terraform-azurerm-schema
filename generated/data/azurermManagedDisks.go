package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermManagedDisks = `{
  "block": {
    "attributes": {
      "disk": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "create_option": "string",
              "disk_access_id": "string",
              "disk_encryption_set_id": "string",
              "disk_iops_read_write": "number",
              "disk_mbps_read_write": "number",
              "disk_size_in_gb": "number",
              "encryption_settings": [
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
              ],
              "id": "string",
              "image_reference_id": "string",
              "location": "string",
              "name": "string",
              "network_access_policy": "string",
              "os_type": "string",
              "source_resource_id": "string",
              "source_uri": "string",
              "storage_account_id": "string",
              "storage_account_type": "string",
              "tags": [
                "map",
                "string"
              ],
              "zones": [
                "list",
                "string"
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
      "resource_group_name": {
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

func AzurermManagedDisksSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermManagedDisks), &result)
	return &result
}
