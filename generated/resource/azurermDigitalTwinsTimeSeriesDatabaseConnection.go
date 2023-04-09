package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermDigitalTwinsTimeSeriesDatabaseConnection = `{
  "block": {
    "attributes": {
      "digital_twins_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "eventhub_consumer_group_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "eventhub_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "eventhub_namespace_endpoint_uri": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "eventhub_namespace_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "kusto_cluster_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "kusto_cluster_uri": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "kusto_database_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "kusto_table_name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
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

func AzurermDigitalTwinsTimeSeriesDatabaseConnectionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermDigitalTwinsTimeSeriesDatabaseConnection), &result)
	return &result
}
