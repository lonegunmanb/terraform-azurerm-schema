package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermRedhatOpenshiftCluster = `{
  "block": {
    "attributes": {
      "console_url": {
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
      "api_server_profile": {
        "block": {
          "attributes": {
            "ip_address": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "url": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "visibility": {
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
      "cluster_profile": {
        "block": {
          "attributes": {
            "domain": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "fips_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "pull_secret": {
              "description_kind": "plain",
              "optional": true,
              "sensitive": true,
              "type": "string"
            },
            "resource_group_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "version": {
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
      "ingress_profile": {
        "block": {
          "attributes": {
            "ip_address": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "visibility": {
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
      "main_profile": {
        "block": {
          "attributes": {
            "disk_encryption_set_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "encryption_at_host_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "subnet_id": {
              "description_kind": "plain",
              "required": true,
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
      "network_profile": {
        "block": {
          "attributes": {
            "outbound_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "pod_cidr": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "service_cidr": {
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
      "service_principal": {
        "block": {
          "attributes": {
            "client_id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "client_secret": {
              "description_kind": "plain",
              "required": true,
              "sensitive": true,
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
      "worker_profile": {
        "block": {
          "attributes": {
            "disk_encryption_set_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "disk_size_gb": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "encryption_at_host_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "node_count": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "subnet_id": {
              "description_kind": "plain",
              "required": true,
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
  "version": 0
}`

func AzurermRedhatOpenshiftClusterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermRedhatOpenshiftCluster), &result)
	return &result
}
