package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermTrafficManagerProfile = `{
  "block": {
    "attributes": {
      "fqdn": {
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
      "max_return": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "profile_status": {
        "description_kind": "plain",
        "optional": true,
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
      "traffic_routing_method": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "traffic_view_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      }
    },
    "block_types": {
      "dns_config": {
        "block": {
          "attributes": {
            "relative_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "ttl": {
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
      "monitor_config": {
        "block": {
          "attributes": {
            "expected_status_code_ranges": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "interval_in_seconds": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "path": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "port": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "protocol": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "timeout_in_seconds": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "tolerated_number_of_failures": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "block_types": {
            "custom_header": {
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

func AzurermTrafficManagerProfileSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermTrafficManagerProfile), &result)
	return &result
}
