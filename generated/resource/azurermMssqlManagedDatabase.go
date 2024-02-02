package resource

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
      "managed_instance_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "short_term_retention_days": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      }
    },
    "block_types": {
      "long_term_retention_policy": {
        "block": {
          "attributes": {
            "immutable_backups_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "monthly_retention": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "week_of_year": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "weekly_retention": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "yearly_retention": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "point_in_time_restore": {
        "block": {
          "attributes": {
            "restore_point_in_time": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "source_database_id": {
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

func AzurermMssqlManagedDatabaseSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermMssqlManagedDatabase), &result)
	return &result
}
