package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermApiManagementIdentityProviderAadb2C = `{
  "block": {
    "attributes": {
      "allowed_tenant": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "api_management_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "authority": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "client_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "client_library": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "client_secret": {
        "description_kind": "plain",
        "required": true,
        "sensitive": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "password_reset_policy": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "profile_editing_policy": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "resource_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "signin_policy": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "signin_tenant": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "signup_policy": {
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

func AzurermApiManagementIdentityProviderAadb2CSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermApiManagementIdentityProviderAadb2C), &result)
	return &result
}
