package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermLabServiceLab = `{
  "block": {
    "attributes": {
      "description": {
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
      "lab_plan_id": {
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
      "resource_group_name": {
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
      "title": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "auto_shutdown": {
        "block": {
          "attributes": {
            "disconnect_delay": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "idle_delay": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "no_connect_delay": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "shutdown_on_idle": {
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
      "connection_setting": {
        "block": {
          "attributes": {
            "client_rdp_access": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "client_ssh_access": {
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
      "network": {
        "block": {
          "attributes": {
            "load_balancer_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "public_ip_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "subnet_id": {
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
      "roster": {
        "block": {
          "attributes": {
            "active_directory_group_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "lms_instance": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "lti_client_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "lti_context_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "lti_roster_endpoint": {
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
      "security": {
        "block": {
          "attributes": {
            "open_access_enabled": {
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            },
            "registration_code": {
              "computed": true,
              "description_kind": "plain",
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
      },
      "virtual_machine": {
        "block": {
          "attributes": {
            "additional_capability_gpu_drivers_installed": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "create_option": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "shared_password_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "usage_quota": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "admin_user": {
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
              "min_items": 1,
              "nesting_mode": "list"
            },
            "image_reference": {
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
            "non_admin_user": {
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
            "sku": {
              "block": {
                "attributes": {
                  "capacity": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "name": {
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
        "min_items": 1,
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AzurermLabServiceLabSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermLabServiceLab), &result)
	return &result
}
