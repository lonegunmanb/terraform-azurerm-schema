package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermKeyVaultEncryptedValue = `{
  "block": {
    "attributes": {
      "algorithm": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "decoded_plain_text_value": {
        "computed": true,
        "description_kind": "plain",
        "sensitive": true,
        "type": "string"
      },
      "encrypted_data": {
        "description_kind": "plain",
        "optional": true,
        "sensitive": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "key_vault_key_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "plain_text_value": {
        "description_kind": "plain",
        "optional": true,
        "sensitive": true,
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

func AzurermKeyVaultEncryptedValueSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermKeyVaultEncryptedValue), &result)
	return &result
}
