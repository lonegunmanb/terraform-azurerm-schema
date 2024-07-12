package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermMachineLearningComputeCluster = `{
  "block": {
    "attributes": {
      "description": {
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
      "local_auth_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "location": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "machine_learning_workspace_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "node_public_ip_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "ssh_public_access_enabled": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "subnet_resource_id": {
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
      "vm_priority": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "vm_size": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
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
      "scale_settings": {
        "block": {
          "attributes": {
            "max_node_count": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "min_node_count": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "scale_down_nodes_after_idle_duration": {
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
      "ssh": {
        "block": {
          "attributes": {
            "admin_password": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "admin_username": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "key_value": {
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

func AzurermMachineLearningComputeClusterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermMachineLearningComputeCluster), &result)
	return &result
}
