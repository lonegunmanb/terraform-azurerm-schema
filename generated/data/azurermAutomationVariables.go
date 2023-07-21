package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermAutomationVariables = `{
  "block": {
    "attributes": {
      "automation_account_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "bool": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "description": "string",
              "encrypted": "bool",
              "id": "string",
              "name": "string",
              "value": "bool"
            }
          ]
        ]
      },
      "datetime": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "description": "string",
              "encrypted": "bool",
              "id": "string",
              "name": "string",
              "value": "string"
            }
          ]
        ]
      },
      "encrypted": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "description": "string",
              "encrypted": "bool",
              "id": "string",
              "name": "string",
              "value": "string"
            }
          ]
        ]
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "int": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "description": "string",
              "encrypted": "bool",
              "id": "string",
              "name": "string",
              "value": "number"
            }
          ]
        ]
      },
      "null": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "description": "string",
              "encrypted": "bool",
              "id": "string",
              "name": "string",
              "value": "string"
            }
          ]
        ]
      },
      "string": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "description": "string",
              "encrypted": "bool",
              "id": "string",
              "name": "string",
              "value": "string"
            }
          ]
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

func AzurermAutomationVariablesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermAutomationVariables), &result)
	return &result
}
