package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermMaintenanceConfiguration = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "in_guest_user_patch_mode": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "install_patches": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "linux": [
                "list",
                [
                  "object",
                  {
                    "classifications_to_include": [
                      "list",
                      "string"
                    ],
                    "package_names_mask_to_exclude": [
                      "list",
                      "string"
                    ],
                    "package_names_mask_to_include": [
                      "list",
                      "string"
                    ]
                  }
                ]
              ],
              "reboot": "string",
              "windows": [
                "list",
                [
                  "object",
                  {
                    "classifications_to_include": [
                      "list",
                      "string"
                    ],
                    "kb_numbers_to_exclude": [
                      "list",
                      "string"
                    ],
                    "kb_numbers_to_include": [
                      "list",
                      "string"
                    ]
                  }
                ]
              ]
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
      "properties": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "resource_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "scope": {
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
      "visibility": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "window": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "duration": "string",
              "expiration_date_time": "string",
              "recur_every": "string",
              "start_date_time": "string",
              "time_zone": "string"
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

func AzurermMaintenanceConfigurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermMaintenanceConfiguration), &result)
	return &result
}
