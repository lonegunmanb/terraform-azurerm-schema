package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermDevCenterCatalog = `{
  "block": {
    "attributes": {
      "dev_center_id": {
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
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "resource_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "catalog_adogit": {
        "block": {
          "attributes": {
            "branch": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "key_vault_key_url": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "path": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "uri": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "catalog_github": {
        "block": {
          "attributes": {
            "branch": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "key_vault_key_url": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "path": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "uri": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
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

func AzurermDevCenterCatalogSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermDevCenterCatalog), &result)
	return &result
}
