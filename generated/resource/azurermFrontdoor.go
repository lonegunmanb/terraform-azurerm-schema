package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermFrontdoor = `{
  "block": {
    "attributes": {
      "backend_pool_health_probes": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "backend_pool_load_balancing_settings": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "backend_pools": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "backend_pools_send_receive_timeout_seconds": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "cname": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "enforce_backend_pools_certificate_name_check": {
        "description_kind": "plain",
        "required": true,
        "type": "bool"
      },
      "explicit_resource_order": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "backend_pool_health_probe_ids": [
                "list",
                "string"
              ],
              "backend_pool_ids": [
                "list",
                "string"
              ],
              "backend_pool_load_balancing_ids": [
                "list",
                "string"
              ],
              "frontend_endpoint_ids": [
                "list",
                "string"
              ],
              "routing_rule_ids": [
                "list",
                "string"
              ]
            }
          ]
        ]
      },
      "friendly_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "frontend_endpoints": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "header_frontdoor_id": {
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
      "load_balancer_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "location": {
        "computed": true,
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
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
      "routing_rules": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
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
      "backend_pool": {
        "block": {
          "attributes": {
            "health_probe_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "load_balancing_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "backend": {
              "block": {
                "attributes": {
                  "address": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "enabled": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "host_header": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "http_port": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "https_port": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "priority": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "weight": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 500,
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
        "min_items": 1,
        "nesting_mode": "list"
      },
      "backend_pool_health_probe": {
        "block": {
          "attributes": {
            "enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "interval_in_seconds": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "path": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "probe_method": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "protocol": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 5000,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "backend_pool_load_balancing": {
        "block": {
          "attributes": {
            "additional_latency_milliseconds": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "sample_size": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "successful_samples_required": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 5000,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "frontend_endpoint": {
        "block": {
          "attributes": {
            "host_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "session_affinity_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "session_affinity_ttl_seconds": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "web_application_firewall_policy_link_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 500,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "routing_rule": {
        "block": {
          "attributes": {
            "accepted_protocols": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "list",
                "string"
              ]
            },
            "enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "frontend_endpoints": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "list",
                "string"
              ]
            },
            "id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "patterns_to_match": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "list",
                "string"
              ]
            }
          },
          "block_types": {
            "forwarding_configuration": {
              "block": {
                "attributes": {
                  "backend_pool_name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "cache_duration": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "cache_enabled": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "cache_query_parameter_strip_directive": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "cache_query_parameters": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "cache_use_dynamic_compression": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "custom_forwarding_path": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "forwarding_protocol": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "redirect_configuration": {
              "block": {
                "attributes": {
                  "custom_fragment": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "custom_host": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "custom_path": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "custom_query_string": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "redirect_protocol": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "redirect_type": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
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
        "max_items": 500,
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
      }
    },
    "description_kind": "plain"
  },
  "version": 2
}`

func AzurermFrontdoorSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermFrontdoor), &result)
	return &result
}
