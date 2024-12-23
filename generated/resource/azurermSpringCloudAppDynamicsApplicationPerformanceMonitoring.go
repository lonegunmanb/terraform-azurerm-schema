package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermSpringCloudAppDynamicsApplicationPerformanceMonitoring = `{
  "block": {
    "attributes": {
      "agent_account_access_key": {
        "description_kind": "plain",
        "required": true,
        "sensitive": true,
        "type": "string"
      },
      "agent_account_name": {
        "description_kind": "plain",
        "required": true,
        "sensitive": true,
        "type": "string"
      },
      "agent_application_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "agent_node_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "agent_tier_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "agent_unique_host_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "controller_host_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "controller_port": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "controller_ssl_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "globally_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
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
      "spring_cloud_service_id": {
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

func AzurermSpringCloudAppDynamicsApplicationPerformanceMonitoringSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermSpringCloudAppDynamicsApplicationPerformanceMonitoring), &result)
	return &result
}
