package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermMobileNetworkSimPolicy = `{
  "block": {
    "attributes": {
      "default_slice_id": {
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
      "location": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "mobile_network_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "rat_frequency_selection_priority_index": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "registration_timer_in_seconds": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "slice": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "data_network": [
                "list",
                [
                  "object",
                  {
                    "additional_allowed_session_types": [
                      "list",
                      "string"
                    ],
                    "allocation_and_retention_priority_level": "number",
                    "allowed_services_ids": [
                      "list",
                      "string"
                    ],
                    "data_network_id": "string",
                    "default_session_type": "string",
                    "max_buffered_packets": "number",
                    "preemption_capability": "string",
                    "preemption_vulnerability": "string",
                    "qos_indicator": "number",
                    "session_aggregate_maximum_bit_rate": [
                      "list",
                      [
                        "object",
                        {
                          "downlink": "string",
                          "uplink": "string"
                        }
                      ]
                    ]
                  }
                ]
              ],
              "default_data_network_id": "string",
              "slice_id": "string"
            }
          ]
        ]
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "user_equipment_aggregate_maximum_bit_rate": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "downlink": "string",
              "uplink": "string"
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

func AzurermMobileNetworkSimPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermMobileNetworkSimPolicy), &result)
	return &result
}
