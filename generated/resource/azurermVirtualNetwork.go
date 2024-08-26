package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermVirtualNetwork = `{
  "block": {
    "attributes": {
      "address_space": {
        "description_kind": "plain",
        "required": true,
        "type": [
          "set",
          "string"
        ]
      },
      "bgp_community": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "dns_servers": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "edge_zone": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "flow_timeout_in_minutes": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "guid": {
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
      "resource_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "subnet": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          [
            "object",
            {
              "address_prefixes": [
                "list",
                "string"
              ],
              "default_outbound_access_enabled": "bool",
              "delegation": [
                "list",
                [
                  "object",
                  {
                    "name": "string",
                    "service_delegation": [
                      "list",
                      [
                        "object",
                        {
                          "actions": [
                            "set",
                            "string"
                          ],
                          "name": "string"
                        }
                      ]
                    ]
                  }
                ]
              ],
              "id": "string",
              "name": "string",
              "private_endpoint_network_policies": "string",
              "private_link_service_network_policies_enabled": "bool",
              "route_table_id": "string",
              "security_group": "string",
              "service_endpoint_policy_ids": [
                "set",
                "string"
              ],
              "service_endpoints": [
                "set",
                "string"
              ]
            }
          ]
        ]
      },
      "tags": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      }
    },
    "block_types": {
      "ddos_protection_plan": {
        "block": {
          "attributes": {
            "enable": {
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            },
            "id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "encryption": {
        "block": {
          "attributes": {
            "enforcement": {
              "description_kind": "plain",
              "required": true,
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

func AzurermVirtualNetworkSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermVirtualNetwork), &result)
	return &result
}
