package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermOracleCloudVmCluster = `{
  "block": {
    "attributes": {
      "backup_subnet_cidr": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "cloud_exadata_infrastructure_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "cluster_name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "cpu_core_count": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "data_storage_percentage": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "data_storage_size_in_tbs": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "db_node_storage_size_in_gbs": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "db_servers": {
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
      },
      "display_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "domain": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "gi_version": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "hostname": {
        "description_kind": "plain",
        "required": true,
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
      "license_model": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "local_backup_enabled": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "location": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "memory_size_in_gbs": {
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
      "ocid": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "resource_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "scan_listener_port_tcp": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "scan_listener_port_tcp_ssl": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "sparse_diskgroup_enabled": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "ssh_public_keys": {
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
      },
      "subnet_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "system_version": {
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
      "time_zone": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "virtual_network_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "zone_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "data_collection_options": {
        "block": {
          "attributes": {
            "diagnostics_events_enabled": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "health_monitoring_enabled": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "incident_logs_enabled": {
              "computed": true,
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

func AzurermOracleCloudVmClusterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermOracleCloudVmCluster), &result)
	return &result
}
