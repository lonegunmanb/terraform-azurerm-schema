package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermStorageDataLakeGen2Path = `{
  "block": {
    "attributes": {
      "filesystem_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "group": {
        "computed": true,
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
      "owner": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "path": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "resource": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "storage_account_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "ace": {
        "block": {
          "attributes": {
            "id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "permissions": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "scope": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
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

func AzurermStorageDataLakeGen2PathSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermStorageDataLakeGen2Path), &result)
	return &result
}
