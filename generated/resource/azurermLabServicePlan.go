package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermLabServicePlan = `{
  "block": {
    "attributes": {
      "allowed_regions": {
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
      },
      "default_network_subnet_id": {
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
      "shared_gallery_id": {
        "description_kind": "plain",
        "optional": true,
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
      "default_auto_shutdown": {
        "block": {
          "attributes": {
            "disconnect_delay": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "idle_delay": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "no_connect_delay": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "shutdown_on_idle": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "default_connection": {
        "block": {
          "attributes": {
            "client_rdp_access": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "client_ssh_access": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "web_rdp_access": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "web_ssh_access": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "support": {
        "block": {
          "attributes": {
            "email": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "instructions": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "phone": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "url": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
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
    "deprecated": true,
    "description_kind": "plain"
  },
  "version": 0
}`

func AzurermLabServicePlanSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermLabServicePlan), &result)
	return &result
}
