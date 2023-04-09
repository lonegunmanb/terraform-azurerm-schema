package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermMobileNetworkSimPolicy = `{
  "block": {
    "attributes": {
      "default_slice_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "location": {
        "description_kind": "plain",
        "required": true,
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
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "registration_timer_in_seconds": {
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
      }
    },
    "block_types": {
      "slice": {
        "block": {
          "attributes": {
            "default_data_network_id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "slice_id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "data_network": {
              "block": {
                "attributes": {
                  "additional_allowed_session_types": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "allocation_and_retention_priority_level": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "allowed_services_ids": {
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "data_network_id": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "default_session_type": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "max_buffered_packets": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "preemption_capability": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "preemption_vulnerability": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "qos_indicator": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  }
                },
                "block_types": {
                  "session_aggregate_maximum_bit_rate": {
                    "block": {
                      "attributes": {
                        "downlink": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "uplink": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "min_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description_kind": "plain"
              },
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
        "min_items": 1,
        "nesting_mode": "list"
      },
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
      },
      "user_equipment_aggregate_maximum_bit_rate": {
        "block": {
          "attributes": {
            "downlink": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "uplink": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
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
