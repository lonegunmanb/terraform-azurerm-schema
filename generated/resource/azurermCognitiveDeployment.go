package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermCognitiveDeployment = `{
  "block": {
    "attributes": {
      "cognitive_account_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "dynamic_throttling_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
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
      "rai_policy_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "version_upgrade_option": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "model": {
        "block": {
          "attributes": {
            "format": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "version": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "sku": {
        "block": {
          "attributes": {
            "capacity": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "family": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "size": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "tier": {
              "description_kind": "plain",
              "optional": true,
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
  "version": 0
}`

func AzurermCognitiveDeploymentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermCognitiveDeployment), &result)
	return &result
}
