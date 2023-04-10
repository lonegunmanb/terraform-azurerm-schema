package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermAppConfigurationKeys = `{
  "block": {
    "attributes": {
      "configuration_store_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "items": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "content_type": "string",
              "etag": "string",
              "key": "string",
              "label": "string",
              "locked": "bool",
              "tags": [
                "map",
                "string"
              ],
              "type": "string",
              "value": "string",
              "vault_key_reference": "string"
            }
          ]
        ]
      },
      "key": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "label": {
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

func AzurermAppConfigurationKeysSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermAppConfigurationKeys), &result)
	return &result
}
