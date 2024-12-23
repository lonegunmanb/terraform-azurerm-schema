package resource

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
      "service_precedence": {
        "description_kind": "plain",
        "required": true,
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
      "pcc_rule": {
        "block": {
          "attributes": {
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "precedence": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "traffic_control_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "block_types": {
            "qos_policy": {
              "block": {
                "attributes": {
                  "allocation_and_retention_priority_level": {
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
                  "guaranteed_bit_rate": {
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
                    "nesting_mode": "list"
                  },
                  "maximum_bit_rate": {
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
              "max_items": 1,
              "nesting_mode": "list"
            },
            "service_data_flow_template": {
              "block": {
                "attributes": {
                  "direction": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "ports": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "protocol": {
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "remote_ip_list": {
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "list",
                      "string"
                    ]
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
      "service_qos_policy": {
        "block": {
          "attributes": {
            "allocation_and_retention_priority_level": {
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
              "optional": true,
              "type": "number"
            }
          },
          "block_types": {
            "maximum_bit_rate": {
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
        "max_items": 1,
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
