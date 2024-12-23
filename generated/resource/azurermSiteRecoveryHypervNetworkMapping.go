package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermSiteRecoveryHypervNetworkMapping = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "recovery_vault_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "source_network_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "source_system_center_virtual_machine_manager_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "target_network_id": {
        "description_kind": "plain",
        "required": true,
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

func AzurermSiteRecoveryHypervNetworkMappingSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermSiteRecoveryHypervNetworkMapping), &result)
	return &result
}
