package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermMariadbServer = `{
  "block": {
    "attributes": {
      "administrator_login": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "fqdn": {
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
      "resource_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "sku_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "ssl_enforcement": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "storage_profile": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "auto_grow": "string",
              "backup_retention_days": "number",
              "geo_redundant_backup": "string",
              "storage_mb": "number"
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
      "version": {
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
    "deprecated": true,
    "description_kind": "plain"
  },
  "version": 0
}`

func AzurermMariadbServerSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermMariadbServer), &result)
	return &result
}
