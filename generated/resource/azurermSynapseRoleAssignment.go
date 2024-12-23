package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermSynapseRoleAssignment = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "principal_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "principal_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "role_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "synapse_spark_pool_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "synapse_workspace_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
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

func AzurermSynapseRoleAssignmentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermSynapseRoleAssignment), &result)
	return &result
}
