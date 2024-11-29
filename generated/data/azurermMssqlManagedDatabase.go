package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermMssqlManagedDatabase = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "long_term_retention_policy": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "immutable_backups_enabled": "bool",
              "monthly_retention": "string",
              "week_of_year": "number",
              "weekly_retention": "string",
              "yearly_retention": "string"
            }
          ]
        ]
      },
      "managed_instance_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "managed_instance_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "point_in_time_restore": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "restore_point_in_time": "string",
              "source_database_id": "string"
            }
          ]
        ]
      },
      "resource_group_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "short_term_retention_days": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
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

func AzurermMssqlManagedDatabaseSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermMssqlManagedDatabase), &result)
	return &result
}
