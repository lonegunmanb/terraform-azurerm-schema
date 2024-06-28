package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermNetworkConnectionMonitor = `{
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
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "network_watcher_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "notes": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "output_workspace_resource_ids": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
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
      "endpoint": {
        "block": {
          "attributes": {
            "address": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "coverage_level": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "excluded_ip_addresses": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "included_ip_addresses": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "target_resource_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "target_resource_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "filter": {
              "block": {
                "attributes": {
                  "type": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "item": {
                    "block": {
                      "attributes": {
                        "address": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "type": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "nesting_mode": "set"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
        "min_items": 1,
        "nesting_mode": "set"
      },
      "test_configuration": {
        "block": {
          "attributes": {
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "preferred_ip_version": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "protocol": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "test_frequency_in_seconds": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "block_types": {
            "http_configuration": {
              "block": {
                "attributes": {
                  "method": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "path": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "port": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "prefer_https": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "valid_status_code_ranges": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  }
                },
                "block_types": {
                  "request_header": {
                    "block": {
                      "attributes": {
                        "name": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "value": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "nesting_mode": "set"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "icmp_configuration": {
              "block": {
                "attributes": {
                  "trace_route_enabled": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "success_threshold": {
              "block": {
                "attributes": {
                  "checks_failed_percent": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "round_trip_time_ms": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "tcp_configuration": {
              "block": {
                "attributes": {
                  "destination_port_behavior": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "port": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "trace_route_enabled": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
        "min_items": 1,
        "nesting_mode": "set"
      },
      "test_group": {
        "block": {
          "attributes": {
            "destination_endpoints": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "set",
                "string"
              ]
            },
            "enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "source_endpoints": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "set",
                "string"
              ]
            },
            "test_configuration_names": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "set",
                "string"
              ]
            }
          },
          "description_kind": "plain"
        },
        "min_items": 1,
        "nesting_mode": "set"
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

func AzurermNetworkConnectionMonitorSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermNetworkConnectionMonitor), &result)
	return &result
}
