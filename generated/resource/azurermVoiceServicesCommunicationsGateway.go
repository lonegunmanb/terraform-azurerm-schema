package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermVoiceServicesCommunicationsGateway = `{
  "block": {
    "attributes": {
      "api_bridge": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "auto_generated_domain_name_label_scope": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "codecs": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "connectivity": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "e911_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "emergency_dial_strings": {
        "description_kind": "plain",
        "optional": true,
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
      "microsoft_teams_voicemail_pilot_number": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "on_prem_mcp_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "platforms": {
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
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
      }
    },
    "block_types": {
      "service_location": {
        "block": {
          "attributes": {
            "allowed_media_source_address_prefixes": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "allowed_signaling_source_address_prefixes": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "esrp_addresses": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "location": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "operator_addresses": {
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
        "min_items": 1,
        "nesting_mode": "set"
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

func AzurermVoiceServicesCommunicationsGatewaySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermVoiceServicesCommunicationsGateway), &result)
	return &result
}
