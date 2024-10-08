package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermSiteRecoveryReplicationRecoveryPlan = `{
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
      "source_recovery_fabric_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "target_recovery_fabric_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "azure_to_azure_settings": {
        "block": {
          "attributes": {
            "primary_edge_zone": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "primary_zone": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "recovery_edge_zone": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "recovery_zone": {
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
      "boot_recovery_group": {
        "block": {
          "attributes": {
            "replicated_protected_items": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            }
          },
          "block_types": {
            "post_action": {
              "block": {
                "attributes": {
                  "fabric_location": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "fail_over_directions": {
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "fail_over_types": {
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "manual_action_instruction": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "runbook_id": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "script_path": {
                    "description_kind": "plain",
                    "optional": true,
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
              "nesting_mode": "list"
            },
            "pre_action": {
              "block": {
                "attributes": {
                  "fabric_location": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "fail_over_directions": {
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "fail_over_types": {
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "manual_action_instruction": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "runbook_id": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "script_path": {
                    "description_kind": "plain",
                    "optional": true,
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
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
        "min_items": 1,
        "nesting_mode": "list"
      },
      "failover_recovery_group": {
        "block": {
          "block_types": {
            "post_action": {
              "block": {
                "attributes": {
                  "fabric_location": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "fail_over_directions": {
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "fail_over_types": {
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "manual_action_instruction": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "runbook_id": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "script_path": {
                    "description_kind": "plain",
                    "optional": true,
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
              "nesting_mode": "list"
            },
            "pre_action": {
              "block": {
                "attributes": {
                  "fabric_location": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "fail_over_directions": {
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "fail_over_types": {
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "manual_action_instruction": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "runbook_id": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "script_path": {
                    "description_kind": "plain",
                    "optional": true,
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
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "shutdown_recovery_group": {
        "block": {
          "block_types": {
            "post_action": {
              "block": {
                "attributes": {
                  "fabric_location": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "fail_over_directions": {
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "fail_over_types": {
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "manual_action_instruction": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "runbook_id": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "script_path": {
                    "description_kind": "plain",
                    "optional": true,
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
              "nesting_mode": "list"
            },
            "pre_action": {
              "block": {
                "attributes": {
                  "fabric_location": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "fail_over_directions": {
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "fail_over_types": {
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "manual_action_instruction": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "runbook_id": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "script_path": {
                    "description_kind": "plain",
                    "optional": true,
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
              "nesting_mode": "list"
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

func AzurermSiteRecoveryReplicationRecoveryPlanSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermSiteRecoveryReplicationRecoveryPlan), &result)
	return &result
}
