package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermAutomationSourceControl = `{
  "block": {
    "attributes": {
      "automatic_sync": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "automation_account_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "branch": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "folder_path": {
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
      },
      "publish_runbook_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "repository_url": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "source_control_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "security": {
        "block": {
          "attributes": {
            "refresh_token": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "token": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "token_type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
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
  "version": 1
}`

func AzurermAutomationSourceControlSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermAutomationSourceControl), &result)
	return &result
}
