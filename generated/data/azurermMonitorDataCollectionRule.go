package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermMonitorDataCollectionRule = `{
  "block": {
    "attributes": {
      "data_flow": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "destinations": [
                "list",
                "string"
              ],
              "streams": [
                "list",
                "string"
              ]
            }
          ]
        ]
      },
      "data_sources": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "extension": [
                "list",
                [
                  "object",
                  {
                    "extension_json": "string",
                    "extension_name": "string",
                    "input_data_sources": [
                      "list",
                      "string"
                    ],
                    "name": "string",
                    "streams": [
                      "list",
                      "string"
                    ]
                  }
                ]
              ],
              "performance_counter": [
                "list",
                [
                  "object",
                  {
                    "counter_specifiers": [
                      "list",
                      "string"
                    ],
                    "name": "string",
                    "sampling_frequency_in_seconds": "number",
                    "streams": [
                      "list",
                      "string"
                    ]
                  }
                ]
              ],
              "syslog": [
                "list",
                [
                  "object",
                  {
                    "facility_names": [
                      "list",
                      "string"
                    ],
                    "log_levels": [
                      "list",
                      "string"
                    ],
                    "name": "string",
                    "streams": [
                      "list",
                      "string"
                    ]
                  }
                ]
              ],
              "windows_event_log": [
                "list",
                [
                  "object",
                  {
                    "name": "string",
                    "streams": [
                      "list",
                      "string"
                    ],
                    "x_path_queries": [
                      "list",
                      "string"
                    ]
                  }
                ]
              ]
            }
          ]
        ]
      },
      "description": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "destinations": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "azure_monitor_metrics": [
                "list",
                [
                  "object",
                  {
                    "name": "string"
                  }
                ]
              ],
              "log_analytics": [
                "list",
                [
                  "object",
                  {
                    "name": "string",
                    "workspace_resource_id": "string"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
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
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
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

func AzurermMonitorDataCollectionRuleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermMonitorDataCollectionRule), &result)
	return &result
}
