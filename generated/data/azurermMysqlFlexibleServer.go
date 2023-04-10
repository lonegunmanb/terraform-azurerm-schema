package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermMysqlFlexibleServer = `{
  "block": {
    "attributes": {
      "administrator_login": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "backup_retention_days": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "delegated_subnet_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "fqdn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "geo_redundant_backup_enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "high_availability": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "mode": "string",
              "standby_availability_zone": "string"
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
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "maintenance_window": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "day_of_week": "number",
              "start_hour": "number",
              "start_minute": "number"
            }
          ]
        ]
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "private_dns_zone_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "public_network_access_enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "replica_capacity": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "replication_role": {
        "computed": true,
        "description_kind": "plain",
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
        "type": "string"
      },
      "sku_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "storage": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "auto_grow_enabled": "bool",
              "iops": "number",
              "size_gb": "number"
            }
          ]
        ]
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "version": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "zone": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
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

func AzurermMysqlFlexibleServerSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermMysqlFlexibleServer), &result)
	return &result
}
