package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermStorageAccountBlobContainerSas = `{
  "block": {
    "attributes": {
      "cache_control": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "connection_string": {
        "description_kind": "plain",
        "required": true,
        "sensitive": true,
        "type": "string"
      },
      "container_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "content_disposition": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "content_encoding": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "content_language": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "content_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "expiry": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "https_only": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ip_address": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "sas": {
        "computed": true,
        "description_kind": "plain",
        "sensitive": true,
        "type": "string"
      },
      "start": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "permissions": {
        "block": {
          "attributes": {
            "add": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "create": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "delete": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "delete_version": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "execute": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "find": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "list": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "move": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "ownership": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "permissions": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "read": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "set_immutability_policy": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "tags": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "write": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
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

func AzurermStorageAccountBlobContainerSasSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermStorageAccountBlobContainerSas), &result)
	return &result
}
