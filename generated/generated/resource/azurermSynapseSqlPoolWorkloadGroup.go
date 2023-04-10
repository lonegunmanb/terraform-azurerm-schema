package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermSynapseSqlPoolWorkloadGroup = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "importance": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "max_resource_percent": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "max_resource_percent_per_request": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "min_resource_percent": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "min_resource_percent_per_request": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "query_execution_timeout_in_seconds": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "sql_pool_id": {
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
  "version": 0
}`

func AzurermSynapseSqlPoolWorkloadGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermSynapseSqlPoolWorkloadGroup), &result)
	return &result
}
