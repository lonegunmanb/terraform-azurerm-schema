package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermOracleDbNodes = `{
  "block": {
    "attributes": {
      "cloud_vm_cluster_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "db_nodes": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "additional_details": "string",
              "backup_ip_id": "string",
              "backup_vnic_2_id": "string",
              "backup_vnic_id": "string",
              "cpu_core_count": "number",
              "db_node_storage_size_in_gbs": "number",
              "db_server_id": "string",
              "db_system_id": "string",
              "fault_domain": "string",
              "host_ip_id": "string",
              "hostname": "string",
              "lifecycle_details": "string",
              "lifecycle_state": "string",
              "maintenance_type": "string",
              "memory_size_in_gbs": "number",
              "ocid": "string",
              "software_storage_size_in_gb": "number",
              "time_created": "string",
              "time_maintenance_window_end": "string",
              "time_maintenance_window_start": "string",
              "vnic_2_id": "string",
              "vnic_id": "string"
            }
          ]
        ]
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
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

func AzurermOracleDbNodesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermOracleDbNodes), &result)
	return &result
}
