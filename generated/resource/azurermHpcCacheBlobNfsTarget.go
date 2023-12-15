package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermHpcCacheBlobNfsTarget = `{
  "block": {
    "attributes": {
      "access_policy_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "cache_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
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
      "namespace_path": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "resource_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "storage_container_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "usage_model": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "verification_timer_in_seconds": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "write_back_timer_in_seconds": {
        "description_kind": "plain",
        "optional": true,
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

func AzurermHpcCacheBlobNfsTargetSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermHpcCacheBlobNfsTarget), &result)
	return &result
}
