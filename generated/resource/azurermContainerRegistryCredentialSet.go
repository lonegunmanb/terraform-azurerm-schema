package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermContainerRegistryCredentialSet = `{
  "block": {
    "attributes": {
      "container_registry_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "login_server": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description": "The name of the credential set.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "authentication_credentials": {
        "block": {
          "attributes": {
            "password_secret_id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "username_secret_id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "identity": {
        "block": {
          "attributes": {
            "principal_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "tenant_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "type": {
              "description_kind": "plain",
              "required": true,
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

func AzurermContainerRegistryCredentialSetSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermContainerRegistryCredentialSet), &result)
	return &result
}
