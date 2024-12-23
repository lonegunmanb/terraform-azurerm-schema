package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermRedisCache = `{
  "block": {
    "attributes": {
      "access_keys_authentication_enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "capacity": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "family": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "hostname": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
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
      "minimum_tls_version": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "non_ssl_port_enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "patch_schedule": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "day_of_week": "string",
              "maintenance_window": "string",
              "start_hour_utc": "number"
            }
          ]
        ]
      },
      "port": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "primary_access_key": {
        "computed": true,
        "description_kind": "plain",
        "sensitive": true,
        "type": "string"
      },
      "primary_connection_string": {
        "computed": true,
        "description_kind": "plain",
        "sensitive": true,
        "type": "string"
      },
      "private_static_ip_address": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "redis_configuration": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "active_directory_authentication_enabled": "bool",
              "aof_backup_enabled": "bool",
              "aof_storage_connection_string_0": "string",
              "aof_storage_connection_string_1": "string",
              "authentication_enabled": "bool",
              "data_persistence_authentication_method": "string",
              "maxclients": "number",
              "maxfragmentationmemory_reserved": "number",
              "maxmemory_delta": "number",
              "maxmemory_policy": "string",
              "maxmemory_reserved": "number",
              "notify_keyspace_events": "string",
              "rdb_backup_enabled": "bool",
              "rdb_backup_frequency": "number",
              "rdb_backup_max_snapshot_count": "number",
              "rdb_storage_connection_string": "string",
              "storage_account_subscription_id": "string"
            }
          ]
        ]
      },
      "resource_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "secondary_access_key": {
        "computed": true,
        "description_kind": "plain",
        "sensitive": true,
        "type": "string"
      },
      "secondary_connection_string": {
        "computed": true,
        "description_kind": "plain",
        "sensitive": true,
        "type": "string"
      },
      "shard_count": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "sku_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "ssl_port": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "subnet_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "zones": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
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

func AzurermRedisCacheSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermRedisCache), &result)
	return &result
}
