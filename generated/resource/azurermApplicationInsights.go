package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermApplicationInsights = `{
  "block": {
    "attributes": {
      "app_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "application_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "connection_string": {
        "computed": true,
        "description_kind": "plain",
        "sensitive": true,
        "type": "string"
      },
      "daily_data_cap_in_gb": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "daily_data_cap_notifications_disabled": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "disable_ip_masking": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "force_customer_storage_for_profiler": {
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
      "instrumentation_key": {
        "computed": true,
        "description_kind": "plain",
        "sensitive": true,
        "type": "string"
      },
      "internet_ingestion_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "internet_query_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "local_authentication_disabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
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
      "retention_in_days": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "sampling_percentage": {
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
      "workspace_id": {
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
  "version": 1
}`

func AzurermApplicationInsightsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermApplicationInsights), &result)
	return &result
}
