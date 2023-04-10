package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermAppConfiguration = `{
  "block": {
    "attributes": {
      "encryption": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "identity_client_id": "string",
              "key_vault_key_identifier": "string"
            }
          ]
        ]
      },
      "endpoint": {
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
      "local_auth_enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
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
      "primary_read_key": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "connection_string": "string",
              "id": "string",
              "secret": "string"
            }
          ]
        ]
      },
      "primary_write_key": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "connection_string": "string",
              "id": "string",
              "secret": "string"
            }
          ]
        ]
      },
      "public_network_access": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "public_network_access_enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "purge_protection_enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "resource_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "secondary_read_key": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "connection_string": "string",
              "id": "string",
              "secret": "string"
            }
          ]
        ]
      },
      "secondary_write_key": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "connection_string": "string",
              "id": "string",
              "secret": "string"
            }
          ]
        ]
      },
      "sku": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "soft_delete_retention_days": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
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

func AzurermAppConfigurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermAppConfiguration), &result)
	return &result
}
