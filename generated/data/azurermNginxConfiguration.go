package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermNginxConfiguration = `{
  "block": {
    "attributes": {
      "config_file": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          [
            "object",
            {
              "content": "string",
              "virtual_path": "string"
            }
          ]
        ]
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "nginx_deployment_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "package_data": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "protected_file": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          [
            "object",
            {
              "content": "string",
              "content_hash": "string",
              "virtual_path": "string"
            }
          ]
        ]
      },
      "root_file": {
        "computed": true,
        "description_kind": "plain",
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

func AzurermNginxConfigurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermNginxConfiguration), &result)
	return &result
}
