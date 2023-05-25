package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermDataShareDatasetDataLakeGen1 = `{
  "block": {
    "attributes": {
      "data_lake_store_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "data_share_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "display_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "file_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "folder_path": {
        "computed": true,
        "description_kind": "plain",
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

func AzurermDataShareDatasetDataLakeGen1Schema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermDataShareDatasetDataLakeGen1), &result)
	return &result
}
