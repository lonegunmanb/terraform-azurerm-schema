package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermVirtualMachine = `{
  "block": {
    "attributes": {
      "availability_set_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "delete_data_disks_on_termination": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "delete_os_disk_on_termination": {
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
      "network_interface_ids": {
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
      },
      "primary_network_interface_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "proximity_placement_group_id": {
        "description_kind": "plain",
        "optional": true,
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
      "vm_size": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "zones": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      }
    },
    "block_types": {
      "additional_capabilities": {
        "block": {
          "attributes": {
            "ultra_ssd_enabled": {
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "boot_diagnostics": {
        "block": {
          "attributes": {
            "enabled": {
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            },
            "storage_uri": {
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
      "identity": {
        "block": {
          "attributes": {
            "identity_ids": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "principal_id": {
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
      "os_profile": {
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
              "required": true,
              "type": "string"
            },
            "computer_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "custom_data": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "set"
      },
      "os_profile_linux_config": {
        "block": {
          "attributes": {
            "disable_password_authentication": {
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            }
          },
          "block_types": {
            "ssh_keys": {
              "block": {
                "attributes": {
                  "key_data": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "path": {
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
        "nesting_mode": "set"
      },
      "os_profile_secrets": {
        "block": {
          "attributes": {
            "source_vault_id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "vault_certificates": {
              "block": {
                "attributes": {
                  "certificate_store": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "certificate_url": {
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
      },
      "os_profile_windows_config": {
        "block": {
          "attributes": {
            "enable_automatic_upgrades": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "provision_vm_agent": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "timezone": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "additional_unattend_config": {
              "block": {
                "attributes": {
                  "component": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "content": {
                    "description_kind": "plain",
                    "required": true,
                    "sensitive": true,
                    "type": "string"
                  },
                  "pass": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "setting_name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "winrm": {
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
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "set"
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
      "storage_data_disk": {
        "block": {
          "attributes": {
            "caching": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "create_option": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "disk_size_gb": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "lun": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "managed_disk_id": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "managed_disk_type": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "vhd_uri": {
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
          "description_kind": "plain"
        },
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
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "set"
      },
      "storage_os_disk": {
        "block": {
          "attributes": {
            "caching": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "create_option": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "disk_size_gb": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "image_uri": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "managed_disk_id": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "managed_disk_type": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "os_type": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "vhd_uri": {
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

func AzurermVirtualMachineSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermVirtualMachine), &result)
	return &result
}
