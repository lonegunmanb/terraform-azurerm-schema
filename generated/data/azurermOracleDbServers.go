package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermOracleDbServers = `{
  "block": {
    "attributes": {
      "cloud_exadata_infrastructure_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "db_servers": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "autonomous_virtual_machine_ds": [
                "list",
                "string"
              ],
              "autonomous_vm_cluster_ids": [
                "list",
                "string"
              ],
              "compartment_id": "string",
              "compute_model": "string",
              "cpu_core_count": "number",
              "db_node_ids": [
                "list",
                "string"
              ],
              "db_node_storage_size_in_gbs": "number",
              "display_name": "string",
              "exadata_infrastructure_id": "string",
              "lifecycle_details": "string",
              "lifecycle_state": "string",
              "max_cpu_count": "number",
              "max_db_node_storage_in_gbs": "number",
              "max_memory_in_gbs": "number",
              "memory_size_in_gbs": "number",
              "ocid": "string",
              "shape": "string",
              "time_created": "string",
              "vm_cluster_ids": [
                "list",
                "string"
              ]
            }
          ]
        ]
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "resource_group_name": {
        "description_kind": "plain",
        "required": true,
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

func AzurermOracleDbServersSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermOracleDbServers), &result)
	return &result
}
