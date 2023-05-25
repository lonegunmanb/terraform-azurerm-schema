package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermStorageAccount = `{
  "block": {
    "attributes": {
      "access_tier": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "account_kind": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "account_replication_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "account_tier": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "allow_blob_public_access": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "custom_domain": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "name": "string"
            }
          ]
        ]
      },
      "enable_https_traffic_only": {
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
      "infrastructure_encryption_enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "is_hns_enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "location": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "min_tls_version": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "primary_access_key": {
        "computed": true,
        "description_kind": "plain",
        "sensitive": true,
        "type": "string"
      },
      "primary_blob_connection_string": {
        "computed": true,
        "description_kind": "plain",
        "sensitive": true,
        "type": "string"
      },
      "primary_blob_endpoint": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "primary_blob_host": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "primary_connection_string": {
        "computed": true,
        "description_kind": "plain",
        "sensitive": true,
        "type": "string"
      },
      "primary_dfs_endpoint": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "primary_dfs_host": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "primary_file_endpoint": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "primary_file_host": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "primary_location": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "primary_queue_endpoint": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "primary_queue_host": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "primary_table_endpoint": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "primary_table_host": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "primary_web_endpoint": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "primary_web_host": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "queue_encryption_key_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "resource_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "secondary_access_key": {
        "computed": true,
        "description_kind": "plain",
        "sensitive": true,
        "type": "string"
      },
      "secondary_blob_connection_string": {
        "computed": true,
        "description_kind": "plain",
        "sensitive": true,
        "type": "string"
      },
      "secondary_blob_endpoint": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "secondary_blob_host": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "secondary_connection_string": {
        "computed": true,
        "description_kind": "plain",
        "sensitive": true,
        "type": "string"
      },
      "secondary_dfs_endpoint": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "secondary_dfs_host": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "secondary_file_endpoint": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "secondary_file_host": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "secondary_location": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "secondary_queue_endpoint": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "secondary_queue_host": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "secondary_table_endpoint": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "secondary_table_host": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "secondary_web_endpoint": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "secondary_web_host": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "table_encryption_key_type": {
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
  "version": 2
}`

func AzurermStorageAccountSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermStorageAccount), &result)
	return &result
}
