package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermHpcCacheAccessPolicy = `{
  "block": {
    "attributes": {
      "hpc_cache_id": {
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
      }
    },
    "block_types": {
      "access_rule": {
        "block": {
          "attributes": {
            "access": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "anonymous_gid": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "anonymous_uid": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "filter": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "root_squash_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "scope": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "submount_access_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "suid_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 3,
        "min_items": 1,
        "nesting_mode": "set"
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

func AzurermHpcCacheAccessPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermHpcCacheAccessPolicy), &result)
	return &result
}
