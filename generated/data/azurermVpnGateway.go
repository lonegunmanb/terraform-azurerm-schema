package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermVpnGateway = `{
  "block": {
    "attributes": {
      "bgp_settings": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "asn": "number",
              "bgp_peering_address": "string",
              "instance_0_bgp_peering_address": [
                "list",
                [
                  "object",
                  {
                    "custom_ips": [
                      "list",
                      "string"
                    ],
                    "default_ips": [
                      "list",
                      "string"
                    ],
                    "ip_configuration_id": "string",
                    "tunnel_ips": [
                      "list",
                      "string"
                    ]
                  }
                ]
              ],
              "instance_1_bgp_peering_address": [
                "list",
                [
                  "object",
                  {
                    "custom_ips": [
                      "list",
                      "string"
                    ],
                    "default_ips": [
                      "list",
                      "string"
                    ],
                    "ip_configuration_id": "string",
                    "tunnel_ips": [
                      "list",
                      "string"
                    ]
                  }
                ]
              ],
              "peer_weight": "number"
            }
          ]
        ]
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
      "scale_unit": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "virtual_hub_id": {
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

func AzurermVpnGatewaySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermVpnGateway), &result)
	return &result
}
