package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermMssqlServerExtendedAuditingPolicy = `{
  "block": {
    "attributes": {
      "audit_actions_and_groups": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "log_monitoring_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "predicate_expression": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "retention_in_days": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "server_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "storage_account_access_key": {
        "description_kind": "plain",
        "optional": true,
        "sensitive": true,
        "type": "string"
      },
      "storage_account_access_key_is_secondary": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "storage_account_subscription_id": {
        "description_kind": "plain",
        "optional": true,
        "sensitive": true,
        "type": "string"
      },
      "storage_endpoint": {
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

func AzurermMssqlServerExtendedAuditingPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermMssqlServerExtendedAuditingPolicy), &result)
	return &result
}
