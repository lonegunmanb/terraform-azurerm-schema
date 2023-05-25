package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermHdinsightMlServicesCluster = `{
  "block": {
    "attributes": {
      "cluster_version": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "edge_ssh_endpoint": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "https_endpoint": {
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
      "rstudio": {
        "description_kind": "plain",
        "required": true,
        "type": "bool"
      },
      "ssh_endpoint": {
        "computed": true,
        "description_kind": "plain",
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
      "tier": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tls_min_version": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "gateway": {
        "block": {
          "attributes": {
            "enabled": {
              "deprecated": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
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
      "roles": {
        "block": {
          "block_types": {
            "edge_node": {
              "block": {
                "attributes": {
                  "password": {
                    "description_kind": "plain",
                    "optional": true,
                    "sensitive": true,
                    "type": "string"
                  },
                  "ssh_keys": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "subnet_id": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "username": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "virtual_network_id": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "vm_size": {
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
            "head_node": {
              "block": {
                "attributes": {
                  "password": {
                    "description_kind": "plain",
                    "optional": true,
                    "sensitive": true,
                    "type": "string"
                  },
                  "ssh_keys": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "subnet_id": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "username": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "virtual_network_id": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "vm_size": {
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
            "worker_node": {
              "block": {
                "attributes": {
                  "min_instance_count": {
                    "computed": true,
                    "deprecated": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "password": {
                    "description_kind": "plain",
                    "optional": true,
                    "sensitive": true,
                    "type": "string"
                  },
                  "ssh_keys": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "subnet_id": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "target_instance_count": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "username": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "virtual_network_id": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "vm_size": {
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
            "zookeeper_node": {
              "block": {
                "attributes": {
                  "password": {
                    "description_kind": "plain",
                    "optional": true,
                    "sensitive": true,
                    "type": "string"
                  },
                  "ssh_keys": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "subnet_id": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "username": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "virtual_network_id": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "vm_size": {
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
      },
      "storage_account": {
        "block": {
          "attributes": {
            "is_default": {
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            },
            "storage_account_key": {
              "description_kind": "plain",
              "required": true,
              "sensitive": true,
              "type": "string"
            },
            "storage_container_id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "storage_resource_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
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
    "deprecated": true,
    "description_kind": "plain"
  },
  "version": 0
}`

func AzurermHdinsightMlServicesClusterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermHdinsightMlServicesCluster), &result)
	return &result
}
