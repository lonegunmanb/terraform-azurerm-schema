package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermPolicyVirtualMachineConfigurationAssignment = `{
  "block": {
    "attributes": {
      "assignment_hash": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "compliance_status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "content_hash": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "content_uri": {
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
      "last_compliance_status_checked": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "latest_report_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "resource_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "virtual_machine_name": {
        "description_kind": "plain",
        "required": true,
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

func AzurermPolicyVirtualMachineConfigurationAssignmentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermPolicyVirtualMachineConfigurationAssignment), &result)
	return &result
}
