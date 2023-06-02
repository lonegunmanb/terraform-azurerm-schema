package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermSharedImageVersions = `{
  "block": {
    "attributes": {
      "gallery_name": {
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
      "image_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "images": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "exclude_from_latest": "bool",
              "id": "string",
              "location": "string",
              "managed_image_id": "string",
              "name": "string",
              "tags": [
                "map",
                "string"
              ],
              "target_region": [
                "list",
                [
                  "object",
                  {
                    "name": "string",
                    "regional_replica_count": "number",
                    "storage_account_type": "string"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "resource_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags_filter": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      }
    },
    "block_types": {
      "timeouts": {
        "block": {
          "attributes": {
            "read": {
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

func AzurermSharedImageVersionsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermSharedImageVersions), &result)
	return &result
}
