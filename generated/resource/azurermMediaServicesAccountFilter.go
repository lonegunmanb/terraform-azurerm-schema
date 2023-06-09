package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermMediaServicesAccountFilter = `{
  "block": {
    "attributes": {
      "first_quality_bitrate": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
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
      }
    },
    "block_types": {
      "presentation_time_range": {
        "block": {
          "attributes": {
            "end_in_units": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "force_end": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "live_backoff_in_units": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "presentation_window_in_units": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "start_in_units": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "unit_timescale_in_milliseconds": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
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
      },
      "track_selection": {
        "block": {
          "block_types": {
            "condition": {
              "block": {
                "attributes": {
                  "operation": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "property": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      }
    },
    "deprecated": true,
    "description_kind": "plain"
  },
  "version": 0
}`

func AzurermMediaServicesAccountFilterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermMediaServicesAccountFilter), &result)
	return &result
}
