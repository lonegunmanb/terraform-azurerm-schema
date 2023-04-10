package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermSiteRecoveryReplicatedVm = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "managed_disk": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          [
            "object",
            {
              "disk_id": "string",
              "staging_storage_account_id": "string",
              "target_disk_encryption": [
                "list",
                [
                  "object",
                  {
                    "disk_encryption_key": [
                      "list",
                      [
                        "object",
                        {
                          "secret_url": "string",
                          "vault_id": "string"
                        }
                      ]
                    ],
                    "key_encryption_key": [
                      "list",
                      [
                        "object",
                        {
                          "key_url": "string",
                          "vault_id": "string"
                        }
                      ]
                    ]
                  }
                ]
              ],
              "target_disk_encryption_set_id": "string",
              "target_disk_type": "string",
              "target_replica_disk_type": "string",
              "target_resource_group_id": "string"
            }
          ]
        ]
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
      "network_interface": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          [
            "object",
            {
              "failover_test_public_ip_address_id": "string",
              "failover_test_static_ip": "string",
              "failover_test_subnet_name": "string",
              "is_primary": "bool",
              "recovery_public_ip_address_id": "string",
              "source_network_interface_id": "string",
              "target_static_ip": "string",
              "target_subnet_name": "string"
            }
          ]
        ]
      },
      "recovery_replication_policy_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "recovery_vault_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "resource_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "source_recovery_fabric_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "source_recovery_protection_container_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "source_vm_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "target_availability_set_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "target_boot_diagnostic_storage_account_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "target_capacity_reservation_group_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "target_edge_zone": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "target_network_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "target_proximity_placement_group_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "target_recovery_fabric_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "target_recovery_protection_container_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "target_resource_group_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "target_virtual_machine_scale_set_id": {
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
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "unmanaged_disk": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          [
            "object",
            {
              "disk_uri": "string",
              "staging_storage_account_id": "string",
              "target_storage_account_id": "string"
            }
          ]
        ]
      }
    },
    "block_types": {
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

func AzurermSiteRecoveryReplicatedVmSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermSiteRecoveryReplicatedVm), &result)
	return &result
}
