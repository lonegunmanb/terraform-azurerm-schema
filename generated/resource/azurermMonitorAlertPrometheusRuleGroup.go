package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermMonitorAlertPrometheusRuleGroup = `{
  "block": {
    "attributes": {
      "cluster_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "interval": {
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
      "resource_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "rule_group_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "scopes": {
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
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
      "rule": {
        "block": {
          "attributes": {
            "alert": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "annotations": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "expression": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "for": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "labels": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "record": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "severity": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "block_types": {
            "action": {
              "block": {
                "attributes": {
                  "action_group_id": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "action_properties": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 5,
              "nesting_mode": "list"
            },
            "alert_resolution": {
              "block": {
                "attributes": {
                  "auto_resolved": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "time_to_resolve": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
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

func AzurermMonitorAlertPrometheusRuleGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermMonitorAlertPrometheusRuleGroup), &result)
	return &result
}
