package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermLbRule = `{
  "block": {
    "attributes": {
      "backend_address_pool_ids": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "backend_port": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "disable_outbound_snat": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "enable_floating_ip": {
        "computed": true,
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "enable_tcp_reset": {
        "computed": true,
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "floating_ip_enabled": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "frontend_ip_configuration_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "frontend_ip_configuration_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "frontend_port": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "idle_timeout_in_minutes": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "load_distribution": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
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
      "probe_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "protocol": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tcp_reset_enabled": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
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

func AzurermLbRuleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermLbRule), &result)
	return &result
}
