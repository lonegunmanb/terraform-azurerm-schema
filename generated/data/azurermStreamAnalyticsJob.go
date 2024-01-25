package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermStreamAnalyticsJob = `{
  "block": {
    "attributes": {
      "compatibility_level": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "data_locale": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "events_late_arrival_max_delay_in_seconds": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "events_out_of_order_max_delay_in_seconds": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "events_out_of_order_policy": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "identity": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "principal_id": "string",
              "tenant_id": "string",
              "type": "string"
            }
          ]
        ]
      },
      "job_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "last_output_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "location": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "output_error_policy": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "resource_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "sku_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "start_mode": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "start_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "streaming_units": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "transformation_query": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
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

func AzurermStreamAnalyticsJobSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermStreamAnalyticsJob), &result)
	return &result
}
