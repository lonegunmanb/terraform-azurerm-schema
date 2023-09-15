package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermManagedDisk = `{
  "block": {
    "attributes": {
      "create_option": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "disk_access_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "disk_encryption_set_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "disk_iops_read_only": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "disk_iops_read_write": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "disk_mbps_read_only": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "disk_mbps_read_write": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "disk_size_gb": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "edge_zone": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "gallery_image_reference_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "hyper_v_generation": {
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
      "image_reference_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "location": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "logical_sector_size": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "max_shares": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "network_access_policy": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "on_demand_bursting_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "optimized_frequent_attach_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "os_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "performance_plus_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "public_network_access_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "resource_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "secure_vm_disk_encryption_set_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "security_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "source_resource_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "source_uri": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "storage_account_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "storage_account_type": {
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
      "tier": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "trusted_launch_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "upload_size_bytes": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "zone": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "encryption_settings": {
        "block": {
          "attributes": {
            "enabled": {
              "deprecated": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "block_types": {
            "disk_encryption_key": {
              "block": {
                "attributes": {
                  "secret_url": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "source_vault_id": {
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
            "key_encryption_key": {
              "block": {
                "attributes": {
                  "key_url": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "source_vault_id": {
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
  "version": 1
}`

func AzurermManagedDiskSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermManagedDisk), &result)
	return &result
}
