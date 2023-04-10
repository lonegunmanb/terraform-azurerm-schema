package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermNetworkSecurityGroup = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
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
      "resource_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "security_rule": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "access": "string",
              "description": "string",
              "destination_address_prefix": "string",
              "destination_address_prefixes": [
                "set",
                "string"
              ],
              "destination_application_security_group_ids": [
                "set",
                "string"
              ],
              "destination_port_range": "string",
              "destination_port_ranges": [
                "set",
                "string"
              ],
              "direction": "string",
              "name": "string",
              "priority": "number",
              "protocol": "string",
              "source_address_prefix": "string",
              "source_address_prefixes": [
                "set",
                "string"
              ],
              "source_application_security_group_ids": [
                "set",
                "string"
              ],
              "source_port_range": "string",
              "source_port_ranges": [
                "set",
                "string"
              ]
            }
          ]
        ]
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

func AzurermNetworkSecurityGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermNetworkSecurityGroup), &result)
	return &result
}
