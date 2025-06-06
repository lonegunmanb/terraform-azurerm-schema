package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermNetappVolume = `{
  "block": {
    "attributes": {
      "account_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "azure_vmware_data_store_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "create_from_snapshot_resource_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "encryption_key_source": {
        "computed": true,
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
      "kerberos_enabled": {
        "description": "Enable to allow Kerberos secured volumes. Requires appropriate export rules as well as the parent ` + "`" + `azurerm_netapp_account` + "`" + ` having a defined AD connection.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "key_vault_private_endpoint_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "large_volume_enabled": {
        "description": "Indicates whether the volume is a large volume.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "location": {
        "description_kind": "plain",
        "required": true,
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
        "optional": true,
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
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "resource_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "security_style": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "service_level": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "smb3_protocol_encryption_enabled": {
        "description": "SMB3 encryption option should be used only for SMB/DualProtocol volumes. Using it for any other workloads is not supported.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "smb_access_based_enumeration_enabled": {
        "description": "Enable access based enumeration setting for SMB/Dual Protocol volume. When enabled, users who do not have permission to access a shared folder or file underneath it, do not see that shared resource displayed in their environment.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "smb_continuous_availability_enabled": {
        "description": "Continuous availability option should be used only for SQL and FSLogix workloads. Using it for any other SMB workloads is not supported.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "smb_non_browsable_enabled": {
        "description": "Enable non browsable share setting for SMB/Dual Protocol volume. When enabled, it restricts windows clients to browse the share",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "snapshot_directory_visible": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "storage_quota_in_gb": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "subnet_id": {
        "description_kind": "plain",
        "required": true,
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
      "throughput_in_mibps": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "volume_path": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "zone": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "data_protection_backup_policy": {
        "block": {
          "attributes": {
            "backup_policy_id": {
              "description": "The ID of the backup policy to associate with this volume.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "backup_vault_id": {
              "description": "The ID of the backup vault to associate with this volume.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "policy_enabled": {
              "description": "If set to false, the backup policy will not be enabled on this volume, thus disabling scheduled backups.",
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
      "data_protection_replication": {
        "block": {
          "attributes": {
            "endpoint_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "remote_volume_location": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "remote_volume_resource_id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "replication_frequency": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "data_protection_snapshot_policy": {
        "block": {
          "attributes": {
            "snapshot_policy_id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "export_policy_rule": {
        "block": {
          "attributes": {
            "allowed_clients": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "set",
                "string"
              ]
            },
            "kerberos_5_read_only_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "kerberos_5_read_write_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "kerberos_5i_read_only_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "kerberos_5i_read_write_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "kerberos_5p_read_only_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "kerberos_5p_read_write_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "protocols_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "root_access_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "rule_index": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "unix_read_only": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "unix_read_write": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 5,
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

func AzurermNetappVolumeSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermNetappVolume), &result)
	return &result
}
