package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermHdinsightKafkaCluster = `{
  "block": {
    "attributes": {
      "cluster_version": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "encryption_in_transit_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
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
      "kafka_rest_proxy_endpoint": {
        "computed": true,
        "description_kind": "plain",
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
      "component_version": {
        "block": {
          "attributes": {
            "kafka": {
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
      "compute_isolation": {
        "block": {
          "attributes": {
            "compute_isolation_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "host_sku": {
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
      "disk_encryption": {
        "block": {
          "attributes": {
            "encryption_algorithm": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "encryption_at_host_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "key_vault_key_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "key_vault_managed_identity_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "extension": {
        "block": {
          "attributes": {
            "log_analytics_workspace_id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "primary_key": {
              "description_kind": "plain",
              "required": true,
              "sensitive": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "gateway": {
        "block": {
          "attributes": {
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
      "metastores": {
        "block": {
          "block_types": {
            "ambari": {
              "block": {
                "attributes": {
                  "database_name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "password": {
                    "description_kind": "plain",
                    "required": true,
                    "sensitive": true,
                    "type": "string"
                  },
                  "server": {
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
              "max_items": 1,
              "nesting_mode": "list"
            },
            "hive": {
              "block": {
                "attributes": {
                  "database_name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "password": {
                    "description_kind": "plain",
                    "required": true,
                    "sensitive": true,
                    "type": "string"
                  },
                  "server": {
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
              "max_items": 1,
              "nesting_mode": "list"
            },
            "oozie": {
              "block": {
                "attributes": {
                  "database_name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "password": {
                    "description_kind": "plain",
                    "required": true,
                    "sensitive": true,
                    "type": "string"
                  },
                  "server": {
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
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "monitor": {
        "block": {
          "attributes": {
            "log_analytics_workspace_id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "primary_key": {
              "description_kind": "plain",
              "required": true,
              "sensitive": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "network": {
        "block": {
          "attributes": {
            "connection_direction": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "private_link_enabled": {
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
      "private_link_configuration": {
        "block": {
          "attributes": {
            "group_id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "ip_configuration": {
              "block": {
                "attributes": {
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
                  "private_ip_address": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "private_ip_allocation_method": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "subnet_id": {
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
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "rest_proxy": {
        "block": {
          "attributes": {
            "security_group_id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "security_group_name": {
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
      "roles": {
        "block": {
          "block_types": {
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
                "block_types": {
                  "script_actions": {
                    "block": {
                      "attributes": {
                        "name": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "parameters": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "uri": {
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
              "min_items": 1,
              "nesting_mode": "list"
            },
            "kafka_management_node": {
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
                "block_types": {
                  "script_actions": {
                    "block": {
                      "attributes": {
                        "name": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "parameters": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "uri": {
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
              "nesting_mode": "list"
            },
            "worker_node": {
              "block": {
                "attributes": {
                  "number_of_disks_per_node": {
                    "description_kind": "plain",
                    "required": true,
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
                "block_types": {
                  "script_actions": {
                    "block": {
                      "attributes": {
                        "name": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "parameters": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "uri": {
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
                "block_types": {
                  "script_actions": {
                    "block": {
                      "attributes": {
                        "name": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "parameters": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "uri": {
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
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "deprecated": true,
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "security_profile": {
        "block": {
          "attributes": {
            "aadds_resource_id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "cluster_users_group_dns": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "domain_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "domain_user_password": {
              "description_kind": "plain",
              "required": true,
              "sensitive": true,
              "type": "string"
            },
            "domain_username": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "ldaps_urls": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "set",
                "string"
              ]
            },
            "msi_resource_id": {
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
      "storage_account_gen2": {
        "block": {
          "attributes": {
            "filesystem_id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "is_default": {
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            },
            "managed_identity_resource_id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "storage_resource_id": {
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

func AzurermHdinsightKafkaClusterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermHdinsightKafkaCluster), &result)
	return &result
}
