package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermNewRelicMonitor = `{
  "block": {
    "attributes": {
      "account_creation_source": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "account_id": {
        "computed": true,
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
      "ingestion_key": {
        "description_kind": "plain",
        "optional": true,
        "sensitive": true,
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
      "org_creation_source": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "organization_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "resource_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "user_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "plan": {
        "block": {
          "attributes": {
            "billing_cycle": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "effective_date": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "plan_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "usage_type": {
              "description_kind": "plain",
              "optional": true,
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
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "single"
      },
      "user": {
        "block": {
          "attributes": {
            "email": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "first_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "last_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "phone_number": {
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
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AzurermNewRelicMonitorSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermNewRelicMonitor), &result)
	return &result
}
