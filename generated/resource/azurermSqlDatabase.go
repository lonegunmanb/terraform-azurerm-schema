package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermSqlDatabase = `{
  "block": {
    "attributes": {
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
      "creation_date": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "default_secondary_location": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "edition": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "elastic_pool_name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "encryption": {
        "computed": true,
        "description_kind": "plain",
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
              "storage_endpoint": "string"
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
      "location": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "max_size_bytes": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "max_size_gb": {
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
      "read_scale": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "requested_service_objective_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "requested_service_objective_name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "resource_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "restore_point_in_time": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "server_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "source_database_deletion_date": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "source_database_id": {
        "computed": true,
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
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      }
    },
    "block_types": {
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
            "operation_mode": {
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
  "version": 0
}`

func AzurermSqlDatabaseSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermSqlDatabase), &result)
	return &result
}
