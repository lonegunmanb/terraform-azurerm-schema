package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermVirtualNetworkPeering = `{
  "block": {
    "attributes": {
      "allow_forwarded_traffic": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "allow_gateway_transit": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "allow_virtual_network_access": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
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
      },
      "only_ipv6_peering_enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "peer_complete_virtual_networks_enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "remote_virtual_network_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "use_remote_gateways": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "virtual_network_id": {
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

func AzurermVirtualNetworkPeeringSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermVirtualNetworkPeering), &result)
	return &result
}
