package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermSiteRecoveryHypervReplicationPolicy = `{
  "block": {
    "attributes": {
      "application_consistent_snapshot_frequency_in_hours": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "recovery_point_retention_in_hours": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "recovery_vault_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "replication_interval_in_seconds": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
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

func AzurermSiteRecoveryHypervReplicationPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermSiteRecoveryHypervReplicationPolicy), &result)
	return &result
}
