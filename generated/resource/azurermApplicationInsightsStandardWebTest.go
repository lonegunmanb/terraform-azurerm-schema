package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermApplicationInsightsStandardWebTest = `{
  "block": {
    "attributes": {
      "application_insights_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
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
        "type": "number"
      },
      "geo_locations": {
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
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
      "retry_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "synthetic_monitor_id": {
        "computed": true,
        "description_kind": "plain",
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
      "timeout": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      }
    },
    "block_types": {
      "request": {
        "block": {
          "attributes": {
            "body": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "follow_redirects_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "http_verb": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "parse_dependent_requests_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "url": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "header": {
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
      },
      "validation_rules": {
        "block": {
          "attributes": {
            "expected_status_code": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "ssl_cert_remaining_lifetime": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "ssl_check_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "block_types": {
            "content": {
              "block": {
                "attributes": {
                  "content_match": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "ignore_case": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "pass_if_text_found": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
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
  "version": 0
}`

func AzurermApplicationInsightsStandardWebTestSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermApplicationInsightsStandardWebTest), &result)
	return &result
}
