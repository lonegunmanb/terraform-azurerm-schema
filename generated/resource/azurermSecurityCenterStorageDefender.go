package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermSecurityCenterStorageDefender = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "malware_scanning_on_upload_cap_gb_per_month": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "malware_scanning_on_upload_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "override_subscription_settings_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "scan_results_event_grid_topic_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "sensitive_data_discovery_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "storage_account_id": {
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

func AzurermSecurityCenterStorageDefenderSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermSecurityCenterStorageDefender), &result)
	return &result
}
