package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermPimActiveRoleAssignment = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "justification": {
        "description": "The justification of the role assignment.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "principal_id": {
        "description": "The principal id.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "principal_type": {
        "computed": true,
        "description": "The type of principal.",
        "description_kind": "plain",
        "type": "string"
      },
      "role_definition_id": {
        "description": "The role definition id.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "scope": {
        "description": "The scope.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "schedule": {
        "block": {
          "attributes": {
            "start_date_time": {
              "computed": true,
              "description": "The start date time.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "expiration": {
              "block": {
                "attributes": {
                  "duration_days": {
                    "computed": true,
                    "description": "The duration of the assignment in days.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "duration_hours": {
                    "computed": true,
                    "description": "The duration of the assignment in hours.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "end_date_time": {
                    "computed": true,
                    "description": "The end date time of the assignment.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "The schedule details of this role assignment.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "ticket": {
        "block": {
          "attributes": {
            "number": {
              "description": "The ticket number.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "system": {
              "description": "The ticket system.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "The ticket details.",
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

func AzurermPimActiveRoleAssignmentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermPimActiveRoleAssignment), &result)
	return &result
}
