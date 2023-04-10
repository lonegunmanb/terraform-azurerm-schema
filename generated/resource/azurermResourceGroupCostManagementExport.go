package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermResourceGroupCostManagementExport = `{
  "block": {
    "attributes": {
      "active": {
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
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "recurrence_period_end_date": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "recurrence_period_start_date": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "recurrence_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "resource_group_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "export_data_options": {
        "block": {
          "attributes": {
            "time_frame": {
              "description_kind": "plain",
              "required": true,
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
        "min_items": 1,
        "nesting_mode": "list"
      },
      "export_data_storage_location": {
        "block": {
          "attributes": {
            "container_id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "root_folder_path": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
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

func AzurermResourceGroupCostManagementExportSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermResourceGroupCostManagementExport), &result)
	return &result
}
