package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermKubernetesClusterNodePool = `{
  "block": {
    "attributes": {
      "auto_scaling_enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "eviction_policy": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "gpu_driver": {
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
      "kubernetes_cluster_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "max_count": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "max_pods": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "min_count": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "mode": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "node_count": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "node_labels": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "node_public_ip_enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "node_public_ip_prefix_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "node_taints": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "orchestrator_version": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "os_disk_size_gb": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "os_disk_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "os_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "priority": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "proximity_placement_group_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "resource_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "spot_max_price": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "upgrade_settings": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "drain_timeout_in_minutes": "number",
              "max_surge": "string",
              "node_soak_duration_in_minutes": "number"
            }
          ]
        ]
      },
      "vm_size": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "vnet_subnet_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "zones": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      }
    },
    "block_types": {
      "timeouts": {
        "block": {
          "attributes": {
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

func AzurermKubernetesClusterNodePoolSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermKubernetesClusterNodePool), &result)
	return &result
}
