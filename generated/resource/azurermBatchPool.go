package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermBatchPool = `{
  "block": {
    "attributes": {
      "account_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "display_name": {
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
      "max_tasks_per_node": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "metadata": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "node_agent_sku_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "resource_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "stop_pending_resize_operation": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "vm_size": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "auto_scale": {
        "block": {
          "attributes": {
            "evaluation_interval": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "formula": {
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
      "certificate": {
        "block": {
          "attributes": {
            "id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "store_location": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "store_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "visibility": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "container_configuration": {
        "block": {
          "attributes": {
            "container_image_names": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "container_registries": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                [
                  "object",
                  {
                    "password": "string",
                    "registry_server": "string",
                    "user_name": "string"
                  }
                ]
              ]
            },
            "type": {
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
      "fixed_scale": {
        "block": {
          "attributes": {
            "resize_timeout": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "target_dedicated_nodes": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "target_low_priority_nodes": {
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
      "identity": {
        "block": {
          "attributes": {
            "identity_ids": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "set",
                "string"
              ]
            },
            "type": {
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
      "network_configuration": {
        "block": {
          "attributes": {
            "public_address_provisioning_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "public_ips": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "subnet_id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "endpoint_configuration": {
              "block": {
                "attributes": {
                  "backend_port": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "frontend_port_range": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "protocol": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "network_security_group_rules": {
                    "block": {
                      "attributes": {
                        "access": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "priority": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        },
                        "source_address_prefix": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "start_task": {
        "block": {
          "attributes": {
            "command_line": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "common_environment_properties": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "environment": {
              "computed": true,
              "deprecated": true,
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "max_task_retry_count": {
              "computed": true,
              "deprecated": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "task_retry_maximum": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "wait_for_success": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "block_types": {
            "resource_file": {
              "block": {
                "attributes": {
                  "auto_storage_container_name": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "blob_prefix": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "file_mode": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "file_path": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "http_url": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "storage_container_url": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "user_identity": {
              "block": {
                "attributes": {
                  "user_name": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "auto_user": {
                    "block": {
                      "attributes": {
                        "elevation_level": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "scope": {
                          "description_kind": "plain",
                          "optional": true,
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
      "storage_image_reference": {
        "block": {
          "attributes": {
            "id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "offer": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "publisher": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "sku": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "version": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
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
  "version": 0
}`

func AzurermBatchPoolSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermBatchPool), &result)
	return &result
}
