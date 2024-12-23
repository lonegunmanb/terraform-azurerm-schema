package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermHealthcareService = `{
  "block": {
    "attributes": {
      "access_policy_object_ids": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "authentication_configuration": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "audience": "string",
              "authority": "string",
              "smart_proxy_enabled": "bool"
            }
          ]
        ]
      },
      "cors_configuration": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "allow_credentials": "bool",
              "allowed_headers": [
                "set",
                "string"
              ],
              "allowed_methods": [
                "list",
                "string"
              ],
              "allowed_origins": [
                "set",
                "string"
              ],
              "max_age_in_seconds": "number"
            }
          ]
        ]
      },
      "cosmosdb_key_vault_key_versionless_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "cosmosdb_throughput": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "kind": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "location": {
        "description_kind": "plain",
        "required": true,
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

func AzurermHealthcareServiceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermHealthcareService), &result)
	return &result
}
