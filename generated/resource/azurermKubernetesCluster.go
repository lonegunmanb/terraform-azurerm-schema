package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermKubernetesCluster = `{
  "block": {
    "attributes": {
      "api_server_authorized_ip_ranges": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "automatic_channel_upgrade": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "azure_policy_enabled": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "disk_encryption_set_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "dns_prefix": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "dns_prefix_private_cluster": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "enable_pod_security_policy": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "fqdn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "http_application_routing_enabled": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "http_application_routing_zone_name": {
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
      "kube_admin_config": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "client_certificate": "string",
              "client_key": "string",
              "cluster_ca_certificate": "string",
              "host": "string",
              "password": "string",
              "username": "string"
            }
          ]
        ]
      },
      "kube_admin_config_raw": {
        "computed": true,
        "description_kind": "plain",
        "sensitive": true,
        "type": "string"
      },
      "kube_config": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "client_certificate": "string",
              "client_key": "string",
              "cluster_ca_certificate": "string",
              "host": "string",
              "password": "string",
              "username": "string"
            }
          ]
        ]
      },
      "kube_config_raw": {
        "computed": true,
        "description_kind": "plain",
        "sensitive": true,
        "type": "string"
      },
      "kubernetes_version": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "local_account_disabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
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
      "node_resource_group": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "open_service_mesh_enabled": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "portal_fqdn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "private_cluster_enabled": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "private_cluster_public_fqdn_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "private_dns_zone_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "private_fqdn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "private_link_enabled": {
        "computed": true,
        "deprecated": true,
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
      "role_based_access_control_enabled": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "sku_tier": {
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
      "aci_connector_linux": {
        "block": {
          "attributes": {
            "subnet_name": {
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
      "addon_profile": {
        "block": {
          "block_types": {
            "aci_connector_linux": {
              "block": {
                "attributes": {
                  "enabled": {
                    "deprecated": true,
                    "description_kind": "plain",
                    "required": true,
                    "type": "bool"
                  },
                  "subnet_name": {
                    "deprecated": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "deprecated": true,
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "azure_keyvault_secrets_provider": {
              "block": {
                "attributes": {
                  "enabled": {
                    "deprecated": true,
                    "description_kind": "plain",
                    "required": true,
                    "type": "bool"
                  },
                  "secret_identity": {
                    "computed": true,
                    "deprecated": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      [
                        "object",
                        {
                          "client_id": "string",
                          "object_id": "string",
                          "user_assigned_identity_id": "string"
                        }
                      ]
                    ]
                  },
                  "secret_rotation_enabled": {
                    "deprecated": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "secret_rotation_interval": {
                    "deprecated": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "deprecated": true,
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "azure_policy": {
              "block": {
                "attributes": {
                  "enabled": {
                    "deprecated": true,
                    "description_kind": "plain",
                    "required": true,
                    "type": "bool"
                  }
                },
                "deprecated": true,
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "http_application_routing": {
              "block": {
                "attributes": {
                  "enabled": {
                    "deprecated": true,
                    "description_kind": "plain",
                    "required": true,
                    "type": "bool"
                  },
                  "http_application_routing_zone_name": {
                    "computed": true,
                    "deprecated": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "deprecated": true,
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "ingress_application_gateway": {
              "block": {
                "attributes": {
                  "effective_gateway_id": {
                    "computed": true,
                    "deprecated": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "enabled": {
                    "deprecated": true,
                    "description_kind": "plain",
                    "required": true,
                    "type": "bool"
                  },
                  "gateway_id": {
                    "deprecated": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "gateway_name": {
                    "deprecated": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "ingress_application_gateway_identity": {
                    "computed": true,
                    "deprecated": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      [
                        "object",
                        {
                          "client_id": "string",
                          "object_id": "string",
                          "user_assigned_identity_id": "string"
                        }
                      ]
                    ]
                  },
                  "subnet_cidr": {
                    "deprecated": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "subnet_id": {
                    "deprecated": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "deprecated": true,
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "kube_dashboard": {
              "block": {
                "attributes": {
                  "enabled": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "bool"
                  }
                },
                "deprecated": true,
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "oms_agent": {
              "block": {
                "attributes": {
                  "enabled": {
                    "deprecated": true,
                    "description_kind": "plain",
                    "required": true,
                    "type": "bool"
                  },
                  "log_analytics_workspace_id": {
                    "deprecated": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "oms_agent_identity": {
                    "computed": true,
                    "deprecated": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      [
                        "object",
                        {
                          "client_id": "string",
                          "object_id": "string",
                          "user_assigned_identity_id": "string"
                        }
                      ]
                    ]
                  }
                },
                "deprecated": true,
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "open_service_mesh": {
              "block": {
                "attributes": {
                  "enabled": {
                    "deprecated": true,
                    "description_kind": "plain",
                    "required": true,
                    "type": "bool"
                  }
                },
                "deprecated": true,
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "deprecated": true,
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "auto_scaler_profile": {
        "block": {
          "attributes": {
            "balance_similar_node_groups": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "empty_bulk_delete_max": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "expander": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "max_graceful_termination_sec": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "max_node_provisioning_time": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "max_unready_nodes": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "max_unready_percentage": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "new_pod_scale_up_delay": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "scale_down_delay_after_add": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "scale_down_delay_after_delete": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "scale_down_delay_after_failure": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "scale_down_unneeded": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "scale_down_unready": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "scale_down_utilization_threshold": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "scan_interval": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "skip_nodes_with_local_storage": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "skip_nodes_with_system_pods": {
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
      "azure_active_directory_role_based_access_control": {
        "block": {
          "attributes": {
            "admin_group_object_ids": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "azure_rbac_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "client_app_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "managed": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "server_app_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "server_app_secret": {
              "description_kind": "plain",
              "optional": true,
              "sensitive": true,
              "type": "string"
            },
            "tenant_id": {
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
      "default_node_pool": {
        "block": {
          "attributes": {
            "availability_zones": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "enable_auto_scaling": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "enable_host_encryption": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "enable_node_public_ip": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "fips_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "kubelet_disk_type": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "max_count": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "max_pods": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "min_count": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "node_count": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "node_labels": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "node_public_ip_prefix_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "node_taints": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "only_critical_addons_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "orchestrator_version": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "os_disk_size_gb": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "os_disk_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "os_sku": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "pod_subnet_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "proximity_placement_group_id": {
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
            "type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "ultra_ssd_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "vm_size": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "vnet_subnet_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "kubelet_config": {
              "block": {
                "attributes": {
                  "allowed_unsafe_sysctls": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "container_log_max_line": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "container_log_max_size_mb": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "cpu_cfs_quota_enabled": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "cpu_cfs_quota_period": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "cpu_manager_policy": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "image_gc_high_threshold": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "image_gc_low_threshold": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "pod_max_pid": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "topology_manager_policy": {
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
            "linux_os_config": {
              "block": {
                "attributes": {
                  "swap_file_size_mb": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "transparent_huge_page_defrag": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "transparent_huge_page_enabled": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "sysctl_config": {
                    "block": {
                      "attributes": {
                        "fs_aio_max_nr": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "fs_file_max": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "fs_inotify_max_user_watches": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "fs_nr_open": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "kernel_threads_max": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "net_core_netdev_max_backlog": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "net_core_optmem_max": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "net_core_rmem_default": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "net_core_rmem_max": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "net_core_somaxconn": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "net_core_wmem_default": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "net_core_wmem_max": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "net_ipv4_ip_local_port_range_max": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "net_ipv4_ip_local_port_range_min": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "net_ipv4_neigh_default_gc_thresh1": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "net_ipv4_neigh_default_gc_thresh2": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "net_ipv4_neigh_default_gc_thresh3": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "net_ipv4_tcp_fin_timeout": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "net_ipv4_tcp_keepalive_intvl": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "net_ipv4_tcp_keepalive_probes": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "net_ipv4_tcp_keepalive_time": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "net_ipv4_tcp_max_syn_backlog": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "net_ipv4_tcp_max_tw_buckets": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "net_ipv4_tcp_tw_reuse": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "net_netfilter_nf_conntrack_buckets": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "net_netfilter_nf_conntrack_max": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "vm_max_map_count": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "vm_swappiness": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "vm_vfs_cache_pressure": {
                          "description_kind": "plain",
                          "optional": true,
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
            "upgrade_settings": {
              "block": {
                "attributes": {
                  "max_surge": {
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
        "min_items": 1,
        "nesting_mode": "list"
      },
      "http_proxy_config": {
        "block": {
          "attributes": {
            "http_proxy": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "https_proxy": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "no_proxy": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "trusted_ca": {
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
      "identity": {
        "block": {
          "attributes": {
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
            },
            "user_assigned_identity_id": {
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
      "ingress_application_gateway": {
        "block": {
          "attributes": {
            "effective_gateway_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "gateway_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "gateway_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "ingress_application_gateway_identity": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                [
                  "object",
                  {
                    "client_id": "string",
                    "object_id": "string",
                    "user_assigned_identity_id": "string"
                  }
                ]
              ]
            },
            "subnet_cidr": {
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
        "nesting_mode": "list"
      },
      "key_vault_secrets_provider": {
        "block": {
          "attributes": {
            "secret_identity": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                [
                  "object",
                  {
                    "client_id": "string",
                    "object_id": "string",
                    "user_assigned_identity_id": "string"
                  }
                ]
              ]
            },
            "secret_rotation_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "secret_rotation_interval": {
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
      "kubelet_identity": {
        "block": {
          "attributes": {
            "client_id": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "object_id": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "user_assigned_identity_id": {
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
      "linux_profile": {
        "block": {
          "attributes": {
            "admin_username": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "ssh_key": {
              "block": {
                "attributes": {
                  "key_data": {
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
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "maintenance_window": {
        "block": {
          "block_types": {
            "allowed": {
              "block": {
                "attributes": {
                  "day": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "hours": {
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "set",
                      "number"
                    ]
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "set"
            },
            "not_allowed": {
              "block": {
                "attributes": {
                  "end": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "start": {
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
      },
      "network_profile": {
        "block": {
          "attributes": {
            "dns_service_ip": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "docker_bridge_cidr": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "load_balancer_sku": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "network_mode": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "network_plugin": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "network_policy": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "outbound_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "pod_cidr": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "service_cidr": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "load_balancer_profile": {
              "block": {
                "attributes": {
                  "effective_outbound_ips": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "idle_timeout_in_minutes": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "managed_outbound_ip_count": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "outbound_ip_address_ids": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "outbound_ip_prefix_ids": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "outbound_ports_allocated": {
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
            "nat_gateway_profile": {
              "block": {
                "attributes": {
                  "effective_outbound_ips": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "idle_timeout_in_minutes": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "managed_outbound_ip_count": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
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
      "oms_agent": {
        "block": {
          "attributes": {
            "log_analytics_workspace_id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "oms_agent_identity": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                [
                  "object",
                  {
                    "client_id": "string",
                    "object_id": "string",
                    "user_assigned_identity_id": "string"
                  }
                ]
              ]
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "role_based_access_control": {
        "block": {
          "attributes": {
            "enabled": {
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            }
          },
          "block_types": {
            "azure_active_directory": {
              "block": {
                "attributes": {
                  "admin_group_object_ids": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "azure_rbac_enabled": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "client_app_id": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "managed": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "server_app_id": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "server_app_secret": {
                    "description_kind": "plain",
                    "optional": true,
                    "sensitive": true,
                    "type": "string"
                  },
                  "tenant_id": {
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
            }
          },
          "deprecated": true,
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "service_principal": {
        "block": {
          "attributes": {
            "client_id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "client_secret": {
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
      "windows_profile": {
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
            "license": {
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
  "version": 2
}`

func AzurermKubernetesClusterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermKubernetesCluster), &result)
	return &result
}
