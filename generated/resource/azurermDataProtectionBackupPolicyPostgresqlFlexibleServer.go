package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermDataProtectionBackupPolicyPostgresqlFlexibleServer = `{
  "block": {
    "attributes": {
      "backup_repeating_time_intervals": {
        "description_kind": "plain",
        "required": true,
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
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "time_zone": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "vault_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "default_retention_rule": {
        "block": {
          "block_types": {
            "life_cycle": {
              "block": {
                "attributes": {
                  "data_store_type": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "duration": {
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
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "retention_rule": {
        "block": {
          "attributes": {
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "priority": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            }
          },
          "block_types": {
            "criteria": {
              "block": {
                "attributes": {
                  "absolute_criteria": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "days_of_week": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "months_of_year": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "scheduled_backup_times": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "weeks_of_month": {
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
            "life_cycle": {
              "block": {
                "attributes": {
                  "data_store_type": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "duration": {
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
    "description_kind": "plain"
  },
  "version": 0
}`

func AzurermDataProtectionBackupPolicyPostgresqlFlexibleServerSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermDataProtectionBackupPolicyPostgresqlFlexibleServer), &result)
	return &result
}
