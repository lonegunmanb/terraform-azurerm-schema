package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermSpringCloudAppCosmosdbAssociation = `{
  "block": {
    "attributes": {
      "api_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "cosmosdb_access_key": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "cosmosdb_account_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "cosmosdb_cassandra_keyspace_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "cosmosdb_gremlin_database_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "cosmosdb_gremlin_graph_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "cosmosdb_mongo_database_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "cosmosdb_sql_database_name": {
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
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "spring_cloud_app_id": {
        "description_kind": "plain",
        "required": true,
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
  "version": 1
}`

func AzurermSpringCloudAppCosmosdbAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermSpringCloudAppCosmosdbAssociation), &result)
	return &result
}
