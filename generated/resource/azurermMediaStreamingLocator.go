package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermMediaStreamingLocator = `{
  "block": {
    "attributes": {
      "alternative_media_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "asset_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "default_content_key_policy_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "end_time": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "filter_names": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "media_services_account_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
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
      },
      "start_time": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "streaming_locator_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "streaming_policy_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "content_key": {
        "block": {
          "attributes": {
            "content_key_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "label_reference_in_streaming_policy": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "policy_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
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
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "single"
      }
    },
    "deprecated": true,
    "description_kind": "plain"
  },
  "version": 1
}`

func AzurermMediaStreamingLocatorSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermMediaStreamingLocator), &result)
	return &result
}
