package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermAppServiceHybridConnection = `{
  "block": {
    "attributes": {
      "app_service_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "hostname": {
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
      "namespace_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "port": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "relay_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "relay_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "resource_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "send_key_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "send_key_value": {
        "computed": true,
        "description_kind": "plain",
        "sensitive": true,
        "type": "string"
      },
      "service_bus_namespace": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "service_bus_suffix": {
        "computed": true,
        "description_kind": "plain",
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

func AzurermAppServiceHybridConnectionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermAppServiceHybridConnection), &result)
	return &result
}
