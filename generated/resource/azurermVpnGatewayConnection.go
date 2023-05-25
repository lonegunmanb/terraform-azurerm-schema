package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermVpnGatewayConnection = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "internet_security_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "remote_vpn_site_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "vpn_gateway_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "routing": {
        "block": {
          "attributes": {
            "associated_route_table": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "propagated_route_tables": {
              "computed": true,
              "deprecated": true,
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            }
          },
          "block_types": {
            "propagated_route_table": {
              "block": {
                "attributes": {
                  "labels": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "route_table_ids": {
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
            "local_address_ranges": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "set",
                "string"
              ]
            },
            "remote_address_ranges": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "set",
                "string"
              ]
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "vpn_link": {
        "block": {
          "attributes": {
            "bandwidth_mbps": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "bgp_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "connection_mode": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "egress_nat_rule_ids": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "ingress_nat_rule_ids": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "local_azure_ip_address_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "policy_based_traffic_selector_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "protocol": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "ratelimit_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "route_weight": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "shared_key": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "vpn_site_link_id": {
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
                  "encryption_algorithm": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "ike_encryption_algorithm": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "ike_integrity_algorithm": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "integrity_algorithm": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "pfs_group": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "sa_data_size_kb": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "sa_lifetime_sec": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
        "min_items": 1,
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AzurermVpnGatewayConnectionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermVpnGatewayConnection), &result)
	return &result
}
