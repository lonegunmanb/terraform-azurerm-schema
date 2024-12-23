package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermNetappVolumeGroupSapHana = `{
  "block": {
    "attributes": {
      "account_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "application_identifier": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "group_description": {
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
      "location": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "resource_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "volume": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "capacity_pool_id": "string",
              "data_protection_replication": [
                "list",
                [
                  "object",
                  {
                    "endpoint_type": "string",
                    "remote_volume_location": "string",
                    "remote_volume_resource_id": "string",
                    "replication_frequency": "string"
                  }
                ]
              ],
              "data_protection_snapshot_policy": [
                "list",
                [
                  "object",
                  {
                    "snapshot_policy_id": "string"
                  }
                ]
              ],
              "export_policy_rule": [
                "list",
                [
                  "object",
                  {
                    "allowed_clients": "string",
                    "nfsv3_enabled": "bool",
                    "nfsv41_enabled": "bool",
                    "root_access_enabled": "bool",
                    "rule_index": "number",
                    "unix_read_only": "bool",
                    "unix_read_write": "bool"
                  }
                ]
              ],
              "id": "string",
              "mount_ip_addresses": [
                "list",
                "string"
              ],
              "name": "string",
              "protocols": [
                "list",
                "string"
              ],
              "proximity_placement_group_id": "string",
              "security_style": "string",
              "service_level": "string",
              "snapshot_directory_visible": "bool",
              "storage_quota_in_gb": "number",
              "subnet_id": "string",
              "tags": [
                "map",
                "string"
              ],
              "throughput_in_mibps": "number",
              "volume_path": "string",
              "volume_spec_name": "string"
            }
          ]
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

func AzurermNetappVolumeGroupSapHanaSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermNetappVolumeGroupSapHana), &result)
	return &result
}
