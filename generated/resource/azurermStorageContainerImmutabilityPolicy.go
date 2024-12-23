package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermStorageContainerImmutabilityPolicy = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "immutability_period_in_days": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "locked": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "protected_append_writes_all_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "protected_append_writes_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "storage_container_resource_manager_id": {
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

func AzurermStorageContainerImmutabilityPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermStorageContainerImmutabilityPolicy), &result)
	return &result
}
