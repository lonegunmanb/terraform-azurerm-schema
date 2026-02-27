package resource

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
        "optional": true,
        "type": "bool"
      },
      "bgp_enabled": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "bgp_route_translation_for_nat_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "default_local_network_gateway_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "dns_forwarding_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "edge_zone": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "enable_bgp": {
        "computed": true,
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "generation": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ip_sec_replay_protection_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
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
      "private_ip_address_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "remote_vnet_traffic_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "resource_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "sku": {
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
      "type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "virtual_wan_traffic_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "vpn_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "bgp_settings": {
        "block": {
          "attributes": {
            "asn": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "peer_weight": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "block_types": {
            "peering_addresses": {
              "block": {
                "attributes": {
                  "apipa_addresses": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "default_addresses": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "ip_configuration_name": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "tunnel_ip_addresses": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 2,
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "custom_route": {
        "block": {
          "attributes": {
            "address_prefixes": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "ip_configuration": {
        "block": {
          "attributes": {
            "name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "private_ip_address_allocation": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "public_ip_address_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "subnet_id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 3,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "policy_group": {
        "block": {
          "attributes": {
            "is_default": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "priority": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "block_types": {
            "policy_member": {
              "block": {
                "attributes": {
                  "name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "type": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
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
      "vpn_client_configuration": {
        "block": {
          "attributes": {
            "aad_audience": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "aad_issuer": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "aad_tenant": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "address_space": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "list",
                "string"
              ]
            },
            "radius_server_address": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "radius_server_secret": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "vpn_auth_types": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "vpn_client_protocols": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
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
                  "sa_data_size_in_kilobytes": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "sa_lifetime_in_seconds": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "radius_server": {
              "block": {
                "attributes": {
                  "address": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "score": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "secret": {
                    "description_kind": "plain",
                    "required": true,
                    "sensitive": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "revoked_certificate": {
              "block": {
                "attributes": {
                  "name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "thumbprint": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "set"
            },
            "root_certificate": {
              "block": {
                "attributes": {
                  "name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "public_cert_data": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "set"
            },
            "virtual_network_gateway_client_connection": {
              "block": {
                "attributes": {
                  "address_prefixes": {
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "policy_group_names": {
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
              "nesting_mode": "list"
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

func AzurermVirtualNetworkGatewaySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermVirtualNetworkGateway), &result)
	return &result
}
