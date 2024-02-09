package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermStorageAccount = `{
  "block": {
    "attributes": {
      "access_tier": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "account_kind": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "account_replication_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "account_tier": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "allow_nested_items_to_be_public": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "allowed_copy_scope": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "cross_tenant_replication_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "default_to_oauth_authentication": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "edge_zone": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "enable_https_traffic_only": {
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
      "infrastructure_encryption_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "is_hns_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "large_file_share_enabled": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "local_user_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "location": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "min_tls_version": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "nfsv3_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "primary_access_key": {
        "computed": true,
        "description_kind": "plain",
        "sensitive": true,
        "type": "string"
      },
      "primary_blob_connection_string": {
        "computed": true,
        "description_kind": "plain",
        "sensitive": true,
        "type": "string"
      },
      "primary_blob_endpoint": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "primary_blob_host": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "primary_blob_internet_endpoint": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "primary_blob_internet_host": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "primary_blob_microsoft_endpoint": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "primary_blob_microsoft_host": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "primary_connection_string": {
        "computed": true,
        "description_kind": "plain",
        "sensitive": true,
        "type": "string"
      },
      "primary_dfs_endpoint": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "primary_dfs_host": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "primary_dfs_internet_endpoint": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "primary_dfs_internet_host": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "primary_dfs_microsoft_endpoint": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "primary_dfs_microsoft_host": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "primary_file_endpoint": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "primary_file_host": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "primary_file_internet_endpoint": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "primary_file_internet_host": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "primary_file_microsoft_endpoint": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "primary_file_microsoft_host": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "primary_location": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "primary_queue_endpoint": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "primary_queue_host": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "primary_queue_microsoft_endpoint": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "primary_queue_microsoft_host": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "primary_table_endpoint": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "primary_table_host": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "primary_table_microsoft_endpoint": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "primary_table_microsoft_host": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "primary_web_endpoint": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "primary_web_host": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "primary_web_internet_endpoint": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "primary_web_internet_host": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "primary_web_microsoft_endpoint": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "primary_web_microsoft_host": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "public_network_access_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "queue_encryption_key_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "resource_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "secondary_access_key": {
        "computed": true,
        "description_kind": "plain",
        "sensitive": true,
        "type": "string"
      },
      "secondary_blob_connection_string": {
        "computed": true,
        "description_kind": "plain",
        "sensitive": true,
        "type": "string"
      },
      "secondary_blob_endpoint": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "secondary_blob_host": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "secondary_blob_internet_endpoint": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "secondary_blob_internet_host": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "secondary_blob_microsoft_endpoint": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "secondary_blob_microsoft_host": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "secondary_connection_string": {
        "computed": true,
        "description_kind": "plain",
        "sensitive": true,
        "type": "string"
      },
      "secondary_dfs_endpoint": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "secondary_dfs_host": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "secondary_dfs_internet_endpoint": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "secondary_dfs_internet_host": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "secondary_dfs_microsoft_endpoint": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "secondary_dfs_microsoft_host": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "secondary_file_endpoint": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "secondary_file_host": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "secondary_file_internet_endpoint": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "secondary_file_internet_host": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "secondary_file_microsoft_endpoint": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "secondary_file_microsoft_host": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "secondary_location": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "secondary_queue_endpoint": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "secondary_queue_host": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "secondary_queue_microsoft_endpoint": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "secondary_queue_microsoft_host": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "secondary_table_endpoint": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "secondary_table_host": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "secondary_table_microsoft_endpoint": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "secondary_table_microsoft_host": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "secondary_web_endpoint": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "secondary_web_host": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "secondary_web_internet_endpoint": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "secondary_web_internet_host": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "secondary_web_microsoft_endpoint": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "secondary_web_microsoft_host": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "sftp_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "shared_access_key_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "table_encryption_key_type": {
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
      }
    },
    "block_types": {
      "azure_files_authentication": {
        "block": {
          "attributes": {
            "directory_type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "active_directory": {
              "block": {
                "attributes": {
                  "domain_guid": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "domain_name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "domain_sid": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "forest_name": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "netbios_domain_name": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "storage_sid": {
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
      "blob_properties": {
        "block": {
          "attributes": {
            "change_feed_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "change_feed_retention_in_days": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "default_service_version": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "last_access_time_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "versioning_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "block_types": {
            "container_delete_retention_policy": {
              "block": {
                "attributes": {
                  "days": {
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
            "cors_rule": {
              "block": {
                "attributes": {
                  "allowed_headers": {
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "allowed_methods": {
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "allowed_origins": {
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "exposed_headers": {
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "max_age_in_seconds": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 5,
              "nesting_mode": "list"
            },
            "delete_retention_policy": {
              "block": {
                "attributes": {
                  "days": {
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
            "restore_policy": {
              "block": {
                "attributes": {
                  "days": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
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
      "custom_domain": {
        "block": {
          "attributes": {
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "use_subdomain": {
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
      "customer_managed_key": {
        "block": {
          "attributes": {
            "key_vault_key_id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "user_assigned_identity_id": {
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
      "immutability_policy": {
        "block": {
          "attributes": {
            "allow_protected_append_writes": {
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            },
            "period_since_creation_in_days": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "state": {
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
      "network_rules": {
        "block": {
          "attributes": {
            "bypass": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "default_action": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "ip_rules": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "virtual_network_subnet_ids": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            }
          },
          "block_types": {
            "private_link_access": {
              "block": {
                "attributes": {
                  "endpoint_resource_id": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "endpoint_tenant_id": {
                    "computed": true,
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
        "max_items": 1,
        "nesting_mode": "list"
      },
      "queue_properties": {
        "block": {
          "block_types": {
            "cors_rule": {
              "block": {
                "attributes": {
                  "allowed_headers": {
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "allowed_methods": {
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "allowed_origins": {
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "exposed_headers": {
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "max_age_in_seconds": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 5,
              "nesting_mode": "list"
            },
            "hour_metrics": {
              "block": {
                "attributes": {
                  "enabled": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "bool"
                  },
                  "include_apis": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "retention_policy_days": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
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
            "logging": {
              "block": {
                "attributes": {
                  "delete": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "bool"
                  },
                  "read": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "bool"
                  },
                  "retention_policy_days": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "version": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "write": {
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
            "minute_metrics": {
              "block": {
                "attributes": {
                  "enabled": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "bool"
                  },
                  "include_apis": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "retention_policy_days": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
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
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "routing": {
        "block": {
          "attributes": {
            "choice": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "publish_internet_endpoints": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "publish_microsoft_endpoints": {
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
      "sas_policy": {
        "block": {
          "attributes": {
            "expiration_action": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "expiration_period": {
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
      "share_properties": {
        "block": {
          "block_types": {
            "cors_rule": {
              "block": {
                "attributes": {
                  "allowed_headers": {
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "allowed_methods": {
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "allowed_origins": {
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "exposed_headers": {
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "max_age_in_seconds": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 5,
              "nesting_mode": "list"
            },
            "retention_policy": {
              "block": {
                "attributes": {
                  "days": {
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
            "smb": {
              "block": {
                "attributes": {
                  "authentication_types": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "channel_encryption_type": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "kerberos_ticket_encryption_type": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "multichannel_enabled": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "versions": {
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
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "static_website": {
        "block": {
          "attributes": {
            "error_404_document": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "index_document": {
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
  "version": 4
}`

func AzurermStorageAccountSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermStorageAccount), &result)
	return &result
}
