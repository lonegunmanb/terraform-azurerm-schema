package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermCosmosdbPostgresqlCluster = `{
  "block": {
    "attributes": {
      "administrator_login_password": {
        "description_kind": "plain",
        "optional": true,
        "sensitive": true,
        "type": "string"
      },
      "citus_version": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "coordinator_public_ip_access_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "coordinator_server_edition": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "coordinator_storage_quota_in_mb": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "coordinator_vcore_count": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "earliest_restore_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "ha_enabled": {
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
      "node_count": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "node_public_ip_access_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "node_server_edition": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "node_storage_quota_in_mb": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "node_vcores": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "point_in_time_in_utc": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "preferred_primary_zone": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "resource_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "shards_on_coordinator_enabled": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "source_location": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "source_resource_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "sql_version": {
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
      }
    },
    "block_types": {
      "maintenance_window": {
        "block": {
          "attributes": {
            "day_of_week": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "start_hour": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "start_minute": {
              "description_kind": "plain",
              "optional": true,
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
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AzurermCosmosdbPostgresqlClusterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermCosmosdbPostgresqlCluster), &result)
	return &result
}
