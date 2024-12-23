package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermContainerAppEnvironmentCustomDomain = `{
  "block": {
    "attributes": {
      "certificate_blob_base64": {
        "description": "The Custom Domain Certificate Private Key as a base64 encoded PFX or PEM.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "certificate_password": {
        "description": "The Custom Domain Certificate password.",
        "description_kind": "plain",
        "required": true,
        "sensitive": true,
        "type": "string"
      },
      "container_app_environment_id": {
        "description": "The Container App Managed Environment ID to configure this Custom Domain on.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "dns_suffix": {
        "description": "The Custom Domain DNS suffix for this Container App Environment.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
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

func AzurermContainerAppEnvironmentCustomDomainSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermContainerAppEnvironmentCustomDomain), &result)
	return &result
}
