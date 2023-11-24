package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermSharedImage = `{
  "block": {
    "attributes": {
      "accelerated_network_support_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "architecture": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "confidential_vm_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "confidential_vm_supported": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "disk_types_not_allowed": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "end_of_life_date": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "eula": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "gallery_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "hyper_v_generation": {
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
      "max_recommended_memory_in_gb": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "max_recommended_vcpu_count": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "min_recommended_memory_in_gb": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "min_recommended_vcpu_count": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "os_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "privacy_statement_uri": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "release_note_uri": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "resource_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "specialized": {
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
      "trusted_launch_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "trusted_launch_supported": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      }
    },
    "block_types": {
      "identifier": {
        "block": {
          "attributes": {
            "offer": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "publisher": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "sku": {
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
      "purchase_plan": {
        "block": {
          "attributes": {
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "product": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "publisher": {
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
    "description_kind": "plain"
  },
  "version": 0
}`

func AzurermSharedImageSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermSharedImage), &result)
	return &result
}
