package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermStorageObjectReplication = `{
  "block": {
    "attributes": {
      "destination_object_replication_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "destination_storage_account_id": {
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
      "source_object_replication_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "source_storage_account_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "rules": {
        "block": {
          "attributes": {
            "copy_blobs_created_after": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "destination_container_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "filter_out_blobs_with_prefix": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "source_container_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
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

func AzurermStorageObjectReplicationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermStorageObjectReplication), &result)
	return &result
}
