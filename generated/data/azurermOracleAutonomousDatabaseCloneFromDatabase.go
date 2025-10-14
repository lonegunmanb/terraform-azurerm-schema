package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermOracleAutonomousDatabaseCloneFromDatabase = `{
  "block": {
    "attributes": {
      "actual_used_data_storage_size_in_tb": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "allocated_storage_size_in_tb": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "allowed_ip_addresses": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
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
      "compute_model": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "connection_strings": {
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
      "customer_contacts": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "data_storage_size_in_gb": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "data_storage_size_in_tb": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "database_version": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "database_workload": {
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
      "in_memory_area_in_gb": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
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
      "local_adg_auto_failover_max_data_loss_limit_in_seconds": {
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
      "long_term_backup_schedule": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "enabled": "bool",
              "repeat_cadence": "string",
              "retention_period_in_days": "number",
              "time_of_backup_in_utc": "string"
            }
          ]
        ]
      },
      "memory_per_oracle_compute_unit_in_gb": {
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
      "next_long_term_backup_timestamp": {
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
      "peer_database_ids": {
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
      "private_endpoint_url": {
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
      "reconnect_clone_enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "refreshable_clone": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "refreshable_status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
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
      "source_autonomous_database_id": {
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
          "string"
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
      "time_created_in_utc": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "time_data_guard_role_changed_in_utc": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "time_deletion_of_free_autonomous_database_in_utc": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "time_local_data_guard_enabled_in_utc": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "time_maintenance_begin_in_utc": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "time_maintenance_end_in_utc": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "time_of_last_failover_in_utc": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "time_of_last_refresh_in_utc": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "time_of_last_refresh_point_in_utc": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "time_of_last_switchover_in_utc": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "time_reclamation_of_free_autonomous_database_in_utc": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "time_until_reconnect_in_utc": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "used_data_storage_size_in_gb": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "used_data_storage_size_in_tb": {
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

func AzurermOracleAutonomousDatabaseCloneFromDatabaseSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermOracleAutonomousDatabaseCloneFromDatabase), &result)
	return &result
}
