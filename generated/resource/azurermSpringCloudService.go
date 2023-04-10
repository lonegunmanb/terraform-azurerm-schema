package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermSpringCloudService = `{
  "block": {
    "attributes": {
      "build_agent_pool_size": {
        "description_kind": "plain",
        "optional": true,
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
      "log_stream_public_endpoint_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "outbound_public_ip_addresses": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "required_network_traffic_rules": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "direction": "string",
              "fqdns": [
                "list",
                "string"
              ],
              "ip_addresses": [
                "list",
                "string"
              ],
              "port": "number",
              "protocol": "string"
            }
          ]
        ]
      },
      "resource_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "service_registry_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "service_registry_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "sku_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tags": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "zone_redundant": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      }
    },
    "block_types": {
      "config_server_git_setting": {
        "block": {
          "attributes": {
            "label": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "search_paths": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "uri": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "http_basic_auth": {
              "block": {
                "attributes": {
                  "password": {
                    "description_kind": "plain",
                    "required": true,
                    "sensitive": true,
                    "type": "string"
                  },
                  "username": {
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
            "repository": {
              "block": {
                "attributes": {
                  "label": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "pattern": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "search_paths": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "uri": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "http_basic_auth": {
                    "block": {
                      "attributes": {
                        "password": {
                          "description_kind": "plain",
                          "required": true,
                          "sensitive": true,
                          "type": "string"
                        },
                        "username": {
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
                  "ssh_auth": {
                    "block": {
                      "attributes": {
                        "host_key": {
                          "description_kind": "plain",
                          "optional": true,
                          "sensitive": true,
                          "type": "string"
                        },
                        "host_key_algorithm": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "private_key": {
                          "description_kind": "plain",
                          "required": true,
                          "sensitive": true,
                          "type": "string"
                        },
                        "strict_host_key_checking_enabled": {
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
              "nesting_mode": "list"
            },
            "ssh_auth": {
              "block": {
                "attributes": {
                  "host_key": {
                    "description_kind": "plain",
                    "optional": true,
                    "sensitive": true,
                    "type": "string"
                  },
                  "host_key_algorithm": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "private_key": {
                    "description_kind": "plain",
                    "required": true,
                    "sensitive": true,
                    "type": "string"
                  },
                  "strict_host_key_checking_enabled": {
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
        "max_items": 1,
        "nesting_mode": "list"
      },
      "network": {
        "block": {
          "attributes": {
            "app_network_resource_group": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "app_subnet_id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "cidr_ranges": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "list",
                "string"
              ]
            },
            "read_timeout_seconds": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "service_runtime_network_resource_group": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "service_runtime_subnet_id": {
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
      "trace": {
        "block": {
          "attributes": {
            "connection_string": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "sample_rate": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
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
  "version": 1
}`

func AzurermSpringCloudServiceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermSpringCloudService), &result)
	return &result
}
