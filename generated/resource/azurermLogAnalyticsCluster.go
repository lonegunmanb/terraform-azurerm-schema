package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermLogAnalyticsCluster = `{
  "block": {
    "attributes": {
      "cluster_id": {
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
      "size_gb": {
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
        "min_items": 1,
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

func AzurermLogAnalyticsClusterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermLogAnalyticsCluster), &result)
	return &result
}
