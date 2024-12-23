package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermRecoveryServicesVault = `{
  "block": {
    "attributes": {
      "classic_vmware_replication_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "cross_region_restore_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "immutability": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "location": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "public_network_access_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "resource_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "sku": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "soft_delete_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "storage_mode_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tags": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      }
    },
    "block_types": {
      "encryption": {
        "block": {
          "attributes": {
            "infrastructure_encryption_enabled": {
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            },
            "key_id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "use_system_assigned_identity": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "user_assigned_identity_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "identity": {
        "block": {
          "attributes": {
            "identity_ids": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "principal_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "tenant_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "type": {
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
      "monitoring": {
        "block": {
          "attributes": {
            "alerts_for_all_job_failures_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "alerts_for_critical_operation_failures_enabled": {
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

func AzurermRecoveryServicesVaultSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermRecoveryServicesVault), &result)
	return &result
}
