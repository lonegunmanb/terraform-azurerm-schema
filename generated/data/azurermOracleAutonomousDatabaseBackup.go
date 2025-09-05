package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermOracleAutonomousDatabaseBackup = `{
  "block": {
    "attributes": {
      "automatic": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "autonomous_database_backup_ocid": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "autonomous_database_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "autonomous_database_ocid": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "database_backup_size_in_tbs": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "database_version": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "display_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "lifecycle_details": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "lifecycle_state": {
        "computed": true,
        "description_kind": "plain",
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
      "provisioning_state": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "restorable": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "retention_period_in_days": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "time_available_til": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "time_ended": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "time_started": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "type": {
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
    "description_kind": "plain"
  },
  "version": 0
}`

func AzurermOracleAutonomousDatabaseBackupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermOracleAutonomousDatabaseBackup), &result)
	return &result
}
