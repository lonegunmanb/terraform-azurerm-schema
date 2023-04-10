package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermSpringCloudGatewayRouteConfig = `{
  "block": {
    "attributes": {
      "filters": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
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
      "predicates": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "protocol": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "spring_cloud_app_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "spring_cloud_gateway_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "sso_validation_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      }
    },
    "block_types": {
      "open_api": {
        "block": {
          "attributes": {
            "uri": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "route": {
        "block": {
          "attributes": {
            "classification_tags": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "description": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "filters": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "order": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "predicates": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "sso_validation_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "title": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "token_relay": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "uri": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
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
  "version": 1
}`

func AzurermSpringCloudGatewayRouteConfigSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermSpringCloudGatewayRouteConfig), &result)
	return &result
}
