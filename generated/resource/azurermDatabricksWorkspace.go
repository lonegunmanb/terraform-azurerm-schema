package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermDatabricksWorkspace = `{
  "block": {
    "attributes": {
      "customer_managed_key_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "disk_encryption_set_id": {
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
      "infrastructure_encryption_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "load_balancer_backend_address_pool_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "location": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "managed_disk_cmk_key_vault_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "managed_disk_cmk_key_vault_key_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "managed_disk_cmk_rotation_to_latest_version_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "managed_disk_identity": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "principal_id": "string",
              "tenant_id": "string",
              "type": "string"
            }
          ]
        ]
      },
      "managed_resource_group_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "managed_resource_group_name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "managed_services_cmk_key_vault_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "managed_services_cmk_key_vault_key_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "network_security_group_rules_required": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "public_network_access_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "resource_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "sku": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "storage_account_identity": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "principal_id": "string",
              "tenant_id": "string",
              "type": "string"
            }
          ]
        ]
      },
      "tags": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "workspace_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "workspace_url": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "custom_parameters": {
        "block": {
          "attributes": {
            "machine_learning_workspace_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "nat_gateway_name": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "no_public_ip": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "private_subnet_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "private_subnet_network_security_group_association_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "public_ip_name": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "public_subnet_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "public_subnet_network_security_group_association_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "storage_account_name": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "storage_account_sku_name": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "virtual_network_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "vnet_address_prefix": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
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

func AzurermDatabricksWorkspaceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermDatabricksWorkspace), &result)
	return &result
}
