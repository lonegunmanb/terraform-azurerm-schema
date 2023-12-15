package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermRedisCache = `{
  "block": {
    "attributes": {
      "capacity": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "enable_non_ssl_port": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "family": {
        "description_kind": "plain",
        "required": true,
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
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "minimum_tls_version": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
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
        "optional": true,
        "type": "string"
      },
      "public_network_access_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "redis_version": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "replicas_per_master": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "replicas_per_primary": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
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
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "sku_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "ssl_port": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "subnet_id": {
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
      "tenant_settings": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "zones": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      }
    },
    "block_types": {
      "identity": {
        "block": {
          "attributes": {
            "identity_ids": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "principal_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "tenant_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
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
      "patch_schedule": {
        "block": {
          "attributes": {
            "day_of_week": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "maintenance_window": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "start_hour_utc": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "redis_configuration": {
        "block": {
          "attributes": {
            "active_directory_authentication_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "aof_backup_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "aof_storage_connection_string_0": {
              "description_kind": "plain",
              "optional": true,
              "sensitive": true,
              "type": "string"
            },
            "aof_storage_connection_string_1": {
              "description_kind": "plain",
              "optional": true,
              "sensitive": true,
              "type": "string"
            },
            "enable_authentication": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "maxclients": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "maxfragmentationmemory_reserved": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "maxmemory_delta": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "maxmemory_policy": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "maxmemory_reserved": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "notify_keyspace_events": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "rdb_backup_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "rdb_backup_frequency": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "rdb_backup_max_snapshot_count": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "rdb_storage_connection_string": {
              "description_kind": "plain",
              "optional": true,
              "sensitive": true,
              "type": "string"
            },
            "storage_account_subscription_id": {
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

func AzurermRedisCacheSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermRedisCache), &result)
	return &result
}
