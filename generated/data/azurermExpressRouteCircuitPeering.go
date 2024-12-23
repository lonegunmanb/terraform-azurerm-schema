package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermExpressRouteCircuitPeering = `{
  "block": {
    "attributes": {
      "azure_asn": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "express_route_circuit_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "gateway_manager_etag": {
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
      "ipv4_enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "peer_asn": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "peering_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "primary_azure_port": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "primary_peer_address_prefix": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "resource_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "route_filter_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "secondary_azure_port": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "secondary_peer_address_prefix": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "shared_key": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "vlan_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
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

func AzurermExpressRouteCircuitPeeringSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermExpressRouteCircuitPeering), &result)
	return &result
}
