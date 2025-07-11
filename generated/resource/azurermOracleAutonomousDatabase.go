package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermOracleAutonomousDatabase = `{
  "block": {
    "attributes": {
      "admin_password": {
        "description_kind": "plain",
        "required": true,
        "sensitive": true,
        "type": "string"
      },
      "allowed_ips": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "auto_scaling_enabled": {
        "description_kind": "plain",
        "required": true,
        "type": "bool"
      },
      "auto_scaling_for_storage_enabled": {
        "description_kind": "plain",
        "required": true,
        "type": "bool"
      },
      "backup_retention_period_in_days": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "character_set": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "compute_count": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "compute_model": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "customer_contacts": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "data_storage_size_in_tbs": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "db_version": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "db_workload": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "display_name": {
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
      "license_model": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "location": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "mtls_connection_required": {
        "description_kind": "plain",
        "required": true,
        "type": "bool"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "national_character_set": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "resource_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "subnet_id": {
        "description_kind": "plain",
        "optional": true,
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
      "virtual_network_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "long_term_backup_schedule": {
        "block": {
          "attributes": {
            "enabled": {
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            },
            "repeat_cadence": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "retention_period_in_days": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "time_of_backup": {
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

func AzurermOracleAutonomousDatabaseSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermOracleAutonomousDatabase), &result)
	return &result
}
