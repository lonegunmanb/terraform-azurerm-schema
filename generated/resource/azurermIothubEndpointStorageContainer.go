package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermIothubEndpointStorageContainer = `{
  "block": {
    "attributes": {
      "authentication_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "batch_frequency_in_seconds": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "connection_string": {
        "description_kind": "plain",
        "optional": true,
        "sensitive": true,
        "type": "string"
      },
      "container_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "encoding": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "endpoint_uri": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "file_name_format": {
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
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "iothub_name": {
        "computed": true,
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "max_chunk_size_in_bytes": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "resource_group_name": {
        "description_kind": "plain",
        "required": true,
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

func AzurermIothubEndpointStorageContainerSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermIothubEndpointStorageContainer), &result)
	return &result
}
