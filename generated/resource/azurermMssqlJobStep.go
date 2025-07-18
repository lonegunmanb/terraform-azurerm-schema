package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermMssqlJobStep = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "initial_retry_interval_seconds": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "job_credential_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "job_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "job_step_index": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "job_target_group_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "maximum_retry_interval_seconds": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "retry_attempts": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "retry_interval_backoff_multiplier": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "sql_script": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "timeout_seconds": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      }
    },
    "block_types": {
      "output_target": {
        "block": {
          "attributes": {
            "job_credential_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "mssql_database_id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "schema_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "table_name": {
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

func AzurermMssqlJobStepSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermMssqlJobStep), &result)
	return &result
}
