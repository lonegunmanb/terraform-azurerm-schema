package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermHdinsightCluster = `{
  "block": {
    "attributes": {
      "cluster_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "cluster_version": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "component_versions": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "edge_ssh_endpoint": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "gateway": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "enabled": "bool",
              "password": "string",
              "username": "string"
            }
          ]
        ]
      },
      "https_endpoint": {
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
      "kafka_rest_proxy_endpoint": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "kind": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "location": {
        "computed": true,
        "description_kind": "plain",
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
      "ssh_endpoint": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "tier": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tls_min_version": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
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

func AzurermHdinsightClusterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermHdinsightCluster), &result)
	return &result
}
