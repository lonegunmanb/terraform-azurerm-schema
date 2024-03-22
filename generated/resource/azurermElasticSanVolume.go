package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermElasticSanVolume = `{
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
      "size_in_gib": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "target_iqn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "target_portal_hostname": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "target_portal_port": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "volume_group_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "volume_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "create_source": {
        "block": {
          "attributes": {
            "source_id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "source_type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
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

func AzurermElasticSanVolumeSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermElasticSanVolume), &result)
	return &result
}
