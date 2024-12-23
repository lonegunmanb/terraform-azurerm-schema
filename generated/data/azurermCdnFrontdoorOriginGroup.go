package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermCdnFrontdoorOriginGroup = `{
  "block": {
    "attributes": {
      "cdn_frontdoor_profile_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "health_probe": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "interval_in_seconds": "number",
              "path": "string",
              "protocol": "string",
              "request_type": "string"
            }
          ]
        ]
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "load_balancing": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "additional_latency_in_milliseconds": "number",
              "sample_size": "number",
              "successful_samples_required": "number"
            }
          ]
        ]
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "profile_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "resource_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "restore_traffic_time_to_healed_or_new_endpoint_in_minutes": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "session_affinity_enabled": {
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

func AzurermCdnFrontdoorOriginGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermCdnFrontdoorOriginGroup), &result)
	return &result
}
