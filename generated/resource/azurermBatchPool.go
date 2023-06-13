package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermBatchPool = `{
  "block": {
    "attributes": {
      "account_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "display_name": {
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
      "inter_node_communication": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "license_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "max_tasks_per_node": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "metadata": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "node_agent_sku_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "os_disk_placement": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "resource_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "stop_pending_resize_operation": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "target_node_communication_mode": {
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
    "block_types": {
      "auto_scale": {
        "block": {
          "attributes": {
            "evaluation_interval": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "formula": {
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
            "id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "store_location": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "store_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "visibility": {
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
        "nesting_mode": "list"
      },
      "container_configuration": {
        "block": {
          "attributes": {
            "container_image_names": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "container_registries": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                [
                  "object",
                  {
                    "password": "string",
                    "registry_server": "string",
                    "user_assigned_identity_id": "string",
                    "user_name": "string"
                  }
                ]
              ]
            },
            "type": {
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
      "data_disks": {
        "block": {
          "attributes": {
            "caching": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "disk_size_gb": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "lun": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "storage_account_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "disk_encryption": {
        "block": {
          "attributes": {
            "disk_encryption_target": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "extensions": {
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
            "settings_json": {
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
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "fixed_scale": {
        "block": {
          "attributes": {
            "node_deallocation_method": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "resize_timeout": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "target_dedicated_nodes": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "target_low_priority_nodes": {
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
      "mount": {
        "block": {
          "block_types": {
            "azure_blob_file_system": {
              "block": {
                "attributes": {
                  "account_key": {
                    "description_kind": "plain",
                    "optional": true,
                    "sensitive": true,
                    "type": "string"
                  },
                  "account_name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "blobfuse_options": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "container_name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "identity_id": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "relative_mount_path": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "sas_key": {
                    "description_kind": "plain",
                    "optional": true,
                    "sensitive": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "azure_file_share": {
              "block": {
                "attributes": {
                  "account_key": {
                    "description_kind": "plain",
                    "required": true,
                    "sensitive": true,
                    "type": "string"
                  },
                  "account_name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "azure_file_url": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "mount_options": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "relative_mount_path": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "cifs_mount": {
              "block": {
                "attributes": {
                  "mount_options": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "password": {
                    "description_kind": "plain",
                    "required": true,
                    "sensitive": true,
                    "type": "string"
                  },
                  "relative_mount_path": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "source": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "user_name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "nfs_mount": {
              "block": {
                "attributes": {
                  "mount_options": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "relative_mount_path": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "source": {
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
      "network_configuration": {
        "block": {
          "attributes": {
            "dynamic_vnet_assignment_scope": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "public_address_provisioning_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "public_ips": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "subnet_id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "endpoint_configuration": {
              "block": {
                "attributes": {
                  "backend_port": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "frontend_port_range": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "protocol": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "network_security_group_rules": {
                    "block": {
                      "attributes": {
                        "access": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "priority": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        },
                        "source_address_prefix": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "source_port_ranges": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
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
        "max_items": 1,
        "nesting_mode": "list"
      },
      "node_placement": {
        "block": {
          "attributes": {
            "policy": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "start_task": {
        "block": {
          "attributes": {
            "command_line": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "common_environment_properties": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "task_retry_maximum": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "wait_for_success": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "block_types": {
            "container": {
              "block": {
                "attributes": {
                  "image_name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "run_options": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "working_directory": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "registry": {
                    "block": {
                      "attributes": {
                        "password": {
                          "description_kind": "plain",
                          "optional": true,
                          "sensitive": true,
                          "type": "string"
                        },
                        "registry_server": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "user_assigned_identity_id": {
                          "description": "The User Assigned Identity to use for Container Registry access.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "user_name": {
                          "description_kind": "plain",
                          "optional": true,
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
            "resource_file": {
              "block": {
                "attributes": {
                  "auto_storage_container_name": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "blob_prefix": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "file_mode": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "file_path": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "http_url": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "storage_container_url": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "user_assigned_identity_id": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "user_identity": {
              "block": {
                "attributes": {
                  "user_name": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "auto_user": {
                    "block": {
                      "attributes": {
                        "elevation_level": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "scope": {
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
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
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
      "task_scheduling_policy": {
        "block": {
          "attributes": {
            "node_fill_type": {
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
      "user_accounts": {
        "block": {
          "attributes": {
            "elevation_level": {
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
              "required": true,
              "sensitive": true,
              "type": "string"
            }
          },
          "block_types": {
            "linux_user_configuration": {
              "block": {
                "attributes": {
                  "gid": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "ssh_private_key": {
                    "description_kind": "plain",
                    "optional": true,
                    "sensitive": true,
                    "type": "string"
                  },
                  "uid": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "windows_user_configuration": {
              "block": {
                "attributes": {
                  "login_mode": {
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
      "windows": {
        "block": {
          "attributes": {
            "enable_automatic_updates": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AzurermBatchPoolSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermBatchPool), &result)
	return &result
}
