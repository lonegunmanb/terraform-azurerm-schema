package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermIothubFileUpload = `{
  "block": {
    "attributes": {
      "authentication_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "connection_string": {
        "description_kind": "plain",
        "required": true,
        "sensitive": true,
        "type": "string"
      },
      "container_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "default_ttl": {
        "description_kind": "plain",
        "optional": true,
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
      "lock_duration": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "max_delivery_count": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "notifications_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "sas_ttl": {
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

func AzurermIothubFileUploadSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermIothubFileUpload), &result)
	return &result
}
