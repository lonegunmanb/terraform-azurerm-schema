package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermStorageAccountLocalUser = `{
  "block": {
    "attributes": {
      "home_directory": {
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
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "password": {
        "computed": true,
        "description_kind": "plain",
        "sensitive": true,
        "type": "string"
      },
      "sid": {
        "computed": true,
        "description_kind": "plain",
        "sensitive": true,
        "type": "string"
      },
      "ssh_key_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "ssh_password_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "storage_account_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "permission_scope": {
        "block": {
          "attributes": {
            "resource_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "service": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "permissions": {
              "block": {
                "attributes": {
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
                  "list": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "read": {
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
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "ssh_authorized_key": {
        "block": {
          "attributes": {
            "description": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "key": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
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

func AzurermStorageAccountLocalUserSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermStorageAccountLocalUser), &result)
	return &result
}
