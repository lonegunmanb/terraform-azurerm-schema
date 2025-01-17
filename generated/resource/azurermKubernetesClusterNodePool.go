package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermKubernetesClusterNodePool = `{
  "block": {
    "attributes": {
      "auto_scaling_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "capacity_reservation_group_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "eviction_policy": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "fips_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "gpu_instance": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "host_encryption_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "host_group_id": {
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
      "kubelet_disk_type": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "kubernetes_cluster_id": {
        "description_kind": "plain",
        "required": true,
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
      "mode": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
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
      "node_public_ip_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
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
      "os_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "pod_subnet_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
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
      "scale_down_mode": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "snapshot_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "spot_max_price": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "tags": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "temporary_name_for_rotation": {
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
      },
      "workload_runtime": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
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
      "node_network_profile": {
        "block": {
          "attributes": {
            "application_security_group_ids": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "node_public_ip_tags": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            }
          },
          "block_types": {
            "allowed_host_ports": {
              "block": {
                "attributes": {
                  "port_end": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "port_start": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "protocol": {
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
      "upgrade_settings": {
        "block": {
          "attributes": {
            "drain_timeout_in_minutes": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "max_surge": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "node_soak_duration_in_minutes": {
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
      "windows_profile": {
        "block": {
          "attributes": {
            "outbound_nat_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
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
  "version": 1
}`

func AzurermKubernetesClusterNodePoolSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermKubernetesClusterNodePool), &result)
	return &result
}
