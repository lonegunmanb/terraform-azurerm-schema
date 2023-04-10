package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermNetworkInterface = `{
  "block": {
    "attributes": {
      "applied_dns_servers": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "dns_servers": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "enable_accelerated_networking": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "enable_ip_forwarding": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "internal_dns_name_label": {
        "computed": true,
        "description_kind": "plain",
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
              "application_gateway_backend_address_pools_ids": [
                "set",
                "string"
              ],
              "application_security_group_ids": [
                "set",
                "string"
              ],
              "gateway_load_balancer_frontend_ip_configuration_id": "string",
              "load_balancer_backend_address_pools_ids": [
                "set",
                "string"
              ],
              "load_balancer_inbound_nat_rules_ids": [
                "set",
                "string"
              ],
              "name": "string",
              "primary": "bool",
              "private_ip_address": "string",
              "private_ip_address_allocation": "string",
              "private_ip_address_version": "string",
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
      "mac_address": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "network_security_group_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "private_ip_address": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "private_ip_addresses": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "resource_group_name": {
        "description_kind": "plain",
        "required": true,
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
      "virtual_machine_id": {
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

func AzurermNetworkInterfaceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermNetworkInterface), &result)
	return &result
}
