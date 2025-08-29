package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermLinuxVirtualMachine = `{
  "block": {
    "attributes": {
      "admin_password": {
        "description_kind": "plain",
        "optional": true,
        "sensitive": true,
        "type": "string"
      },
      "admin_username": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "allow_extension_operations": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "availability_set_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "bypass_platform_safety_checks_on_user_schedule_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "capacity_reservation_group_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "computer_name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "custom_data": {
        "description_kind": "plain",
        "optional": true,
        "sensitive": true,
        "type": "string"
      },
      "dedicated_host_group_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "dedicated_host_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "disable_password_authentication": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "disk_controller_type": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "edge_zone": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "encryption_at_host_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "eviction_policy": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "extensions_time_budget": {
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
      "license_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "location": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "max_bid_price": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "network_interface_ids": {
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
      },
      "os_managed_disk_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "patch_assessment_mode": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "patch_mode": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "platform_fault_domain": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "priority": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "private_ip_address": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "private_ip_addresses": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "provision_vm_agent": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "proximity_placement_group_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "public_ip_address": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "public_ip_addresses": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "reboot_setting": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "resource_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "secure_boot_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "size": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "source_image_id": {
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
      "user_data": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "virtual_machine_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "virtual_machine_scale_set_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "vm_agent_platform_updates_enabled": {
        "computed": true,
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "vtpm_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "zone": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "additional_capabilities": {
        "block": {
          "attributes": {
            "hibernation_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "ultra_ssd_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "admin_ssh_key": {
        "block": {
          "attributes": {
            "public_key": {
              "description_kind": "plain",
              "required": true,
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
        "nesting_mode": "set"
      },
      "boot_diagnostics": {
        "block": {
          "attributes": {
            "storage_account_uri": {
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
      "gallery_application": {
        "block": {
          "attributes": {
            "automatic_upgrade_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "configuration_blob_uri": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "order": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "tag": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "treat_failure_as_deployment_failure_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "version_id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 100,
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
      "os_disk": {
        "block": {
          "attributes": {
            "caching": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "disk_encryption_set_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "disk_size_gb": {
              "computed": true,
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
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "secure_vm_disk_encryption_set_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "security_encryption_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "storage_account_type": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "write_accelerator_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "block_types": {
            "diff_disk_settings": {
              "block": {
                "attributes": {
                  "option": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "placement": {
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
      },
      "os_image_notification": {
        "block": {
          "attributes": {
            "timeout": {
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
      "plan": {
        "block": {
          "attributes": {
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "product": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "publisher": {
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
      "secret": {
        "block": {
          "attributes": {
            "key_vault_id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "certificate": {
              "block": {
                "attributes": {
                  "url": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "min_items": 1,
              "nesting_mode": "set"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "source_image_reference": {
        "block": {
          "attributes": {
            "offer": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "publisher": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "sku": {
              "description_kind": "plain",
              "required": true,
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
        "nesting_mode": "list"
      },
      "termination_notification": {
        "block": {
          "attributes": {
            "enabled": {
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            },
            "timeout": {
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

func AzurermLinuxVirtualMachineSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermLinuxVirtualMachine), &result)
	return &result
}
