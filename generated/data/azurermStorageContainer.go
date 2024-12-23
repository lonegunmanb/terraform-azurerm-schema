package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermStorageContainer = `{
  "block": {
    "attributes": {
      "container_access_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "default_encryption_scope": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "encryption_scope_override_enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "has_immutability_policy": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "has_legal_hold": {
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
      "metadata": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "resource_manager_id": {
        "computed": true,
        "deprecated": true,
        "description_kind": "plain",
        "type": "string"
      },
      "storage_account_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "storage_account_name": {
        "description_kind": "plain",
        "optional": true,
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

func AzurermStorageContainerSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermStorageContainer), &result)
	return &result
}
