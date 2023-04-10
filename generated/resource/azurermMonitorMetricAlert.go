package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermMonitorMetricAlert = `{
  "block": {
    "attributes": {
      "auto_mitigate": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
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
      "frequency": {
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
      "severity": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "tags": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "target_resource_location": {
        "computed": true,
        "description": "The location of the target pluginsdk. Required when using subscription, resource group scope or multiple scopes.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "target_resource_type": {
        "computed": true,
        "description": "The resource type (e.g. Microsoft.Compute/virtualMachines) of the target pluginsdk. Required when using subscription, resource group scope or multiple scopes.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "window_size": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
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
        "nesting_mode": "set"
      },
      "application_insights_web_test_location_availability_criteria": {
        "block": {
          "attributes": {
            "component_id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "failed_location_count": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "web_test_id": {
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
      "criteria": {
        "block": {
          "attributes": {
            "aggregation": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "metric_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "metric_namespace": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "operator": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "skip_metric_validation": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "threshold": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            }
          },
          "block_types": {
            "dimension": {
              "block": {
                "attributes": {
                  "name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "operator": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "values": {
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "list",
                      "string"
                    ]
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
      "dynamic_criteria": {
        "block": {
          "attributes": {
            "aggregation": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "alert_sensitivity": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "evaluation_failure_count": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "evaluation_total_count": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "ignore_data_before": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "metric_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "metric_namespace": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "operator": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "skip_metric_validation": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "block_types": {
            "dimension": {
              "block": {
                "attributes": {
                  "name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "operator": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "values": {
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "list",
                      "string"
                    ]
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

func AzurermMonitorMetricAlertSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermMonitorMetricAlert), &result)
	return &result
}
