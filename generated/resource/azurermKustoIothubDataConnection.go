package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermKustoIothubDataConnection = `{
  "block": {
    "attributes": {
      "cluster_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "consumer_group": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "data_format": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "database_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "database_routing_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "event_system_properties": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "iothub_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "location": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "mapping_rule_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "resource_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "shared_access_policy_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "table_name": {
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

func AzurermKustoIothubDataConnectionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermKustoIothubDataConnection), &result)
	return &result
}
