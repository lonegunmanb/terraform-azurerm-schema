package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermOracleAutonomousDatabase = `{
  "block": {
    "attributes": {
      "actual_used_data_storage_size_in_tbs": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "allocated_storage_size_in_tbs": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "allowed_ips": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "number"
        ]
      },
      "auto_scaling_enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "auto_scaling_for_storage_enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "autonomous_database_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "available_upgrade_versions": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "backup_retention_period_in_days": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "character_set": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "compute_count": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "cpu_core_count": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "data_storage_size_in_gbs": {
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
      "db_version": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "display_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "failed_data_recovery_in_seconds": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "in_memory_area_in_gbs": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "lifecycle_details": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "local_adg_auto_failover_max_data_loss_limit": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "local_data_guard_enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "location": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "memory_per_oracle_compute_unit_in_gbs": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "mtls_connection_required": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "national_character_set": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "next_long_term_backup_time_stamp": {
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
      "peer_db_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "peer_db_ids": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "preview": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "preview_version_with_service_terms_accepted": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "private_endpoint": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "private_endpoint_ip": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "private_endpoint_label": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "provisionable_cpus": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "number"
        ]
      },
      "remote_data_guard_enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "resource_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "service_console_url": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "sql_web_developer_url": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "subnet_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "supported_regions_to_clone_to": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "number"
        ]
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
      "time_data_guard_role_changed": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "time_deletion_of_free_autonomous_database": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "time_local_data_guard_enabled_on": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "time_maintenance_begin": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "time_maintenance_end": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "time_of_last_failover": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "time_of_last_refresh": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "time_of_last_refresh_point": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "time_of_last_switchover": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "time_reclamation_of_free_autonomous_database": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "used_data_storage_size_in_gbs": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "used_data_storage_size_in_tbs": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "virtual_network_id": {
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

func AzurermOracleAutonomousDatabaseSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermOracleAutonomousDatabase), &result)
	return &result
}
