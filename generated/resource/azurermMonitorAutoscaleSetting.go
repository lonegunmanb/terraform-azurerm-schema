package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermMonitorAutoscaleSetting = `{
  "block": {
    "attributes": {
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
      "tags": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "target_resource_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "notification": {
        "block": {
          "block_types": {
            "email": {
              "block": {
                "attributes": {
                  "custom_emails": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "send_to_subscription_administrator": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "send_to_subscription_co_administrator": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "webhook": {
              "block": {
                "attributes": {
                  "properties": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  },
                  "service_uri": {
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
        "nesting_mode": "list"
      },
      "profile": {
        "block": {
          "attributes": {
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "capacity": {
              "block": {
                "attributes": {
                  "default": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "maximum": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "minimum": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            },
            "fixed_date": {
              "block": {
                "attributes": {
                  "end": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "start": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "timezone": {
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
            "recurrence": {
              "block": {
                "attributes": {
                  "days": {
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "hours": {
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "list",
                      "number"
                    ]
                  },
                  "minutes": {
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "list",
                      "number"
                    ]
                  },
                  "timezone": {
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
            "rule": {
              "block": {
                "block_types": {
                  "metric_trigger": {
                    "block": {
                      "attributes": {
                        "divide_by_instance_count": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "metric_name": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "metric_namespace": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "metric_resource_id": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "operator": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "statistic": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "threshold": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        },
                        "time_aggregation": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "time_grain": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "time_window": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "dimensions": {
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
                    "min_items": 1,
                    "nesting_mode": "list"
                  },
                  "scale_action": {
                    "block": {
                      "attributes": {
                        "cooldown": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "direction": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "type": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "value": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
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
              "max_items": 10,
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 20,
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
  "version": 2
}`

func AzurermMonitorAutoscaleSettingSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermMonitorAutoscaleSetting), &result)
	return &result
}
