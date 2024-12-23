package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermContainerRegistryCacheRule = `{
  "block": {
    "attributes": {
      "container_registry_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "credential_set_id": {
        "description": "The ARM resource ID of the credential store which is associated with the cache rule.",
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
      "name": {
        "description": "The name of the cache rule.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "source_repo": {
        "description": "The full source repository path such as 'docker.io/library/ubuntu'.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "target_repo": {
        "description": "The target repository namespace such as 'ubuntu'.",
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

func AzurermContainerRegistryCacheRuleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermContainerRegistryCacheRule), &result)
	return &result
}
