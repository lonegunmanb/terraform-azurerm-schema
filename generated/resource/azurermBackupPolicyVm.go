package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermBackupPolicyVm = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "instant_restore_retention_days": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "policy_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "recovery_vault_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "resource_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "timezone": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "backup": {
        "block": {
          "attributes": {
            "frequency": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "hour_duration": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "hour_interval": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "time": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "weekdays": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "instant_restore_resource_group": {
        "block": {
          "attributes": {
            "prefix": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "suffix": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "retention_daily": {
        "block": {
          "attributes": {
            "count": {
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
      "retention_monthly": {
        "block": {
          "attributes": {
            "count": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "days": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "number"
              ]
            },
            "include_last_days": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "weekdays": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "weeks": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "retention_weekly": {
        "block": {
          "attributes": {
            "count": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "weekdays": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "set",
                "string"
              ]
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "retention_yearly": {
        "block": {
          "attributes": {
            "count": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "days": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "number"
              ]
            },
            "include_last_days": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "months": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "set",
                "string"
              ]
            },
            "weekdays": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "weeks": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
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
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AzurermBackupPolicyVmSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermBackupPolicyVm), &result)
	return &result
}
