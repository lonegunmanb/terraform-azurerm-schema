package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermGalleryApplicationVersion = `{
  "block": {
    "attributes": {
      "config_file": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "enable_health_check": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "end_of_life_date": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "exclude_from_latest": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "gallery_application_id": {
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
      "package_file": {
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
      "manage_action": {
        "block": {
          "attributes": {
            "install": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "remove": {
              "description_kind": "plain",
              "required": true,
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
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "source": {
        "block": {
          "attributes": {
            "default_configuration_link": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "media_link": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "target_region": {
        "block": {
          "attributes": {
            "exclude_from_latest": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "regional_replica_count": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "storage_account_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "min_items": 1,
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

func AzurermGalleryApplicationVersionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermGalleryApplicationVersion), &result)
	return &result
}
