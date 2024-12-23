package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermDevCenterProjectPool = `{
  "block": {
    "attributes": {
      "dev_box_definition_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "dev_center_attached_network_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "dev_center_project_id": {
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
      "local_administrator_enabled": {
        "description_kind": "plain",
        "required": true,
        "type": "bool"
      },
      "location": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "stop_on_disconnect_grace_period_minutes": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "tags": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
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

func AzurermDevCenterProjectPoolSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermDevCenterProjectPool), &result)
	return &result
}
