package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermOracleCloudVmCluster = `{
  "block": {
    "attributes": {
      "backup_subnet_cidr": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "cloud_exadata_infrastructure_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "cluster_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "compartment_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "compute_nodes": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "cpu_core_count": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "data_collection_options": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "diagnostics_events_enabled": "bool",
              "health_monitoring_enabled": "bool",
              "incident_logs_enabled": "bool"
            }
          ]
        ]
      },
      "data_storage_percentage": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "data_storage_size_in_tbs": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "db_node_storage_size_in_gbs": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "db_servers": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "disk_redundancy": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "display_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "domain": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "gi_version": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "hostname": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "hostname_actual": {
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
      "iorm_config_cache": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "db_plans": [
                "list",
                [
                  "object",
                  {
                    "db_name": "string",
                    "flash_cache_limit": "string",
                    "share": "number"
                  }
                ]
              ],
              "lifecycle_details": "string",
              "lifecycle_state": "string",
              "objective": "string"
            }
          ]
        ]
      },
      "last_update_history_entry_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "license_model": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "lifecycle_details": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "lifecycle_state": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "listener_port": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "local_backup_enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "location": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "memory_size_in_gbs": {
        "computed": true,
        "description_kind": "plain",
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
        "type": "number"
      },
      "nsg_url": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "oci_url": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "ocid": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "ocpu_count": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "resource_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "scan_dns_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "scan_dns_record_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "scan_ip_ids": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "scan_listener_port_tcp": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "scan_listener_port_tcp_ssl": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "shape": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "sparse_diskgroup_enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "ssh_public_keys": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "storage_size_in_gbs": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "subnet_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "subnet_ocid": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "system_version": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "time_created": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "time_zone": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "vip_ods": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "virtual_network_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "zone_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "timeouts": {
        "block": {
          "attributes": {
            "read": {
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

func AzurermOracleCloudVmClusterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermOracleCloudVmCluster), &result)
	return &result
}
