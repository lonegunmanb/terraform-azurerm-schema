package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermMssqlDatabase = `{
  "block": {
    "attributes": {
      "auto_pause_delay_in_minutes": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "collation": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "create_mode": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "creation_source_database_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "elastic_pool_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "extended_auditing_policy": {
        "computed": true,
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          [
            "object",
            {
              "log_monitoring_enabled": "bool",
              "retention_in_days": "number",
              "storage_account_access_key": "string",
              "storage_account_access_key_is_secondary": "bool",
              "storage_account_subscription_id": "string",
              "storage_endpoint": "string"
            }
          ]
        ]
      },
      "geo_backup_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "license_type": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "max_size_gb": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "min_capacity": {
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
      "read_replica_count": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "read_scale": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "recover_database_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "restore_dropped_database_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "restore_point_in_time": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "sample_name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "server_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "sku_name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "storage_account_type": {
        "description_kind": "plain",
        "optional": true,
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
      "zone_redundant": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      }
    },
    "block_types": {
      "long_term_retention_policy": {
        "block": {
          "attributes": {
            "monthly_retention": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "week_of_year": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "weekly_retention": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "yearly_retention": {
              "computed": true,
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
      "short_term_retention_policy": {
        "block": {
          "attributes": {
            "retention_days": {
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
      "threat_detection_policy": {
        "block": {
          "attributes": {
            "disabled_alerts": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "email_account_admins": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "email_addresses": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "retention_days": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "state": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "storage_account_access_key": {
              "description_kind": "plain",
              "optional": true,
              "sensitive": true,
              "type": "string"
            },
            "storage_endpoint": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "use_server_default": {
              "deprecated": true,
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

func AzurermMssqlDatabaseSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermMssqlDatabase), &result)
	return &result
}
