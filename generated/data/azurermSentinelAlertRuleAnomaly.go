package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermSentinelAlertRuleAnomaly = `{
  "block": {
    "attributes": {
      "anomaly_settings_version": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "anomaly_version": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "display_name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "frequency": {
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
      "log_analytics_workspace_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "mode": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "multi_select_observation": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "description": "string",
              "name": "string",
              "supported_values": [
                "list",
                "string"
              ],
              "values": [
                "list",
                "string"
              ]
            }
          ]
        ]
      },
      "name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "prioritized_exclude_observation": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "description": "string",
              "exclude": "string",
              "name": "string",
              "prioritize": "string"
            }
          ]
        ]
      },
      "required_data_connector": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "connector_id": "string",
              "data_types": [
                "list",
                "string"
              ]
            }
          ]
        ]
      },
      "settings_definition_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "single_select_observation": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "description": "string",
              "name": "string",
              "supported_values": [
                "list",
                "string"
              ],
              "value": "string"
            }
          ]
        ]
      },
      "tactics": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "techniques": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "threshold_observation": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "description": "string",
              "max": "string",
              "min": "string",
              "name": "string",
              "value": "string"
            }
          ]
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

func AzurermSentinelAlertRuleAnomalySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermSentinelAlertRuleAnomaly), &result)
	return &result
}
