package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermElasticSanVolumeSnapshot = `{
  "block": {
    "attributes": {
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
      "source_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "source_volume_size_in_gib": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "volume_group_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "volume_name": {
        "computed": true,
        "description_kind": "plain",
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

func AzurermElasticSanVolumeSnapshotSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermElasticSanVolumeSnapshot), &result)
	return &result
}
