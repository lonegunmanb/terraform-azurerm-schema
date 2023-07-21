package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermMobileNetworkAttachedDataNetwork = `{
  "block": {
    "attributes": {
      "dns_addresses": {
        "description_kind": "plain",
        "required": true,
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
        "description_kind": "plain",
        "required": true,
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
      "tags": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "user_equipment_address_pool_prefixes": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "user_equipment_static_address_pool_prefixes": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "user_plane_access_ipv4_address": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "user_plane_access_ipv4_gateway": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "user_plane_access_ipv4_subnet": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "user_plane_access_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "network_address_port_translation": {
        "block": {
          "attributes": {
            "icmp_pinhole_timeout_in_seconds": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "pinhole_maximum_number": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "tcp_pinhole_timeout_in_seconds": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "tcp_port_reuse_minimum_hold_time_in_seconds": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "udp_pinhole_timeout_in_seconds": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "udp_port_reuse_minimum_hold_time_in_seconds": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "block_types": {
            "port_range": {
              "block": {
                "attributes": {
                  "maximum": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "minimum": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
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

func AzurermMobileNetworkAttachedDataNetworkSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermMobileNetworkAttachedDataNetwork), &result)
	return &result
}
