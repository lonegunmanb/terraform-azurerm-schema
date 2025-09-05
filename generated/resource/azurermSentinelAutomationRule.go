package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermSentinelAutomationRule = `{
  "block": {
    "attributes": {
      "condition_json": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "display_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "expiration": {
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
      "log_analytics_workspace_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "order": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "triggers_on": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "triggers_when": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "action_incident": {
        "block": {
          "attributes": {
            "classification": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "classification_comment": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "labels": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "order": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "owner_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "severity": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "status": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "action_incident_task": {
        "block": {
          "attributes": {
            "description": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "order": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "title": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "action_playbook": {
        "block": {
          "attributes": {
            "logic_app_id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "order": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "tenant_id": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
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

func AzurermSentinelAutomationRuleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermSentinelAutomationRule), &result)
	return &result
}
