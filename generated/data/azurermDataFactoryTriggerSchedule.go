package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermDataFactoryTriggerSchedule = `{
  "block": {
    "attributes": {
      "activated": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "annotations": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "data_factory_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "description": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "end_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "frequency": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "interval": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "pipeline_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "schedule": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "days_of_month": [
                "list",
                "number"
              ],
              "days_of_week": [
                "list",
                "string"
              ],
              "hours": [
                "list",
                "number"
              ],
              "minutes": [
                "list",
                "number"
              ],
              "monthly": [
                "list",
                [
                  "object",
                  {
                    "week": "number",
                    "weekday": "string"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "start_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "time_zone": {
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

func AzurermDataFactoryTriggerScheduleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermDataFactoryTriggerSchedule), &result)
	return &result
}
