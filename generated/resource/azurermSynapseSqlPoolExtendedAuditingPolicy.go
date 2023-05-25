package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermSynapseSqlPoolExtendedAuditingPolicy = `{
  "block": {
    "attributes": {
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
      "retention_in_days": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "sql_pool_id": {
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

func AzurermSynapseSqlPoolExtendedAuditingPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermSynapseSqlPoolExtendedAuditingPolicy), &result)
	return &result
}
