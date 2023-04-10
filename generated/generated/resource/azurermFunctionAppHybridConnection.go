package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermFunctionAppHybridConnection = `{
  "block": {
    "attributes": {
      "function_app_id": {
        "description": "The ID of the Function App for this Hybrid Connection.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "hostname": {
        "description": "The hostname of the endpoint.",
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
        "description": "The name of the Relay Namespace.",
        "description_kind": "plain",
        "type": "string"
      },
      "port": {
        "description": "The port to use for the endpoint",
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "relay_id": {
        "description": "The ID of the Relay Hybrid Connection to use.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "relay_name": {
        "computed": true,
        "description": "The name of the Relay in use.",
        "description_kind": "plain",
        "type": "string"
      },
      "send_key_name": {
        "description": "The name of the Relay key with ` + "`" + `Send` + "`" + ` permission to use. Defaults to ` + "`" + `RootManageSharedAccessKey` + "`" + `",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "send_key_value": {
        "computed": true,
        "description": "The Primary Access Key for the ` + "`" + `send_key_name` + "`" + `",
        "description_kind": "plain",
        "sensitive": true,
        "type": "string"
      },
      "service_bus_namespace": {
        "computed": true,
        "description": "The Service Bus Namespace.",
        "description_kind": "plain",
        "type": "string"
      },
      "service_bus_suffix": {
        "computed": true,
        "description": "The suffix for the endpoint.",
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

func AzurermFunctionAppHybridConnectionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermFunctionAppHybridConnection), &result)
	return &result
}
