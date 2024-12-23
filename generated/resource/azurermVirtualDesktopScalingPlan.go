package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermVirtualDesktopScalingPlan = `{
  "block": {
    "attributes": {
      "description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "exclusion_tag": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "friendly_name": {
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
      "location": {
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
      "tags": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "time_zone": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "host_pool": {
        "block": {
          "attributes": {
            "hostpool_id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "scaling_plan_enabled": {
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "schedule": {
        "block": {
          "attributes": {
            "days_of_week": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "set",
                "string"
              ]
            },
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "off_peak_load_balancing_algorithm": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "off_peak_start_time": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "peak_load_balancing_algorithm": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "peak_start_time": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "ramp_down_capacity_threshold_percent": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "ramp_down_force_logoff_users": {
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            },
            "ramp_down_load_balancing_algorithm": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "ramp_down_minimum_hosts_percent": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "ramp_down_notification_message": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "ramp_down_start_time": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "ramp_down_stop_hosts_when": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "ramp_down_wait_time_minutes": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "ramp_up_capacity_threshold_percent": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "ramp_up_load_balancing_algorithm": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "ramp_up_minimum_hosts_percent": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "ramp_up_start_time": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
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
  "version": 0
}`

func AzurermVirtualDesktopScalingPlanSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermVirtualDesktopScalingPlan), &result)
	return &result
}
