package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermSharedImage = `{
  "block": {
    "attributes": {
      "architecture": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "eula": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "gallery_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "hyper_v_generation": {
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
      "identifier": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "offer": "string",
              "publisher": "string",
              "sku": "string"
            }
          ]
        ]
      },
      "location": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "os_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "privacy_statement_uri": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "purchase_plan": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "name": "string",
              "product": "string",
              "publisher": "string"
            }
          ]
        ]
      },
      "release_note_uri": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "resource_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "specialized": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      }
    },
    "block_types": {
      "timeouts": {
        "block": {
          "attributes": {
            "read": {
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

func AzurermSharedImageSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermSharedImage), &result)
	return &result
}
