package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermExpressRouteCircuit = `{
  "block": {
    "attributes": {
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
      "peerings": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "azure_asn": "number",
              "peer_asn": "number",
              "peering_type": "string",
              "primary_peer_address_prefix": "string",
              "secondary_peer_address_prefix": "string",
              "shared_key": "string",
              "vlan_id": "number"
            }
          ]
        ]
      },
      "resource_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "service_key": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "service_provider_properties": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "bandwidth_in_mbps": "number",
              "peering_location": "string",
              "service_provider_name": "string"
            }
          ]
        ]
      },
      "service_provider_provisioning_state": {
        "computed": true,
        "description_kind": "plain",
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
              "family": "string",
              "tier": "string"
            }
          ]
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

func AzurermExpressRouteCircuitSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermExpressRouteCircuit), &result)
	return &result
}
