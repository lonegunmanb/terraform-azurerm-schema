package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermKeyVaultKey = `{
  "block": {
    "attributes": {
      "curve": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "e": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "expiration_date": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "key_opts": {
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
      },
      "key_size": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "key_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "key_vault_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "n": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "not_before_date": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "public_key_openssh": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "public_key_pem": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "version": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "versionless_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "x": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "y": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "timeouts": {
        "block": {
          "attributes": {
            "create": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "delete": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "read": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "update": {
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

func AzurermKeyVaultKeySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermKeyVaultKey), &result)
	return &result
}
