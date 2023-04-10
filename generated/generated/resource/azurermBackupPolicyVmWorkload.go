package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermBackupPolicyVmWorkload = `{
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
      "recovery_vault_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "resource_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "workload_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "protection_policy": {
        "block": {
          "attributes": {
            "policy_type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "backup": {
              "block": {
                "attributes": {
                  "frequency": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "frequency_in_minutes": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "time": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "weekdays": {
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
              "min_items": 1,
              "nesting_mode": "list"
            },
            "retention_daily": {
              "block": {
                "attributes": {
                  "count": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "retention_monthly": {
              "block": {
                "attributes": {
                  "count": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "format_type": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "monthdays": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "number"
                    ]
                  },
                  "weekdays": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "weeks": {
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
            "retention_weekly": {
              "block": {
                "attributes": {
                  "count": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "weekdays": {
                    "description_kind": "plain",
                    "required": true,
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
            "retention_yearly": {
              "block": {
                "attributes": {
                  "count": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "format_type": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "monthdays": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "number"
                    ]
                  },
                  "months": {
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "weekdays": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "weeks": {
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
            "simple_retention": {
              "block": {
                "attributes": {
                  "count": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
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
        "nesting_mode": "set"
      },
      "settings": {
        "block": {
          "attributes": {
            "compression_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "time_zone": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
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

func AzurermBackupPolicyVmWorkloadSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermBackupPolicyVmWorkload), &result)
	return &result
}
