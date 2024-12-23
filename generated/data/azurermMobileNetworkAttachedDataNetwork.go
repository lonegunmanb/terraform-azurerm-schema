package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermMobileNetworkAttachedDataNetwork = `{
  "block": {
    "attributes": {
      "dns_addresses": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
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
      "mobile_network_data_network_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "mobile_network_packet_core_data_plane_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "network_address_port_translation": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "icmp_pinhole_timeout_in_seconds": "number",
              "pinhole_maximum_number": "number",
              "port_range": [
                "list",
                [
                  "object",
                  {
                    "maximum": "number",
                    "minimum": "number"
                  }
                ]
              ],
              "tcp_pinhole_timeout_in_seconds": "number",
              "tcp_port_reuse_minimum_hold_time_in_seconds": "number",
              "udp_pinhole_timeout_in_seconds": "number",
              "udp_port_reuse_minimum_hold_time_in_seconds": "number"
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
      "user_equipment_address_pool_prefixes": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "user_equipment_static_address_pool_prefixes": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "user_plane_access_ipv4_address": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "user_plane_access_ipv4_gateway": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "user_plane_access_ipv4_subnet": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "user_plane_access_name": {
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

func AzurermMobileNetworkAttachedDataNetworkSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermMobileNetworkAttachedDataNetwork), &result)
	return &result
}
