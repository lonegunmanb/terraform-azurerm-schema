package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermContainerRegistry = `{
  "block": {
    "attributes": {
      "admin_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "admin_password": {
        "computed": true,
        "description_kind": "plain",
        "sensitive": true,
        "type": "string"
      },
      "admin_username": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "anonymous_pull_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "data_endpoint_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "encryption": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          [
            "object",
            {
              "enabled": "bool",
              "identity_client_id": "string",
              "key_vault_key_id": "string"
            }
          ]
        ]
      },
      "export_policy_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "georeplication_locations": {
        "computed": true,
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "georeplications": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          [
            "object",
            {
              "location": "string",
              "regional_endpoint_enabled": "bool",
              "tags": [
                "map",
                "string"
              ],
              "zone_redundancy_enabled": "bool"
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
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "login_server": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "network_rule_bypass_option": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "network_rule_set": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          [
            "object",
            {
              "default_action": "string",
              "ip_rule": [
                "set",
                [
                  "object",
                  {
                    "action": "string",
                    "ip_range": "string"
                  }
                ]
              ],
              "virtual_network": [
                "set",
                [
                  "object",
                  {
                    "action": "string",
                    "subnet_id": "string"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "public_network_access_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "quarantine_policy_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "resource_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "retention_policy": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          [
            "object",
            {
              "days": "number",
              "enabled": "bool"
            }
          ]
        ]
      },
      "sku": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "storage_account_id": {
        "computed": true,
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tags": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "trust_policy": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          [
            "object",
            {
              "enabled": "bool"
            }
          ]
        ]
      },
      "zone_redundancy_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      }
    },
    "block_types": {
      "identity": {
        "block": {
          "attributes": {
            "identity_ids": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
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
  "version": 2
}`

func AzurermContainerRegistrySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermContainerRegistry), &result)
	return &result
}
