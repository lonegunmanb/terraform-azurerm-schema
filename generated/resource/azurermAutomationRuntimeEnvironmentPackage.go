package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermAutomationRuntimeEnvironmentPackage = `{
  "block": {
    "attributes": {
      "automation_runtime_environment_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "content_uri": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "content_version": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "default": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "hash_algorithm": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "hash_value": {
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
      "size_in_bytes": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "version": {
        "computed": true,
        "description_kind": "plain",
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

func AzurermAutomationRuntimeEnvironmentPackageSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermAutomationRuntimeEnvironmentPackage), &result)
	return &result
}
