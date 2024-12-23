package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermSynapseWorkspaceSecurityAlertPolicy = `{
  "block": {
    "attributes": {
      "disabled_alerts": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "email_account_admins_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "email_addresses": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "policy_state": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "retention_days": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "storage_account_access_key": {
        "description_kind": "plain",
        "optional": true,
        "sensitive": true,
        "type": "string"
      },
      "storage_endpoint": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "synapse_workspace_id": {
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

func AzurermSynapseWorkspaceSecurityAlertPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermSynapseWorkspaceSecurityAlertPolicy), &result)
	return &result
}
