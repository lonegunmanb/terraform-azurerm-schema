package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermSpringCloudDynatraceApplicationPerformanceMonitoring = `{
  "block": {
    "attributes": {
      "api_token": {
        "description_kind": "plain",
        "optional": true,
        "sensitive": true,
        "type": "string"
      },
      "api_url": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "connection_point": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "environment_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
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
      },
      "tenant": {
        "description_kind": "plain",
        "required": true,
        "sensitive": true,
        "type": "string"
      },
      "tenant_token": {
        "description_kind": "plain",
        "required": true,
        "sensitive": true,
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

func AzurermSpringCloudDynatraceApplicationPerformanceMonitoringSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermSpringCloudDynatraceApplicationPerformanceMonitoring), &result)
	return &result
}
