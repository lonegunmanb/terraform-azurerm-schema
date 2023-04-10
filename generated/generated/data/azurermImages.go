package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermImages = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
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
              "data_disk": [
                "list",
                [
                  "object",
                  {
                    "blob_uri": "string",
                    "caching": "string",
                    "lun": "number",
                    "managed_disk_id": "string",
                    "size_gb": "number"
                  }
                ]
              ],
              "location": "string",
              "name": "string",
              "os_disk": [
                "list",
                [
                  "object",
                  {
                    "blob_uri": "string",
                    "caching": "string",
                    "managed_disk_id": "string",
                    "os_state": "string",
                    "os_type": "string",
                    "size_gb": "number"
                  }
                ]
              ],
              "tags": [
                "map",
                "string"
              ],
              "zone_resilient": "bool"
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

func AzurermImagesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermImages), &result)
	return &result
}
