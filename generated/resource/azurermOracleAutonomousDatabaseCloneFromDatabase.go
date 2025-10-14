package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermOracleAutonomousDatabaseCloneFromDatabase = `{
  "block": {
    "attributes": {
      "admin_password": {
        "description_kind": "plain",
        "required": true,
        "sensitive": true,
        "type": "string"
      },
      "allowed_ip_addresses": {
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
      "clone_type": {
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
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "data_storage_size_in_tb": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "database_version": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "database_workload": {
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
      "refreshable_model": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "resource_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "source_autonomous_database_id": {
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

func AzurermOracleAutonomousDatabaseCloneFromDatabaseSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermOracleAutonomousDatabaseCloneFromDatabase), &result)
	return &result
}
