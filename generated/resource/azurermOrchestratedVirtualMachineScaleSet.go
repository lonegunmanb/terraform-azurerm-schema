package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermOrchestratedVirtualMachineScaleSet = `{
  "block": {
    "attributes": {
      "capacity_reservation_group_id": {
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
      "extension_operations_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
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
      "instances": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
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
      "platform_fault_domain_count": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
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
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "sku_name": {
        "description_kind": "plain",
        "optional": true,
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
      "unique_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "upgrade_mode": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "user_data_base64": {
        "description_kind": "plain",
        "optional": true,
        "sensitive": true,
        "type": "string"
      },
      "zone_balance": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "zones": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
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
              "optional": true,
              "type": "bool"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "automatic_instance_repair": {
        "block": {
          "attributes": {
            "action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "enabled": {
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            },
            "grace_period": {
              "computed": true,
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
      "data_disk": {
        "block": {
          "attributes": {
            "caching": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "create_option": {
              "description_kind": "plain",
              "optional": true,
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
            "lun": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "storage_account_type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "ultra_ssd_disk_iops_read_write": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "ultra_ssd_disk_mbps_read_write": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
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
      "extension": {
        "block": {
          "attributes": {
            "auto_upgrade_minor_version_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "extensions_to_provision_after_vm_creation": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "failure_suppression_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "force_extension_execution_on_change": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
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
          "block_types": {
            "protected_settings_from_key_vault": {
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
              "required": true,
              "type": [
                "set",
                "string"
              ]
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
      "network_interface": {
        "block": {
          "attributes": {
            "dns_servers": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "enable_accelerated_networking": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "enable_ip_forwarding": {
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
              "optional": true,
              "type": "bool"
            }
          },
          "block_types": {
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
                  "name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "primary": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "subnet_id": {
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
                "block_types": {
                  "public_ip_address": {
                    "block": {
                      "attributes": {
                        "domain_name_label": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "idle_timeout_in_minutes": {
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
                        "public_ip_prefix_id": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "sku_name": {
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
                      "block_types": {
                        "ip_tag": {
                          "block": {
                            "attributes": {
                              "tag": {
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
        "nesting_mode": "list"
      },
      "os_profile": {
        "block": {
          "attributes": {
            "custom_data": {
              "description_kind": "plain",
              "optional": true,
              "sensitive": true,
              "type": "string"
            }
          },
          "block_types": {
            "linux_configuration": {
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
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "disable_password_authentication": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
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
                  "provision_vm_agent": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "block_types": {
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
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "windows_configuration": {
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
                  "computer_name_prefix": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "enable_automatic_updates": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "hotpatching_enabled": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
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
              "max_items": 1,
              "nesting_mode": "list"
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
      "priority_mix": {
        "block": {
          "attributes": {
            "base_regular_count": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "regular_percentage_above_base": {
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
      "rolling_upgrade_policy": {
        "block": {
          "attributes": {
            "cross_zone_upgrades_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "max_batch_instance_percent": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "max_unhealthy_instance_percent": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "max_unhealthy_upgraded_instance_percent": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "maximum_surge_instances_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "pause_time_between_batches": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "prioritize_unhealthy_instances_enabled": {
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
      "sku_profile": {
        "block": {
          "attributes": {
            "allocation_strategy": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "vm_sizes": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "set",
                "string"
              ]
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
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

func AzurermOrchestratedVirtualMachineScaleSetSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermOrchestratedVirtualMachineScaleSet), &result)
	return &result
}
