package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermBotChannelDirectLineSpeech = `{
  "block": {
    "attributes": {
      "bot_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "cognitive_account_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "cognitive_service_access_key": {
        "description_kind": "plain",
        "required": true,
        "sensitive": true,
        "type": "string"
      },
      "cognitive_service_location": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "custom_speech_model_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "custom_voice_deployment_id": {
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
      "location": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "resource_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
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

func AzurermBotChannelDirectLineSpeechSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermBotChannelDirectLineSpeech), &result)
	return &result
}
