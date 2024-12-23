package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermCosmosdbCassandraDatacenter = `{
  "block": {
    "attributes": {
      "availability_zones_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "backup_storage_customer_key_uri": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "base64_encoded_yaml_fragment": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "cassandra_cluster_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "delegated_management_subnet_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "disk_count": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "disk_sku": {
        "description_kind": "plain",
        "optional": true,
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
      "managed_disk_customer_key_uri": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "node_count": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "seed_node_ip_addresses": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "sku_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
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

func AzurermCosmosdbCassandraDatacenterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermCosmosdbCassandraDatacenter), &result)
	return &result
}
