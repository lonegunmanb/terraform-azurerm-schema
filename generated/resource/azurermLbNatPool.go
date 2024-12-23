package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermLbNatPool = `{
  "block": {
    "attributes": {
      "backend_port": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "floating_ip_enabled": {
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
      "frontend_port_end": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "frontend_port_start": {
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
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "resource_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tcp_reset_enabled": {
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

func AzurermLbNatPoolSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermLbNatPool), &result)
	return &result
}
