package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermStreamAnalyticsStreamInputEventhubV2 = `{
  "block": {
    "attributes": {
      "authentication_mode": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "eventhub_consumer_group_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "eventhub_name": {
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
      },
      "partition_key": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "servicebus_namespace": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "shared_access_policy_key": {
        "description_kind": "plain",
        "optional": true,
        "sensitive": true,
        "type": "string"
      },
      "shared_access_policy_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "stream_analytics_job_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "serialization": {
        "block": {
          "attributes": {
            "encoding": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "field_delimiter": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
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

func AzurermStreamAnalyticsStreamInputEventhubV2Schema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermStreamAnalyticsStreamInputEventhubV2), &result)
	return &result
}
