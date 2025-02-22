package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermHealthcareDicomService = `{
  "block": {
    "attributes": {
      "authentication": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "audience": [
                "list",
                "string"
              ],
              "authority": "string"
            }
          ]
        ]
      },
      "cors": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "allow_credentials": "bool",
              "allowed_headers": [
                "list",
                "string"
              ],
              "allowed_methods": [
                "list",
                "string"
              ],
              "allowed_origins": [
                "list",
                "string"
              ],
              "max_age_in_seconds": "number"
            }
          ]
        ]
      },
      "data_partitions_enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "encryption_key_url": {
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
      "private_endpoint": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "id": "string",
              "name": "string"
            }
          ]
        ]
      },
      "service_url": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "storage": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "file_system_name": "string",
              "storage_account_id": "string"
            }
          ]
        ]
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "workspace_id": {
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

func AzurermHealthcareDicomServiceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermHealthcareDicomService), &result)
	return &result
}
