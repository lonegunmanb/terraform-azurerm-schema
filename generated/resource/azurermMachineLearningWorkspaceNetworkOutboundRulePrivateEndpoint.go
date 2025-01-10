package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermMachineLearningWorkspaceNetworkOutboundRulePrivateEndpoint = `{
  "block": {
    "attributes": {
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
      "service_resource_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "spark_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "sub_resource_target": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "workspace_id": {
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

func AzurermMachineLearningWorkspaceNetworkOutboundRulePrivateEndpointSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermMachineLearningWorkspaceNetworkOutboundRulePrivateEndpoint), &result)
	return &result
}
