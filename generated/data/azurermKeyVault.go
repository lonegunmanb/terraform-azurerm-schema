package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermKeyVault = `{
  "block": {
    "attributes": {
      "access_policy": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "application_id": "string",
              "certificate_permissions": [
                "list",
                "string"
              ],
              "key_permissions": [
                "list",
                "string"
              ],
              "object_id": "string",
              "secret_permissions": [
                "list",
                "string"
              ],
              "storage_permissions": [
                "list",
                "string"
              ],
              "tenant_id": "string"
            }
          ]
        ]
      },
      "enable_rbac_authorization": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "enabled_for_deployment": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "enabled_for_disk_encryption": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "enabled_for_template_deployment": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "location": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "network_acls": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "bypass": "string",
              "default_action": "string",
              "ip_rules": [
                "list",
                "string"
              ],
              "virtual_network_subnet_ids": [
                "list",
                "string"
              ]
            }
          ]
        ]
      },
      "public_network_access_enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "purge_protection_enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "resource_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "sku_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "tenant_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "vault_uri": {
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

func AzurermKeyVaultSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermKeyVault), &result)
	return &result
}
