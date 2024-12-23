package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermStorageMoverJobDefinition = `{
  "block": {
    "attributes": {
      "agent_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "copy_mode": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "description": {
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
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "source_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "source_sub_path": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "storage_mover_project_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "target_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "target_sub_path": {
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

func AzurermStorageMoverJobDefinitionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermStorageMoverJobDefinition), &result)
	return &result
}
