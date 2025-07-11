package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermOracleExadataInfrastructure = `{
  "block": {
    "attributes": {
      "activated_storage_count": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "additional_storage_count": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "available_storage_size_in_gbs": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "compute_count": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "compute_model": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "cpu_count": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "customer_contacts": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "data_storage_size_in_tbs": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "database_server_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "db_node_storage_size_in_gbs": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "db_server_version": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "display_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "estimated_patching_time": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "estimated_db_server_patching_time": "number",
              "estimated_network_switches_patching_time": "number",
              "estimated_storage_server_patching_time": "number",
              "total_estimated_patching_time": "number"
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
      "last_maintenance_run_id": {
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
      "location": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "maintenance_window": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "custom_action_timeout_enabled": "bool",
              "custom_action_timeout_in_mins": "number",
              "days_of_week": [
                "list",
                "string"
              ],
              "hours_of_day": [
                "list",
                "number"
              ],
              "lead_time_in_weeks": "number",
              "monthly_patching_enabled": "bool",
              "months": [
                "list",
                "string"
              ],
              "patching_mode": "string",
              "preference": "string",
              "weeks_of_month": [
                "list",
                "number"
              ]
            }
          ]
        ]
      },
      "max_cpu_count": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "max_data_storage_in_tbs": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "max_db_node_storage_size_in_gbs": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "max_memory_in_gbs": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "memory_size_in_gbs": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "monthly_db_server_version": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "monthly_storage_server_version": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "next_maintenance_run_id": {
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
      "resource_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "shape": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "storage_count": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "storage_server_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "storage_server_version": {
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
      "total_storage_size_in_gbs": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "zones": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
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

func AzurermOracleExadataInfrastructureSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermOracleExadataInfrastructure), &result)
	return &result
}
