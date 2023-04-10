package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermApiManagementGatewayHostNameConfiguration = `{
  "block": {
    "attributes": {
      "api_management_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "certificate_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "gateway_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "host_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "http2_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "request_client_certificate_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "tls10_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "tls11_enabled": {
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

func AzurermApiManagementGatewayHostNameConfigurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermApiManagementGatewayHostNameConfiguration), &result)
	return &result
}
