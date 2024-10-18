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
      "enclave_type": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
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
      "ledger_enabled": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "license_type": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "maintenance_configuration_name": {
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
      "recovery_point_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "restore_dropped_database_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "restore_long_term_retention_backup_id": {
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
      "secondary_type": {
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
      "transparent_data_encryption_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "transparent_data_encryption_key_automatic_rotation_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "transparent_data_encryption_key_vault_key_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "zone_redundant": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      }
    },
    "block_types": {
      "identity": {
        "block": {
          "attributes": {
            "identity_ids": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "set",
                "string"
              ]
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
        "nesting_mode": "list"
      },
      "import": {
        "block": {
          "attributes": {
            "administrator_login": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "administrator_login_password": {
              "description_kind": "plain",
              "required": true,
              "sensitive": true,
              "type": "string"
            },
            "authentication_type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "storage_account_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "storage_key": {
              "description_kind": "plain",
              "required": true,
              "sensitive": true,
              "type": "string"
            },
            "storage_key_type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "storage_uri": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
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
            "backup_interval_in_hours": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
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
