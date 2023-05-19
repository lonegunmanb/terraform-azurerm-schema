package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermVirtualHubConnection = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "internet_security_enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "remote_virtual_network_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "resource_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "routing": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "associated_route_table_id": "string",
              "propagated_route_table": [
                "list",
                [
                  "object",
                  {
                    "labels": [
                      "list",
                      "string"
                    ],
                    "route_table_ids": [
                      "list",
                      "string"
                    ]
                  }
                ]
              ],
              "static_vnet_route": [
                "list",
                [
                  "object",
                  {
                    "address_prefixes": [
                      "list",
                      "string"
                    ],
                    "name": "string",
                    "next_hop_ip_address": "string"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "virtual_hub_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "virtual_hub_name": {
        "description_kind": "plain",
        "required": true,
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

func AzurermVirtualHubConnectionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermVirtualHubConnection), &result)
	return &result
}
