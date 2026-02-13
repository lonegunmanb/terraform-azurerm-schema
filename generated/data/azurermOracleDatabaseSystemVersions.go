package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermOracleDatabaseSystemVersions = `{
  "block": {
    "attributes": {
      "database_software_image_supported": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "database_system_shape": {
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
      "location": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "shape_family": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "storage_management": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "upgrade_supported": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "versions": {
        "computed": true,
        "description": "A list of available Oracle Database versions and their properties.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "latest_version": "bool",
              "name": "string",
              "pluggable_database_supported": "bool",
              "version": "string"
            }
          ]
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

func AzurermOracleDatabaseSystemVersionsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermOracleDatabaseSystemVersions), &result)
	return &result
}
