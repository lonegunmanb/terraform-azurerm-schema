package data

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
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "gateway_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "host_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "http2_enabled": {
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
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "request_client_certificate_enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "tls10_enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "tls11_enabled": {
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

func AzurermApiManagementGatewayHostNameConfigurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermApiManagementGatewayHostNameConfiguration), &result)
	return &result
}
