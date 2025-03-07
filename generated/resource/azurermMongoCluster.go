package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermMongoCluster = `{
  "block": {
    "attributes": {
      "administrator_password": {
        "description_kind": "plain",
        "optional": true,
        "sensitive": true,
        "type": "string"
      },
      "administrator_username": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "compute_tier": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "connection_strings": {
        "computed": true,
        "description_kind": "plain",
        "sensitive": true,
        "type": [
          "list",
          [
            "object",
            {
              "description": "string",
              "name": "string",
              "value": "string"
            }
          ]
        ]
      },
      "create_mode": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "high_availability_mode": {
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
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "preview_features": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "public_network_access": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "resource_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "shard_count": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "source_location": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "source_server_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "storage_size_in_gb": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "tags": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "version": {
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

func AzurermMongoClusterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermMongoCluster), &result)
	return &result
}
