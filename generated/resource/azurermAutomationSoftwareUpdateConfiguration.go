package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermAutomationSoftwareUpdateConfiguration = `{
  "block": {
    "attributes": {
      "automation_account_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "duration": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "error_code": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "error_meesage": {
        "computed": true,
        "deprecated": true,
        "description_kind": "plain",
        "type": "string"
      },
      "error_message": {
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
      "non_azure_computer_names": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "operating_system": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "virtual_machine_ids": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      }
    },
    "block_types": {
      "linux": {
        "block": {
          "attributes": {
            "classification_included": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "excluded_packages": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "included_packages": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "reboot": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "post_task": {
        "block": {
          "attributes": {
            "parameters": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "source": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "pre_task": {
        "block": {
          "attributes": {
            "parameters": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "source": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "schedule": {
        "block": {
          "attributes": {
            "advanced_month_days": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "number"
              ]
            },
            "advanced_week_days": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "creation_time": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "description": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "expiry_time": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "expiry_time_offset_minutes": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "frequency": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "interval": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "is_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "last_modified_time": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "next_run": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "next_run_offset_minutes": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "start_time": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "start_time_offset_minutes": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "time_zone": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "monthly_occurrence": {
              "block": {
                "attributes": {
                  "day": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "occurrence": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
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
      "target": {
        "block": {
          "block_types": {
            "azure_query": {
              "block": {
                "attributes": {
                  "locations": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "scope": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "tag_filter": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "tags": {
                    "block": {
                      "attributes": {
                        "tag": {
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
            "non_azure_query": {
              "block": {
                "attributes": {
                  "function_alias": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "workspace_id": {
                    "description_kind": "plain",
                    "optional": true,
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
      },
      "windows": {
        "block": {
          "attributes": {
            "classification_included": {
              "computed": true,
              "deprecated": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "classifications_included": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "excluded_knowledge_base_numbers": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "included_knowledge_base_numbers": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "reboot": {
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
  "version": 0
}`

func AzurermAutomationSoftwareUpdateConfigurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermAutomationSoftwareUpdateConfiguration), &result)
	return &result
}
