package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermVirtualNetworkGatewayConnection = `{
  "block": {
    "attributes": {
      "authorization_key": {
        "description_kind": "plain",
        "optional": true,
        "sensitive": true,
        "type": "string"
      },
      "connection_mode": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "connection_protocol": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "dpd_timeout_seconds": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "enable_bgp": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "express_route_circuit_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "express_route_gateway_bypass": {
        "computed": true,
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
      "local_azure_ip_address_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "local_network_gateway_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "location": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "peer_virtual_network_gateway_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "resource_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "routing_weight": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "shared_key": {
        "description_kind": "plain",
        "optional": true,
        "sensitive": true,
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
      "type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "use_policy_based_traffic_selectors": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "virtual_network_gateway_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "ipsec_policy": {
        "block": {
          "attributes": {
            "dh_group": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "ike_encryption": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "ike_integrity": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "ipsec_encryption": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "ipsec_integrity": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "pfs_group": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "sa_datasize": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "sa_lifetime": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
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
      },
      "traffic_selector_policy": {
        "block": {
          "attributes": {
            "local_address_cidrs": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "list",
                "string"
              ]
            },
            "remote_address_cidrs": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "list",
                "string"
              ]
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
  "version": 0
}`

func AzurermVirtualNetworkGatewayConnectionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermVirtualNetworkGatewayConnection), &result)
	return &result
}
