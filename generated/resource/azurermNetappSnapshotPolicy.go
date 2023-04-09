package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermNetappSnapshotPolicy = `{
  "block": {
    "attributes": {
      "account_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "enabled": {
        "description_kind": "plain",
        "required": true,
        "type": "bool"
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
      }
    },
    "block_types": {
      "daily_schedule": {
        "block": {
          "attributes": {
            "hour": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "minute": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "snapshots_to_keep": {
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
      "hourly_schedule": {
        "block": {
          "attributes": {
            "minute": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "snapshots_to_keep": {
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
      "monthly_schedule": {
        "block": {
          "attributes": {
            "days_of_month": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "set",
                "number"
              ]
            },
            "hour": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "minute": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "snapshots_to_keep": {
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
      "weekly_schedule": {
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
            "hour": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "minute": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "snapshots_to_keep": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AzurermNetappSnapshotPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermNetappSnapshotPolicy), &result)
	return &result
}
