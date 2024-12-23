package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermSpringCloudNewRelicApplicationPerformanceMonitoring = `{
  "block": {
    "attributes": {
      "agent_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "app_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "app_server_port": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "audit_mode_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "auto_app_naming_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "auto_transaction_naming_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "custom_tracing_enabled": {
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
      "labels": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "license_key": {
        "description_kind": "plain",
        "required": true,
        "sensitive": true,
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

func AzurermSpringCloudNewRelicApplicationPerformanceMonitoringSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermSpringCloudNewRelicApplicationPerformanceMonitoring), &result)
	return &result
}
