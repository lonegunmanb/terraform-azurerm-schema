package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermStorageBlob = `{
  "block": {
    "attributes": {
      "access_tier": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "cache_control": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "content_md5": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "content_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "encryption_scope": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "metadata": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "parallelism": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "size": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "source": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "source_content": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "source_uri": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "storage_account_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "storage_container_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "url": {
        "computed": true,
        "description_kind": "plain",
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
  "version": 1
}`

func AzurermStorageBlobSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermStorageBlob), &result)
	return &result
}
