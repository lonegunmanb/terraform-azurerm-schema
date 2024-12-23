package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermLogicAppTriggerRecurrence = `{
  "block": {
    "attributes": {
      "frequency": {
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
      "interval": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "logic_app_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "start_time": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "time_zone": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "schedule": {
        "block": {
          "attributes": {
            "at_these_hours": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "number"
              ]
            },
            "at_these_minutes": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "number"
              ]
            },
            "on_these_days": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
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

func AzurermLogicAppTriggerRecurrenceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermLogicAppTriggerRecurrence), &result)
	return &result
}
