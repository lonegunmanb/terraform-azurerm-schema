package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermLbRule = `{
  "block": {
    "attributes": {
      "backend_address_pool_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "backend_port": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "disable_outbound_snat": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "enable_floating_ip": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "enable_tcp_reset": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "frontend_ip_configuration_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "frontend_port": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
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
      "load_distribution": {
        "computed": true,
        "description_kind": "plain",
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
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "protocol": {
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

func AzurermLbRuleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermLbRule), &result)
	return &result
}
