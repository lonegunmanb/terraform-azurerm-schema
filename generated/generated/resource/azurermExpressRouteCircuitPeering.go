package resource

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
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "peer_asn": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
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
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "resource_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "route_filter_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "secondary_azure_port": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "secondary_peer_address_prefix": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "shared_key": {
        "description_kind": "plain",
        "optional": true,
        "sensitive": true,
        "type": "string"
      },
      "vlan_id": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      }
    },
    "block_types": {
      "ipv6": {
        "block": {
          "attributes": {
            "enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "primary_peer_address_prefix": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "route_filter_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "secondary_peer_address_prefix": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "microsoft_peering": {
              "block": {
                "attributes": {
                  "advertised_communities": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "advertised_public_prefixes": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "customer_asn": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "routing_registry_name": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
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
      "microsoft_peering_config": {
        "block": {
          "attributes": {
            "advertised_communities": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "advertised_public_prefixes": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "list",
                "string"
              ]
            },
            "customer_asn": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "routing_registry_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
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

func AzurermExpressRouteCircuitPeeringSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermExpressRouteCircuitPeering), &result)
	return &result
}
