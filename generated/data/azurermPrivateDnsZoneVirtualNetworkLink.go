package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermPrivateDnsZoneVirtualNetworkLink = `{
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
      "private_dns_zone_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "registration_enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "resolution_policy": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "resource_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "virtual_network_id": {
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

func AzurermPrivateDnsZoneVirtualNetworkLinkSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermPrivateDnsZoneVirtualNetworkLink), &result)
	return &result
}
