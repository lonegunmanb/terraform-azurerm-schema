package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermElasticSan = `{
  "block": {
    "attributes": {
      "base_size_in_tib": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "extended_size_in_tib": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "location": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "resource_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "sku": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "name": "string",
              "tier": "string"
            }
          ]
        ]
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "total_iops": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "total_mbps": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "total_size_in_tib": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "total_volume_size_in_gib": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "volume_group_count": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "zones": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
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

func AzurermElasticSanSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermElasticSan), &result)
	return &result
}
