package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermResourceGroupPolicyRemediation = `{
  "block": {
    "attributes": {
      "failure_percentage": {
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
      "location_filters": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "parallel_deployments": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "policy_assignment_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "policy_definition_reference_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "resource_count": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "resource_discovery_mode": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "resource_group_id": {
        "description_kind": "plain",
        "required": true,
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

func AzurermResourceGroupPolicyRemediationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermResourceGroupPolicyRemediation), &result)
	return &result
}
