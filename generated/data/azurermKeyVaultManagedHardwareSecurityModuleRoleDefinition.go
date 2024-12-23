package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermKeyVaultManagedHardwareSecurityModuleRoleDefinition = `{
  "block": {
    "attributes": {
      "assignable_scopes": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "description": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "managed_hsm_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "permission": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "actions": [
                "list",
                "string"
              ],
              "data_actions": [
                "list",
                "string"
              ],
              "not_actions": [
                "list",
                "string"
              ],
              "not_data_actions": [
                "list",
                "string"
              ]
            }
          ]
        ]
      },
      "resource_manager_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "role_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "role_type": {
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

func AzurermKeyVaultManagedHardwareSecurityModuleRoleDefinitionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermKeyVaultManagedHardwareSecurityModuleRoleDefinition), &result)
	return &result
}
