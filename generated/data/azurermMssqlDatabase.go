package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermMssqlDatabase = `{
  "block": {
    "attributes": {
      "collation": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "elastic_pool_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "enclave_type": {
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
      "identity": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "identity_ids": [
                "list",
                "string"
              ],
              "type": "string"
            }
          ]
        ]
      },
      "license_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "max_size_gb": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "read_replica_count": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "read_scale": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "server_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "sku_name": {
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
      "transparent_data_encryption_enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "transparent_data_encryption_key_automatic_rotation_enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "transparent_data_encryption_key_vault_key_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "zone_redundant": {
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

func AzurermMssqlDatabaseSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermMssqlDatabase), &result)
	return &result
}
