package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermLbBackendAddressPoolAddress = `{
  "block": {
    "attributes": {
      "backend_address_ip_configuration_id": {
        "description": "For global load balancer, user needs to specify the ` + "`" + `backend_address_ip_configuration_id` + "`" + ` of the added regional load balancers",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "backend_address_pool_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "inbound_nat_rule_port_mapping": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "backend_port": "number",
              "frontend_port": "number",
              "inbound_nat_rule_name": "string"
            }
          ]
        ]
      },
      "ip_address": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "virtual_network_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
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

func AzurermLbBackendAddressPoolAddressSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermLbBackendAddressPoolAddress), &result)
	return &result
}
