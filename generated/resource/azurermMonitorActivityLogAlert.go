package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermMonitorActivityLogAlert = `{
  "block": {
    "attributes": {
      "description": {
        "description_kind": "plain",
        "optional": true,
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
      "location": {
        "description_kind": "plain",
        "optional": true,
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
      "scopes": {
        "description_kind": "plain",
        "required": true,
        "type": [
          "set",
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
      "action": {
        "block": {
          "attributes": {
            "action_group_id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "webhook_properties": {
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
        "nesting_mode": "list"
      },
      "criteria": {
        "block": {
          "attributes": {
            "caller": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "category": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "level": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "levels": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "operation_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "recommendation_category": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "recommendation_impact": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "recommendation_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "resource_group": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "resource_groups": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "resource_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "resource_ids": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "resource_provider": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "resource_providers": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "resource_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "resource_types": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "status": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "statuses": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "sub_status": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "sub_statuses": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            }
          },
          "block_types": {
            "resource_health": {
              "block": {
                "attributes": {
                  "current": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "previous": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "reason": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "service_health": {
              "block": {
                "attributes": {
                  "events": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "locations": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "services": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
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
  "version": 1
}`

func AzurermMonitorActivityLogAlertSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermMonitorActivityLogAlert), &result)
	return &result
}
