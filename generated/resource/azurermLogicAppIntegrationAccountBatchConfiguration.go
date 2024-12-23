package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermLogicAppIntegrationAccountBatchConfiguration = `{
  "block": {
    "attributes": {
      "batch_group_name": {
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
      "integration_account_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "metadata": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
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
      }
    },
    "block_types": {
      "release_criteria": {
        "block": {
          "attributes": {
            "batch_size": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "message_count": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "block_types": {
            "recurrence": {
              "block": {
                "attributes": {
                  "end_time": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "frequency": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "interval": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "start_time": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "time_zone": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "schedule": {
                    "block": {
                      "attributes": {
                        "hours": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "set",
                            "number"
                          ]
                        },
                        "minutes": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "set",
                            "number"
                          ]
                        },
                        "month_days": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "set",
                            "number"
                          ]
                        },
                        "week_days": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "set",
                            "string"
                          ]
                        }
                      },
                      "block_types": {
                        "monthly": {
                          "block": {
                            "attributes": {
                              "week": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "number"
                              },
                              "weekday": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "description_kind": "plain"
                          },
                          "nesting_mode": "set"
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

func AzurermLogicAppIntegrationAccountBatchConfigurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermLogicAppIntegrationAccountBatchConfiguration), &result)
	return &result
}
