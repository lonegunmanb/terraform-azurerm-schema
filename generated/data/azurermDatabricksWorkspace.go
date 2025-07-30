package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermDatabricksWorkspace = `{
  "block": {
    "attributes": {
      "custom_parameters": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "machine_learning_workspace_id": "string",
              "nat_gateway_name": "string",
              "no_public_ip": "bool",
              "private_subnet_name": "string",
              "public_ip_name": "string",
              "public_subnet_name": "string",
              "storage_account_name": "string",
              "storage_account_sku_name": "string",
              "virtual_network_id": "string",
              "vnet_address_prefix": "string"
            }
          ]
        ]
      },
      "enhanced_security_compliance": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "automatic_cluster_update_enabled": "bool",
              "compliance_security_profile_enabled": "bool",
              "compliance_security_profile_standards": [
                "set",
                "string"
              ],
              "enhanced_security_monitoring_enabled": "bool"
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
      "location": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
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
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "resource_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "sku": {
        "computed": true,
        "description_kind": "plain",
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

func AzurermDatabricksWorkspaceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermDatabricksWorkspace), &result)
	return &result
}
