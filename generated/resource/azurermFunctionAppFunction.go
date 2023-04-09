package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermFunctionAppFunction = `{
  "block": {
    "attributes": {
      "config_json": {
        "description": "The config for this Function in JSON format.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "config_url": {
        "computed": true,
        "description": "The URL of the configuration JSON.",
        "description_kind": "plain",
        "type": "string"
      },
      "enabled": {
        "description": "Should this function be enabled. Defaults to ` + "`" + `true` + "`" + `.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "function_app_id": {
        "description": "The ID of the Function App in which this function should reside.",
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
      "invocation_url": {
        "computed": true,
        "description": "The invocation URL.",
        "description_kind": "plain",
        "type": "string"
      },
      "language": {
        "description": "The language the Function is written in.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description": "The name of the function.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "script_root_path_url": {
        "computed": true,
        "description": "The Script root path URL.",
        "description_kind": "plain",
        "type": "string"
      },
      "script_url": {
        "computed": true,
        "description": "The script URL.",
        "description_kind": "plain",
        "type": "string"
      },
      "secrets_file_url": {
        "computed": true,
        "description": "The URL for the Secrets File.",
        "description_kind": "plain",
        "type": "string"
      },
      "test_data": {
        "description": "The test data for the function.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "test_data_url": {
        "computed": true,
        "description": "The Test data URL.",
        "description_kind": "plain",
        "type": "string"
      },
      "url": {
        "computed": true,
        "description": "The function URL.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "file": {
        "block": {
          "attributes": {
            "content": {
              "description": "The content of the file.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "name": {
              "description": "The filename of the file to be uploaded.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
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

func AzurermFunctionAppFunctionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermFunctionAppFunction), &result)
	return &result
}
