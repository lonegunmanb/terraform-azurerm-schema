package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermBastionHost = `{
  "block": {
    "attributes": {
      "copy_paste_enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "dns_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "file_copy_enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ip_configuration": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "name": "string",
              "public_ip_address_id": "string",
              "subnet_id": "string"
            }
          ]
        ]
      },
      "ip_connect_enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
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
      "scale_units": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "session_recording_enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "shareable_link_enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "sku": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "tunneling_enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
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

func AzurermBastionHostSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermBastionHost), &result)
	return &result
}
