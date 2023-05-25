package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermVirtualDesktopHostPool = `{
  "block": {
    "attributes": {
      "custom_rdp_properties": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "friendly_name": {
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
      "load_balancer_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "location": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "maximum_sessions_allowed": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "personal_desktop_assignment_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "preferred_app_group_type": {
        "description": "Preferred App Group type to display",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "resource_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "start_vm_on_connect": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "tags": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "validate_environment": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      }
    },
    "block_types": {
      "registration_info": {
        "block": {
          "attributes": {
            "expiration_date": {
              "deprecated": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "reset_token": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "token": {
              "computed": true,
              "description_kind": "plain",
              "sensitive": true,
              "type": "string"
            }
          },
          "description": "This block is now non-functional and will be removed in version 3.0 of the Azure Provider - use the ` + "`" + `azurerm_virtual_desktop_host_pool_registration_info` + "`" + ` resource instead.",
          "description_kind": "plain"
        },
        "max_items": 1,
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
  "version": 1
}`

func AzurermVirtualDesktopHostPoolSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermVirtualDesktopHostPool), &result)
	return &result
}
