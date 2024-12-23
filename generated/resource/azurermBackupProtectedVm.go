package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermBackupProtectedVm = `{
  "block": {
    "attributes": {
      "backup_policy_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "exclude_disk_luns": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "number"
        ]
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "include_disk_luns": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "number"
        ]
      },
      "protection_state": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
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
      "source_vm_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
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

func AzurermBackupProtectedVmSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermBackupProtectedVm), &result)
	return &result
}
