package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermAutomationPowershell72Module = `{
  "block": {
    "attributes": {
      "automation_account_id": {
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
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "module_link": {
        "block": {
          "attributes": {
            "uri": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "hash": {
              "block": {
                "attributes": {
                  "algorithm": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
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

func AzurermAutomationPowershell72ModuleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermAutomationPowershell72Module), &result)
	return &result
}
