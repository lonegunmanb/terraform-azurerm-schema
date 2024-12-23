package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermVirtualNetworkGateway = `{
  "block": {
    "attributes": {
      "active_active": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "bgp_settings": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "asn": "number",
              "peer_weight": "number",
              "peering_address": "string"
            }
          ]
        ]
      },
      "custom_route": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "address_prefixes": [
                "set",
                "string"
              ]
            }
          ]
        ]
      },
      "default_local_network_gateway_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "enable_bgp": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "generation": {
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
      "ip_configuration": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "id": "string",
              "name": "string",
              "private_ip_address": "string",
              "private_ip_address_allocation": "string",
              "public_ip_address_id": "string",
              "subnet_id": "string"
            }
          ]
        ]
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
      "private_ip_address_enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "resource_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "sku": {
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
      "type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "vpn_client_configuration": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "aad_audience": "string",
              "aad_issuer": "string",
              "aad_tenant": "string",
              "address_space": [
                "list",
                "string"
              ],
              "radius_server_address": "string",
              "radius_server_secret": "string",
              "revoked_certificate": [
                "list",
                [
                  "object",
                  {
                    "name": "string",
                    "thumbprint": "string"
                  }
                ]
              ],
              "root_certificate": [
                "list",
                [
                  "object",
                  {
                    "name": "string",
                    "public_cert_data": "string"
                  }
                ]
              ],
              "vpn_client_protocols": [
                "set",
                "string"
              ]
            }
          ]
        ]
      },
      "vpn_type": {
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

func AzurermVirtualNetworkGatewaySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermVirtualNetworkGateway), &result)
	return &result
}
