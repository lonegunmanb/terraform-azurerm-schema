package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermSystemCenterVirtualMachineManagerVirtualMachineInstance = `{
  "block": {
    "attributes": {
      "custom_location_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "scoped_resource_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "system_center_virtual_machine_manager_availability_set_ids": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      }
    },
    "block_types": {
      "hardware": {
        "block": {
          "attributes": {
            "cpu_count": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "dynamic_memory_max_in_mb": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "dynamic_memory_min_in_mb": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "limit_cpu_for_migration_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "memory_in_mb": {
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
      "infrastructure": {
        "block": {
          "attributes": {
            "checkpoint_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "system_center_virtual_machine_manager_cloud_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "system_center_virtual_machine_manager_inventory_item_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "system_center_virtual_machine_manager_template_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "system_center_virtual_machine_manager_virtual_machine_server_id": {
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
      "network_interface": {
        "block": {
          "attributes": {
            "ipv4_address_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "ipv6_address_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "mac_address_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "virtual_network_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "operating_system": {
        "block": {
          "attributes": {
            "admin_password": {
              "description_kind": "plain",
              "optional": true,
              "sensitive": true,
              "type": "string"
            },
            "computer_name": {
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
      "storage_disk": {
        "block": {
          "attributes": {
            "bus": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "bus_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "disk_size_gb": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "lun": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "storage_qos_policy_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "template_disk_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "vhd_type": {
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
    "description_kind": "plain"
  },
  "version": 0
}`

func AzurermSystemCenterVirtualMachineManagerVirtualMachineInstanceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermSystemCenterVirtualMachineManagerVirtualMachineInstance), &result)
	return &result
}
