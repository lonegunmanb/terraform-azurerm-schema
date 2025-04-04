package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermDataFactoryLinkedServiceSftp = `{
  "block": {
    "attributes": {
      "additional_properties": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "annotations": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "authentication_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "data_factory_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "host": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "host_key_fingerprint": {
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
      "integration_runtime_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "parameters": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "password": {
        "description_kind": "plain",
        "optional": true,
        "sensitive": true,
        "type": "string"
      },
      "port": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "private_key_content_base64": {
        "description_kind": "plain",
        "optional": true,
        "sensitive": true,
        "type": "string"
      },
      "private_key_passphrase": {
        "description_kind": "plain",
        "optional": true,
        "sensitive": true,
        "type": "string"
      },
      "private_key_path": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "skip_host_key_validation": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "username": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "key_vault_password": {
        "block": {
          "attributes": {
            "linked_service_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "secret_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "key_vault_private_key_content_base64": {
        "block": {
          "attributes": {
            "linked_service_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "secret_name": {
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
      "key_vault_private_key_passphrase": {
        "block": {
          "attributes": {
            "linked_service_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "secret_name": {
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

func AzurermDataFactoryLinkedServiceSftpSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermDataFactoryLinkedServiceSftp), &result)
	return &result
}
