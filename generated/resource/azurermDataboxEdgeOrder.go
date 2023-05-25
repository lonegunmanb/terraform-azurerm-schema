package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermDataboxEdgeOrder = `{
  "block": {
    "attributes": {
      "device_name": {
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
      "name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "resource_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "return_tracking": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          [
            "object",
            {
              "carrier_name": "string",
              "serial_number": "string",
              "tracking_id": "string",
              "tracking_url": "string"
            }
          ]
        ]
      },
      "serial_number": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "shipment_history": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          [
            "object",
            {
              "additional_details": [
                "map",
                "string"
              ],
              "comments": "string",
              "last_update": "string"
            }
          ]
        ]
      },
      "shipment_tracking": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          [
            "object",
            {
              "carrier_name": "string",
              "serial_number": "string",
              "tracking_id": "string",
              "tracking_url": "string"
            }
          ]
        ]
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "additional_details": [
                "map",
                "string"
              ],
              "comments": "string",
              "info": "string",
              "last_update": "string"
            }
          ]
        ]
      }
    },
    "block_types": {
      "contact": {
        "block": {
          "attributes": {
            "company_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "emails": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "set",
                "string"
              ]
            },
            "name": {
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
      },
      "shipment_address": {
        "block": {
          "attributes": {
            "address": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "list",
                "string"
              ]
            },
            "city": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "country": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "postal_code": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "state": {
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

func AzurermDataboxEdgeOrderSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermDataboxEdgeOrder), &result)
	return &result
}
