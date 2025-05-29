package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermServiceFabricManagedCluster = `{
  "block": {
    "attributes": {
      "backup_service_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "client_connection_port": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "dns_name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "dns_service_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "http_gateway_port": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
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
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "password": {
        "description_kind": "plain",
        "optional": true,
        "sensitive": true,
        "type": "string"
      },
      "resource_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "sku": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "subnet_id": {
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
      "upgrade_wave": {
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
    "block_types": {
      "authentication": {
        "block": {
          "block_types": {
            "active_directory": {
              "block": {
                "attributes": {
                  "client_application_id": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "cluster_application_id": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "tenant_id": {
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
                  "common_name": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "thumbprint": {
                    "description_kind": "plain",
                    "required": true,
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
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "custom_fabric_setting": {
        "block": {
          "attributes": {
            "parameter": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "section": {
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
        "nesting_mode": "list"
      },
      "lb_rule": {
        "block": {
          "attributes": {
            "backend_port": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "frontend_port": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "probe_protocol": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "probe_request_path": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "protocol": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "min_items": 1,
        "nesting_mode": "list"
      },
      "node_type": {
        "block": {
          "attributes": {
            "application_port_range": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "capacities": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "data_disk_size_gb": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "data_disk_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "ephemeral_port_range": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "multiple_placement_groups_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "placement_properties": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "primary": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "stateless": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "vm_image_offer": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "vm_image_publisher": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "vm_image_sku": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "vm_image_version": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "vm_instance_count": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "vm_size": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "vm_secrets": {
              "block": {
                "attributes": {
                  "vault_id": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "certificates": {
                    "block": {
                      "attributes": {
                        "store": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "url": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
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
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AzurermServiceFabricManagedClusterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermServiceFabricManagedCluster), &result)
	return &result
}
