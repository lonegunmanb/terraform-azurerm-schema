package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermKubernetesClusterExtension = `{
  "block": {
    "attributes": {
      "aks_assigned_identity": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "principal_id": "string",
              "tenant_id": "string",
              "type": "string"
            }
          ]
        ]
      },
      "cluster_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "configuration_protected_settings": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "configuration_settings": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "current_version": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "extension_type": {
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
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "release_namespace": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "release_train": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "target_namespace": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "version": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "plan": {
        "block": {
          "attributes": {
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "product": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "promotion_code": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "publisher": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "version": {
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
  "version": 0
}`

func AzurermKubernetesClusterExtensionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermKubernetesClusterExtension), &result)
	return &result
}
