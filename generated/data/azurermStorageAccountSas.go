package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermStorageAccountSas = `{
  "block": {
    "attributes": {
      "connection_string": {
        "description_kind": "plain",
        "required": true,
        "sensitive": true,
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
      "ip_addresses": {
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
      "signed_version": {
        "description_kind": "plain",
        "optional": true,
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
              "required": true,
              "type": "bool"
            },
            "create": {
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            },
            "delete": {
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            },
            "filter": {
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            },
            "list": {
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            },
            "process": {
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            },
            "read": {
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            },
            "tag": {
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            },
            "update": {
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            },
            "write": {
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "resource_types": {
        "block": {
          "attributes": {
            "container": {
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            },
            "object": {
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            },
            "service": {
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "services": {
        "block": {
          "attributes": {
            "blob": {
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            },
            "file": {
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            },
            "queue": {
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            },
            "table": {
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
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

func AzurermStorageAccountSasSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermStorageAccountSas), &result)
	return &result
}
