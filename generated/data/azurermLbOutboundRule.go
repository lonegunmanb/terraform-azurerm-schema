package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermLbOutboundRule = `{
  "block": {
    "attributes": {
      "allocated_outbound_ports": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "backend_address_pool_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "enable_tcp_reset": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "frontend_ip_configuration": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "id": "string",
              "name": "string"
            }
          ]
        ]
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "idle_timeout_in_minutes": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "loadbalancer_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "protocol": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tcp_reset_enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
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

func AzurermLbOutboundRuleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermLbOutboundRule), &result)
	return &result
}
