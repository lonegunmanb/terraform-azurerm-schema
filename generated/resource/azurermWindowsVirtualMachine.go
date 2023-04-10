package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermWindowsVirtualMachine = `{
  "block": {
    "attributes": {
      "admin_password": {
        "description_kind": "plain",
        "required": true,
        "sensitive": true,
        "type": "string"
      },
      "admin_username": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "allow_extension_operations": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "availability_set_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
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
      "edge_zone": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "enable_automatic_updates": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
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
      "hotpatching_enabled": {
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
      "patch_assessment_mode": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "patch_mode": {
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
      "timezone": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
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
      "additional_unattend_content": {
        "block": {
          "attributes": {
            "content": {
              "description_kind": "plain",
              "required": true,
              "sensitive": true,
              "type": "string"
            },
            "setting": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
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
              "description_kind": "plain",
              "required": true,
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
      },
      "winrm_listener": {
        "block": {
          "attributes": {
            "certificate_url": {
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
        "nesting_mode": "set"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AzurermWindowsVirtualMachineSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermWindowsVirtualMachine), &result)
	return &result
}
