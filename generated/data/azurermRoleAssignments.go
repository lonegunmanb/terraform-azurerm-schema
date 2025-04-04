package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermRoleAssignments = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "limit_at_scope": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "principal_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "role_assignments": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "condition": "string",
              "condition_version": "string",
              "delegated_managed_identity_resource_id": "string",
              "description": "string",
              "principal_id": "string",
              "principal_type": "string",
              "role_assignment_id": "string",
              "role_assignment_name": "string",
              "role_assignment_scope": "string",
              "role_definition_id": "string"
            }
          ]
        ]
      },
      "scope": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tenant_id": {
        "description_kind": "plain",
        "optional": true,
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

func AzurermRoleAssignmentsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermRoleAssignments), &result)
	return &result
}
