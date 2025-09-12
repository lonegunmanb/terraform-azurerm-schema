package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermNetappVolume = `{
  "block": {
    "attributes": {
      "accept_grow_capacity_pool_for_short_term_clone_split": {
        "computed": true,
        "description": "The accept grow capacity pool for short term clone split property.",
        "description_kind": "plain",
        "type": "string"
      },
      "account_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "data_protection_backup_policy": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "backup_policy_id": "string",
              "backup_vault_id": "string",
              "policy_enabled": "bool"
            }
          ]
        ]
      },
      "data_protection_replication": {
        "computed": true,
        "description_kind": "plain",
        "type": [
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
        ]
      },
      "encryption_key_source": {
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
      "key_vault_private_endpoint_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "large_volume_enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "location": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "mount_ip_addresses": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "network_features": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "pool_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "protocols": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "resource_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "security_style": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "service_level": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "smb_access_based_enumeration_enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "smb_non_browsable_enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "storage_quota_in_gb": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "subnet_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "volume_path": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "zone": {
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

func AzurermNetappVolumeSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermNetappVolume), &result)
	return &result
}
