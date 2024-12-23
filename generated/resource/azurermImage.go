package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermImage = `{
  "block": {
    "attributes": {
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
      "source_virtual_machine_id": {
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
      },
      "zone_resilient": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      }
    },
    "block_types": {
      "data_disk": {
        "block": {
          "attributes": {
            "blob_uri": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "caching": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "disk_encryption_set_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "lun": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "managed_disk_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "size_gb": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "storage_type": {
              "description": "The type of storage disk",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "os_disk": {
        "block": {
          "attributes": {
            "blob_uri": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "caching": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "disk_encryption_set_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "managed_disk_id": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "os_state": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "os_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "size_gb": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "storage_type": {
              "description": "The type of storage disk",
              "description_kind": "plain",
              "required": true,
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

func AzurermImageSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermImage), &result)
	return &result
}
