package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermOracleExadataInfrastructure = `{
  "block": {
    "attributes": {
      "compute_count": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "customer_contacts": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "database_server_type": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "display_name": {
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
      "shape": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "storage_count": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "storage_server_type": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
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
      "zones": {
        "description_kind": "plain",
        "required": true,
        "type": [
          "set",
          "string"
        ]
      }
    },
    "block_types": {
      "maintenance_window": {
        "block": {
          "attributes": {
            "days_of_week": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "hours_of_day": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "number"
              ]
            },
            "lead_time_in_weeks": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "months": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "patching_mode": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "preference": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "weeks_of_month": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "number"
              ]
            }
          },
          "description_kind": "plain"
        },
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

func AzurermOracleExadataInfrastructureSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermOracleExadataInfrastructure), &result)
	return &result
}
