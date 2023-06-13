package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermSiteRecoveryReplicationRecoveryPlan = `{
  "block": {
    "attributes": {
      "azure_to_azure_settings": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "primary_edge_zone": "string",
              "primary_zone": "string",
              "recovery_edge_zone": "string",
              "recovery_zone": "string"
            }
          ]
        ]
      },
      "failover_deployment_model": {
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
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "recovery_group": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          [
            "object",
            {
              "post_action": [
                "set",
                [
                  "list",
                  [
                    "object",
                    {
                      "fabric_location": "string",
                      "fail_over_directions": [
                        "set",
                        "string"
                      ],
                      "fail_over_types": [
                        "set",
                        "string"
                      ],
                      "manual_action_instruction": "string",
                      "name": "string",
                      "runbook_id": "string",
                      "script_path": "string",
                      "type": "string"
                    }
                  ]
                ]
              ],
              "pre_action": [
                "set",
                [
                  "list",
                  [
                    "object",
                    {
                      "fabric_location": "string",
                      "fail_over_directions": [
                        "set",
                        "string"
                      ],
                      "fail_over_types": [
                        "set",
                        "string"
                      ],
                      "manual_action_instruction": "string",
                      "name": "string",
                      "runbook_id": "string",
                      "script_path": "string",
                      "type": "string"
                    }
                  ]
                ]
              ],
              "replicated_protected_items": [
                "list",
                "string"
              ],
              "type": "string"
            }
          ]
        ]
      },
      "recovery_vault_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "source_recovery_fabric_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "target_recovery_fabric_id": {
        "computed": true,
        "description_kind": "plain",
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

func AzurermSiteRecoveryReplicationRecoveryPlanSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermSiteRecoveryReplicationRecoveryPlan), &result)
	return &result
}
