package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermVirtualMachineScaleSet = `{
  "block": {
    "attributes": {
      "automatic_os_upgrade": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "eviction_policy": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "health_probe_id": {
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
      "overprovision": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "priority": {
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
      "single_placement_group": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "tags": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "upgrade_policy_mode": {
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
      "boot_diagnostics": {
        "block": {
          "attributes": {
            "enabled": {
              "description_kind": "plain",
              "optional": true,
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
      "extension": {
        "block": {
          "attributes": {
            "auto_upgrade_minor_version": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "protected_settings": {
              "description_kind": "plain",
              "optional": true,
              "sensitive": true,
              "type": "string"
            },
            "provision_after_extensions": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "publisher": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "settings": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "type_handler_version": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
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
      "network_profile": {
        "block": {
          "attributes": {
            "accelerated_networking": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "ip_forwarding": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "network_security_group_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "primary": {
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            }
          },
          "block_types": {
            "dns_settings": {
              "block": {
                "attributes": {
                  "dns_servers": {
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "ip_configuration": {
              "block": {
                "attributes": {
                  "application_gateway_backend_address_pool_ids": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "application_security_group_ids": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "load_balancer_backend_address_pool_ids": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "load_balancer_inbound_nat_rules_ids": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "primary": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "bool"
                  },
                  "subnet_id": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "public_ip_address_configuration": {
                    "block": {
                      "attributes": {
                        "domain_name_label": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "idle_timeout": {
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
                    "nesting_mode": "list"
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
        "min_items": 1,
        "nesting_mode": "set"
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
            "computer_name_prefix": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "custom_data": {
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
      "os_profile_linux_config": {
        "block": {
          "attributes": {
            "disable_password_authentication": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "block_types": {
            "ssh_keys": {
              "block": {
                "attributes": {
                  "key_data": {
                    "description_kind": "plain",
                    "optional": true,
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
        "nesting_mode": "set"
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
        "nesting_mode": "set"
      },
      "rolling_upgrade_policy": {
        "block": {
          "attributes": {
            "max_batch_instance_percent": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "max_unhealthy_instance_percent": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "max_unhealthy_upgraded_instance_percent": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "pause_time_between_batches": {
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
            },
            "tier": {
              "computed": true,
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
      "storage_profile_data_disk": {
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
            "managed_disk_type": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "storage_profile_image_reference": {
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
        "nesting_mode": "set"
      },
      "storage_profile_os_disk": {
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
            "image": {
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
              "optional": true,
              "type": "string"
            },
            "os_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "vhd_containers": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "set"
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
  "version": 1
}`

func AzurermVirtualMachineScaleSetSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermVirtualMachineScaleSet), &result)
	return &result
}
