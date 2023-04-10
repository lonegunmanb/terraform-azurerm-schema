package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermSentinelAlertRuleNrt = `{
  "block": {
    "attributes": {
      "alert_rule_template_guid": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "alert_rule_template_version": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "custom_details": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "display_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "enabled": {
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
      "log_analytics_workspace_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "query": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "severity": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "suppression_duration": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "suppression_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "tactics": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "techniques": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      }
    },
    "block_types": {
      "alert_details_override": {
        "block": {
          "attributes": {
            "description_format": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "display_name_format": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "severity_column_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "tactics_column_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "dynamic_property": {
              "block": {
                "attributes": {
                  "name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
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
        "nesting_mode": "list"
      },
      "entity_mapping": {
        "block": {
          "attributes": {
            "entity_type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "field_mapping": {
              "block": {
                "attributes": {
                  "column_name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "identifier": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 3,
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 5,
        "nesting_mode": "list"
      },
      "event_grouping": {
        "block": {
          "attributes": {
            "aggregation_method": {
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
      "incident": {
        "block": {
          "attributes": {
            "create_incident_enabled": {
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            }
          },
          "block_types": {
            "grouping": {
              "block": {
                "attributes": {
                  "by_alert_details": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "by_custom_details": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "by_entities": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "enabled": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "entity_matching_method": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "lookback_duration": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "reopen_closed_incidents": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "sentinel_entity_mapping": {
        "block": {
          "attributes": {
            "column_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 5,
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

func AzurermSentinelAlertRuleNrtSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermSentinelAlertRuleNrt), &result)
	return &result
}
