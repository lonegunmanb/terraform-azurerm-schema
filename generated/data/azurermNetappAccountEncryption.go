package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermNetappAccountEncryption = `{
  "block": {
    "attributes": {
      "encryption_key": {
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
      "netapp_account_id": {
        "description": "The ID of the NetApp Account where encryption will be set.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "system_assigned_identity_principal_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "user_assigned_identity_id": {
        "computed": true,
        "description_kind": "plain",
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

func AzurermNetappAccountEncryptionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermNetappAccountEncryption), &result)
	return &result
}
