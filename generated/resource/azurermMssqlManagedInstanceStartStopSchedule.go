package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermMssqlManagedInstanceStartStopSchedule = `{
  "block": {
    "attributes": {
      "description": {
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
      "managed_instance_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "next_execution_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "next_run_action": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "timezone_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "schedule": {
        "block": {
          "attributes": {
            "start_day": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "start_time": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "stop_day": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "stop_time": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "min_items": 1,
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

func AzurermMssqlManagedInstanceStartStopScheduleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermMssqlManagedInstanceStartStopSchedule), &result)
	return &result
}
