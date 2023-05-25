package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermContainerRegistryTask = `{
  "block": {
    "attributes": {
      "agent_pool_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "container_registry_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "is_system_task": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "log_template": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
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
      "timeout_in_seconds": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      }
    },
    "block_types": {
      "agent_setting": {
        "block": {
          "attributes": {
            "cpu": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "base_image_trigger": {
        "block": {
          "attributes": {
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
            "type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "update_trigger_endpoint": {
              "description_kind": "plain",
              "optional": true,
              "sensitive": true,
              "type": "string"
            },
            "update_trigger_payload_type": {
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
      "docker_step": {
        "block": {
          "attributes": {
            "arguments": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "cache_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "context_access_token": {
              "description_kind": "plain",
              "required": true,
              "sensitive": true,
              "type": "string"
            },
            "context_path": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "dockerfile_path": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "image_names": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "push_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "secret_arguments": {
              "description_kind": "plain",
              "optional": true,
              "sensitive": true,
              "type": [
                "map",
                "string"
              ]
            },
            "target": {
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
      "encoded_step": {
        "block": {
          "attributes": {
            "context_access_token": {
              "description_kind": "plain",
              "optional": true,
              "sensitive": true,
              "type": "string"
            },
            "context_path": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "secret_values": {
              "description_kind": "plain",
              "optional": true,
              "sensitive": true,
              "type": [
                "map",
                "string"
              ]
            },
            "task_content": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value_content": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "values": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "file_step": {
        "block": {
          "attributes": {
            "context_access_token": {
              "description_kind": "plain",
              "optional": true,
              "sensitive": true,
              "type": "string"
            },
            "context_path": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "secret_values": {
              "description_kind": "plain",
              "optional": true,
              "sensitive": true,
              "type": [
                "map",
                "string"
              ]
            },
            "task_file_path": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value_file_path": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "values": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
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
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "principal_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "tenant_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
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
      "platform": {
        "block": {
          "attributes": {
            "architecture": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "os": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "variant": {
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
      "registry_credential": {
        "block": {
          "block_types": {
            "custom": {
              "block": {
                "attributes": {
                  "identity": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "login_server": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "password": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "username": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "set"
            },
            "source": {
              "block": {
                "attributes": {
                  "login_mode": {
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
        "max_items": 1,
        "nesting_mode": "list"
      },
      "source_trigger": {
        "block": {
          "attributes": {
            "branch": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "events": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "list",
                "string"
              ]
            },
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "repository_url": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "source_type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "authentication": {
              "block": {
                "attributes": {
                  "expire_in_seconds": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "refresh_token": {
                    "description_kind": "plain",
                    "optional": true,
                    "sensitive": true,
                    "type": "string"
                  },
                  "scope": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "token": {
                    "description_kind": "plain",
                    "required": true,
                    "sensitive": true,
                    "type": "string"
                  },
                  "token_type": {
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
      "timer_trigger": {
        "block": {
          "attributes": {
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
            "schedule": {
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
  "version": 0
}`

func AzurermContainerRegistryTaskSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermContainerRegistryTask), &result)
	return &result
}
