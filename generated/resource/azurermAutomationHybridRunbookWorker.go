package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermAutomationHybridRunbookWorker = `{
  "block": {
    "attributes": {
      "automation_account_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ip": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "last_seen_date_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "registration_date_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "resource_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "vm_resource_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "worker_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "worker_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "worker_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "worker_type": {
        "computed": true,
        "description_kind": "plain",
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

func AzurermAutomationHybridRunbookWorkerSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermAutomationHybridRunbookWorker), &result)
	return &result
}
