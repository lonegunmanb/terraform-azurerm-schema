package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermVirtualNetworkGatewayConnection = `{
  "block": {
    "attributes": {
      "authorization_key": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "connection_protocol": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "dpd_timeout_seconds": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "egress_bytes_transferred": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "enable_bgp": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "express_route_circuit_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "express_route_gateway_bypass": {
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
      "ingress_bytes_transferred": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "ipsec_policy": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "dh_group": "string",
              "ike_encryption": "string",
              "ike_integrity": "string",
              "ipsec_encryption": "string",
              "ipsec_integrity": "string",
              "pfs_group": "string",
              "sa_datasize": "number",
              "sa_lifetime": "number"
            }
          ]
        ]
      },
      "local_azure_ip_address_enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "local_network_gateway_id": {
        "computed": true,
        "description_kind": "plain",
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
      "peer_virtual_network_gateway_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "resource_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "resource_guid": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "routing_weight": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "shared_key": {
        "computed": true,
        "description_kind": "plain",
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
      "traffic_selector_policy": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "local_address_cidrs": [
                "list",
                "string"
              ],
              "remote_address_cidrs": [
                "list",
                "string"
              ]
            }
          ]
        ]
      },
      "type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "use_policy_based_traffic_selectors": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "virtual_network_gateway_id": {
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

func AzurermVirtualNetworkGatewayConnectionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermVirtualNetworkGatewayConnection), &result)
	return &result
}
