package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermSiteRecoveryVmwareReplicatedVm = `{
  "block": {
    "attributes": {
      "appliance_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "default_log_storage_account_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "default_recovery_disk_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "default_target_disk_encryption_set_id": {
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
      "license_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "multi_vm_group_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "physical_server_credential_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "recovery_replication_policy_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "recovery_vault_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "source_vm_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "target_availability_set_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "target_boot_diagnostics_storage_account_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "target_network_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "target_proximity_placement_group_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "target_resource_group_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "target_vm_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "target_vm_size": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "target_zone": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "test_network_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "managed_disk": {
        "block": {
          "attributes": {
            "disk_id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "log_storage_account_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "target_disk_encryption_set_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "target_disk_type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "network_interface": {
        "block": {
          "attributes": {
            "is_primary": {
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            },
            "source_mac_address": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "target_static_ip": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "target_subnet_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "test_subnet_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
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

func AzurermSiteRecoveryVmwareReplicatedVmSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermSiteRecoveryVmwareReplicatedVm), &result)
	return &result
}
