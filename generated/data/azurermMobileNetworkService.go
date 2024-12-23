package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermMobileNetworkService = `{
  "block": {
    "attributes": {
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
      "pcc_rule": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "name": "string",
              "precedence": "number",
              "qos_policy": [
                "list",
                [
                  "object",
                  {
                    "allocation_and_retention_priority_level": "number",
                    "guaranteed_bit_rate": [
                      "list",
                      [
                        "object",
                        {
                          "downlink": "string",
                          "uplink": "string"
                        }
                      ]
                    ],
                    "maximum_bit_rate": [
                      "list",
                      [
                        "object",
                        {
                          "downlink": "string",
                          "uplink": "string"
                        }
                      ]
                    ],
                    "preemption_capability": "string",
                    "preemption_vulnerability": "string",
                    "qos_indicator": "number"
                  }
                ]
              ],
              "service_data_flow_template": [
                "list",
                [
                  "object",
                  {
                    "direction": "string",
                    "name": "string",
                    "ports": [
                      "list",
                      "string"
                    ],
                    "protocol": [
                      "list",
                      "string"
                    ],
                    "remote_ip_list": [
                      "list",
                      "string"
                    ]
                  }
                ]
              ],
              "traffic_control_enabled": "bool"
            }
          ]
        ]
      },
      "service_precedence": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "service_qos_policy": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "allocation_and_retention_priority_level": "number",
              "maximum_bit_rate": [
                "list",
                [
                  "object",
                  {
                    "downlink": "string",
                    "uplink": "string"
                  }
                ]
              ],
              "preemption_capability": "string",
              "preemption_vulnerability": "string",
              "qos_indicator": "number"
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

func AzurermMobileNetworkServiceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermMobileNetworkService), &result)
	return &result
}
