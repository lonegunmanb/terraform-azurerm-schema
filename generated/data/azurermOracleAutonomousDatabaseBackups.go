package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermOracleAutonomousDatabaseBackups = `{
  "block": {
    "attributes": {
      "autonomous_database_backups": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "automatic": "bool",
              "autonomous_database_backup_ocid": "string",
              "autonomous_database_ocid": "string",
              "database_backup_size_in_tbs": "number",
              "database_version": "string",
              "display_name": "string",
              "id": "string",
              "lifecycle_details": "string",
              "lifecycle_state": "string",
              "location": "string",
              "provisioning_state": "string",
              "restorable": "bool",
              "retention_period_in_days": "number",
              "time_available_til": "string",
              "time_ended": "string",
              "time_started": "string",
              "type": "string"
            }
          ]
        ]
      },
      "autonomous_database_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
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

func AzurermOracleAutonomousDatabaseBackupsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermOracleAutonomousDatabaseBackups), &result)
	return &result
}
