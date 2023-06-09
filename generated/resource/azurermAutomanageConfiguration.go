package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermAutomanageConfiguration = `{
  "block": {
    "attributes": {
      "automation_account_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "boot_diagnostics_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "defender_for_cloud_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "guest_configuration_enabled": {
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
      "status_change_alert_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
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
      "antimalware": {
        "block": {
          "attributes": {
            "real_time_protection_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "scheduled_scan_day": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "scheduled_scan_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "scheduled_scan_time_in_minutes": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "scheduled_scan_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "exclusions": {
              "block": {
                "attributes": {
                  "extensions": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "paths": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "processes": {
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
        "max_items": 1,
        "nesting_mode": "list"
      },
      "azure_security_baseline": {
        "block": {
          "attributes": {
            "assignment_type": {
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
      "backup": {
        "block": {
          "attributes": {
            "instant_rp_retention_range_in_days": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "policy_name": {
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
            "retention_policy": {
              "block": {
                "attributes": {
                  "retention_policy_type": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "daily_schedule": {
                    "block": {
                      "attributes": {
                        "retention_times": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        }
                      },
                      "block_types": {
                        "retention_duration": {
                          "block": {
                            "attributes": {
                              "count": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "duration_type": {
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
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "weekly_schedule": {
                    "block": {
                      "attributes": {
                        "retention_times": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        }
                      },
                      "block_types": {
                        "retention_duration": {
                          "block": {
                            "attributes": {
                              "count": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "duration_type": {
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
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "schedule_policy": {
              "block": {
                "attributes": {
                  "schedule_policy_type": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "schedule_run_days": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "schedule_run_frequency": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "schedule_run_times": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
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

func AzurermAutomanageConfigurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermAutomanageConfiguration), &result)
	return &result
}
