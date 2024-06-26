package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermVirtualNetworkPeering = `{
  "block": {
    "attributes": {
      "allow_forwarded_traffic": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "allow_gateway_transit": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "allow_virtual_network_access": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "local_subnet_names": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "only_ipv6_peering_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "peer_complete_virtual_networks_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "remote_subnet_names": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "remote_virtual_network_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "resource_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "triggers": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "use_remote_gateways": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "virtual_network_name": {
        "description_kind": "plain",
        "required": true,
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
  "version": 0
}`

func AzurermVirtualNetworkPeeringSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermVirtualNetworkPeering), &result)
	return &result
}
