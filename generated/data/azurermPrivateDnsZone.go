package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermPrivateDnsZone = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "max_number_of_record_sets": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "max_number_of_virtual_network_links": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "max_number_of_virtual_network_links_with_registration": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "number_of_record_sets": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "resource_group_name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tags": {
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

func AzurermPrivateDnsZoneSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermPrivateDnsZone), &result)
	return &result
}
