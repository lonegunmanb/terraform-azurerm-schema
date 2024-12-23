package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermPimEligibleRoleAssignment = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "justification": {
        "computed": true,
        "description": "The justification for this eligible role assignment",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "principal_id": {
        "description": "Object ID of the principal for this eligible role assignment",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "principal_type": {
        "computed": true,
        "description": "Type of principal to which the role will be assigned",
        "description_kind": "plain",
        "type": "string"
      },
      "role_definition_id": {
        "description": "Role definition ID for this eligible role assignment",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "scope": {
        "description": "Scope for this eligible role assignment, should be a valid resource ID",
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
              "description": "The start date/time",
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
                    "description": "The duration of the eligible role assignment in days",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "duration_hours": {
                    "computed": true,
                    "description": "The duration of the eligible role assignment in hours",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "end_date_time": {
                    "computed": true,
                    "description": "The end date/time of the eligible role assignment",
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
          "description": "The schedule details for this eligible role assignment",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "ticket": {
        "block": {
          "attributes": {
            "number": {
              "description": "User-supplied ticket number to be included with the request",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "system": {
              "description": "User-supplied ticket system name to be included with the request",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Ticket details relating to the eligible assignment",
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

func AzurermPimEligibleRoleAssignmentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermPimEligibleRoleAssignment), &result)
	return &result
}
