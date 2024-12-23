package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermCostManagementScheduledAction = `{
  "block": {
    "attributes": {
      "day_of_month": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "days_of_week": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "display_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "email_address_sender": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "email_addresses": {
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
      },
      "email_subject": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "end_date": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "frequency": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "hour_of_day": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "message": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "start_date": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "view_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "weeks_of_month": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
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

func AzurermCostManagementScheduledActionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermCostManagementScheduledAction), &result)
	return &result
}
