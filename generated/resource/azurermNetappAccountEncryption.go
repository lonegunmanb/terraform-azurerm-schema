package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermNetappAccountEncryption = `{
  "block": {
    "attributes": {
      "cross_tenant_key_vault_resource_id": {
        "description": "The full resource ID of the cross-tenant key vault. Required when using federated_client_id for cross-tenant scenarios.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "encryption_key": {
        "description": "The versionless encryption key url.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "federated_client_id": {
        "description": "The Client ID of the multi-tenant Entra ID application used to access cross-tenant key vaults.",
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
      "netapp_account_id": {
        "description": "The ID of the NetApp Account where encryption will be set.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "system_assigned_identity_principal_id": {
        "description": "The Principal ID of the System Assigned Identity to use for encryption.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "user_assigned_identity_id": {
        "description": "The resource ID of the User Assigned Identity to use for encryption.",
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

func AzurermNetappAccountEncryptionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermNetappAccountEncryption), &result)
	return &result
}
