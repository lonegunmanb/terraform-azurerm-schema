package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermManagedRedis = `{
  "block": {
    "attributes": {
      "customer_managed_key": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "key_vault_key_id": "string",
              "user_assigned_identity_id": "string"
            }
          ]
        ]
      },
      "default_database": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "access_keys_authentication_enabled": "bool",
              "client_protocol": "string",
              "clustering_policy": "string",
              "eviction_policy": "string",
              "geo_replication_group_name": "string",
              "geo_replication_linked_database_ids": [
                "set",
                "string"
              ],
              "module": [
                "list",
                [
                  "object",
                  {
                    "args": "string",
                    "name": "string",
                    "version": "string"
                  }
                ]
              ],
              "port": "number",
              "primary_access_key": "string",
              "secondary_access_key": "string"
            }
          ]
        ]
      },
      "high_availability_enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "hostname": {
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
              "principal_id": "string",
              "tenant_id": "string",
              "type": "string"
            }
          ]
        ]
      },
      "location": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "resource_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "sku_name": {
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
  "version": 0
}`

func AzurermManagedRedisSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermManagedRedis), &result)
	return &result
}
