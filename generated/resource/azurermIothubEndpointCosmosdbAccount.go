package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermIothubEndpointCosmosdbAccount = `{
  "block": {
    "attributes": {
      "authentication_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "container_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "database_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "endpoint_uri": {
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
      "identity_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "iothub_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "partition_key_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "partition_key_template": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "primary_key": {
        "description_kind": "plain",
        "optional": true,
        "sensitive": true,
        "type": "string"
      },
      "resource_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "secondary_key": {
        "description_kind": "plain",
        "optional": true,
        "sensitive": true,
        "type": "string"
      },
      "subscription_id": {
        "computed": true,
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

func AzurermIothubEndpointCosmosdbAccountSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermIothubEndpointCosmosdbAccount), &result)
	return &result
}
