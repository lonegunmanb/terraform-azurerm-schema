package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermDigitalTwinsEndpointEventgrid = `{
  "block": {
    "attributes": {
      "dead_letter_storage_secret": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "digital_twins_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "eventgrid_topic_endpoint": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "eventgrid_topic_primary_access_key": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "eventgrid_topic_secondary_access_key": {
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
      "name": {
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

func AzurermDigitalTwinsEndpointEventgridSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermDigitalTwinsEndpointEventgrid), &result)
	return &result
}
