package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermStorageSyncServerEndpoint = `{
  "block": {
    "attributes": {
      "cloud_tiering_enabled": {
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
      "initial_download_policy": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "local_cache_mode": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "registered_server_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "server_local_path": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "storage_sync_group_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tier_files_older_than_days": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "volume_free_space_percent": {
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

func AzurermStorageSyncServerEndpointSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermStorageSyncServerEndpoint), &result)
	return &result
}
